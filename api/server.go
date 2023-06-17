package api

import (
	"fmt"
	"strings"

	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/token"
	"github.com/AntoninoAdornetto/mogged-lift-tracker-service/util"
	"github.com/gin-contrib/cors"
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
	server.setupCors(server.router)
	return server, nil
}

func (server *Server) setupCors(router *gin.Engine) {
	config := cors.DefaultConfig()
	origins := strings.Split(server.config.AllowedOrigins, ",")
	fmt.Println(origins)
	config.AllowOrigins = origins
	config.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(config))
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/createUser", server.createUser)
	router.POST("/auth/login", server.login)
	router.POST("/auth/renewToken", server.renewToken)

	protected := router.Group("/").Use(authMiddleware(server.tokenMaker))

	protected.GET("/getUserByEmail/:email", server.getUserByEmail)
	protected.PATCH("/updateUser", server.updateUser)
	protected.PATCH("/changePassword", server.changePassword)
	protected.DELETE("/deleteUser/:id", server.deleteUser)

	protected.POST("/createProfile", server.createProfile)
	protected.GET("/getProfile/:user_id", server.getProfile)
	protected.PATCH("/updateProfile", server.updateProfile)
	protected.DELETE("/deleteProfile/:user_id", server.deleteProfile)

	protected.POST("/createExercise", server.createExercise)
	protected.GET("/getExercise/:id/:user_id", server.getExercise)
	protected.GET("/getExerciseByName/:exercise_name/:user_id", server.getExerciseByName)
	protected.GET("/listExercises/:user_id", server.listExercises)
	protected.PATCH("/updateExercise", server.updateExercise)
	protected.DELETE("/deleteExercise/:id/:user_id", server.deleteExercise)

	protected.POST("/createWorkout", server.createWorkout)
	protected.GET("/getWorkout/:id/:user_id", server.getWorkout)
	protected.GET("/listWorkouts", server.listWorkouts)
	protected.PATCH("/updateWorkout", server.updateWorkout)
	protected.DELETE("/deleteWorkout/:id", server.deleteWorkout)

	protected.GET("/getLift/:id", server.getLift)
	protected.GET("/listLiftsFromWorkout/:workout_id", server.listLiftsFromWorkout)
	protected.GET("/getMaxLifts/:limit", server.getMaxLifts)
	protected.GET("/getMaxLiftsByExercise/:exercise_name", server.getMaxLiftsByExercise)
	protected.GET("/getMaxLiftsByMuscleGroup/:muscle_group", server.getMaxLiftsByMuscleGroup)
	protected.GET("/getMaxRepLifts/:limit", server.getMaxRepLifts)
	protected.PATCH("/updateLift", server.updateLift)
	protected.DELETE("/deleteLift/:id", server.deleteLift)

	protected.POST("/createTemplate", server.createTemplate)
	protected.GET("/getTemplate/:id", server.getTemplate)
	protected.GET("/listTemplates", server.listTemplates)
	protected.PATCH("/updateTemplate", server.updateTemplate)
	protected.DELETE("/deleteTemplate/:id", server.deleteTemplate)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
