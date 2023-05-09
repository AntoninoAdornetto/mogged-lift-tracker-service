package api

import (
	db "github.com/AntoninoAdornetto/mogged-lift-tracker-service/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

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

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
