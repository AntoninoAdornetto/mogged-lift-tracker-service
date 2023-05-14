package api

import (
	"fmt"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/token"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("auth/login", server.login)

	router.POST("/createUser", server.createUser)
	router.GET("/getUserByEmail/:email", server.getUserByEmail)
	router.PATCH("/updateUser", server.updateUser)
	router.PATCH("/changePassword", server.changePassword)
	router.DELETE("/deleteUser/:id", server.deleteUser)

	router.POST("/createProfile", server.createProfile)
	router.GET("/getProfile/:user_id", server.getProfile)
	router.PATCH("/updateProfile", server.updateProfile)
	router.DELETE("/deleteProfile/:user_id", server.deleteProfile)

	router.POST("/createExercise", server.createExercise)
	router.GET("/getExercise/:id/:user_id", server.getExercise)
	router.GET("/getExerciseByName/:exercise_name/:user_id", server.getExerciseByName)
	router.GET("/listExercises/:user_id", server.listExercises)
	router.PATCH("/updateExercise", server.updateExercise)
	router.DELETE("/deleteExercise/:id/:user_id", server.deleteExercise)

	router.POST("/createWorkout", server.createWorkout)
	router.GET("/getWorkout/:id/:user_id", server.getWorkout)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
