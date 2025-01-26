package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                             // Allow all origins (use specific origins in production)
		AllowMethods:     []string{"GET", "POST", "DELETE"},         // Allow only specific methods
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // Allow specific headers
		AllowCredentials: true,
	}))

	router.GET("/classes", GetGymClasses)
	router.GET("/classes/:id", GetGymClassByID)
	router.GET("/classes/info/:class_name", GetGymClassByClassName)
	router.POST("/classes", CreateGymClass)
	router.DELETE("/classes/:id", DeleteGymClassByID)

	router.GET("/bookings", GetBookings)
	router.GET("/bookings/:id", GetBookingByID)
	router.GET("/bookings/info/:booking_name", GetBookingByBookingName)
	router.POST("/bookings", CreateBooking)
	router.DELETE("/bookings/:id", DeleteBookingByID)

	router.Run("localhost:9090")
}
