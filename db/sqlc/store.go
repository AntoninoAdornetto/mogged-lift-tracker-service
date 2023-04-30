package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type NewUserTxResults struct {
	FirstName    string
	LastName     string
	EmailAddress string
	ID           uuid.UUID
}

func (store *Store) NewUserTx(ctx context.Context, args CreateUserParams) (NewUserTxResults, error) {
	user := &NewUserTxResults{}

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		err = store.CreateUser(ctx, CreateUserParams{
			LastName:     args.LastName,
			Password:     args.Password,
			FirstName:    args.FirstName,
			EmailAddress: args.EmailAddress,
		})
		if err != nil {
			return err
		}

		query, err := store.GetUser(ctx, args.EmailAddress)
		if err != nil {
			return err
		}

		userId, err := uuid.Parse(query.ID)

		statement := "INSERT INTO exercise (name, muscle_group, category, isStock, user_id) SELECT name, muscle_group, category, true, UUID_TO_BIN(?) FROM stock_exercise"
		_, err = store.db.ExecContext(ctx, statement, userId)
		if err != nil {
			return err
		}

		user.FirstName = query.FirstName
		user.LastName = query.LastName
		user.EmailAddress = query.EmailAddress
		user.ID = userId
		return nil
	})

	return *user, err
}

type WorkoutTxParams struct {
	UserID   string
	LiftsMap json.RawMessage
	Duration string
}

func (store *Store) WorkoutTx(ctx context.Context, args WorkoutTxParams) (Workout, error) {
	wo := &Workout{}

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		userId, err := uuid.Parse(args.UserID)
		if err != nil {
			return err
		}

		record, err := q.CreateWorkout(ctx, CreateWorkoutParams{
			Duration: args.Duration,
			UserID:   args.UserID,
		})
		if err != nil {
			return err
		}

		workoutId, err := record.LastInsertId()
		if err != nil {
			return err
		}

		liftsMap := make(map[string][]Lift)
		if err = json.Unmarshal(args.LiftsMap, &liftsMap); err != nil {
			return err
		}

		for _, lifts := range liftsMap {
			for _, lift := range lifts {
				_, err := q.CreateLift(ctx, CreateLiftParams{
					ExerciseName: lift.ExerciseName,
					WeightLifted: lift.WeightLifted,
					Reps:         lift.Reps,
					SetType:      lift.SetType,
					UserID:       userId.String(),
					WorkoutID:    int32(workoutId),
				})

				if err != nil {
					return err
				}
			}
		}

		_, err = q.UpdateWorkout(ctx, UpdateWorkoutParams{
			Lifts:  args.LiftsMap,
			ID:     int32(workoutId),
			UserID: args.UserID,
		})
		if err != nil {
			return err
		}

		query, err := q.GetWorkout(ctx, GetWorkoutParams{
			ID:     int32(workoutId),
			UserID: args.UserID,
		})
		if err != nil {
			return err
		}

		wo = &query
		return nil
	})

	return *wo, err
}
