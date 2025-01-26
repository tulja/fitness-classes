package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var booking_id int = 0

type Bookings struct {
	Id          string `json:"id"`
	ClassName   string `json:"class_name"`
	Date        string `json:"date"`
	BookingName string `json:"booking_name"`
}

var myBookings = make(map[string]Bookings)

func mapValuesToSliceForBookings(m map[string]Bookings) []Bookings {
	values := []Bookings{}
	for id, value := range m {
		value.Id = id
		values = append(values, value)
	}
	return values
}

// GetBookings responds with the list of all myBookings as JSON.
func GetBookings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mapValuesToSliceForBookings(myBookings))
}

// CreateBooking adds an Booking from JSON received in the request body.
func CreateBooking(c *gin.Context) {
	var newBooking Bookings

	// Call BindJSON to bind the received JSON to
	// newBooking.
	if err := c.BindJSON(&newBooking); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	// Loop over the list of GymClass, looking for
	// an GymClass whose class name value matches the parameter.
	for _, GymClass := range GymClasses {
		if GymClass.ClassName == newBooking.ClassName {
			_, err := validateDate(newBooking.Date)
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, err.Error())
				return
			}
			// Check if the given end date falls within the range [start_date, end_date] of existing class
			isBookingDateValid, err := checkDateInRange(newBooking.Date, GymClass.StartDate, GymClass.EndDate)
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, err.Error())
				return
			}
			if !isBookingDateValid {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "class not available in given date"})
				return
			}
			booking_id++
			// Add the new booking to the map.
			newBooking.Id = strconv.Itoa(booking_id)
			myBookings[strconv.Itoa(booking_id)] = newBooking
			GymClass.Capacity--
			c.IndentedJSON(http.StatusCreated, newBooking)
			return
		}
	}
}

// GetBookingByID locates the Bookings whose ID value matches the id
// parameter sent by the client, then returns that Booking as a response.
func GetBookingByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of myBookings, looking for
	// an Booking whose ID value matches the parameter.
	for booking_id, booking := range myBookings {
		if booking_id == id {
			booking.Id = id
			c.IndentedJSON(http.StatusOK, booking)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Booking not found"})
}

// GetBookingByBookingName locates the Booking whose ID value matches the id
// parameter sent by the client, then returns that GymClass as a response.
func GetBookingByBookingName(c *gin.Context) {
	bookingName := c.Param("booking_name")

	// Loop over the list of GymClasss, looking for
	// an GymClass whose ClassName value matches the parameter.
	for id, booking := range myBookings {
		if booking.BookingName == bookingName {
			booking.Id = id
			c.IndentedJSON(http.StatusOK, booking)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Booking not found"})
}

// DeleteBookingByID locates the Booking whose ID value matches the id and deletes it
// parameter sent by the client, then returns if deletion is success or failure.
func DeleteBookingByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of myBookings, looking for
	// an Booking whose ID value matches the parameter.
	for booking_id, booking := range myBookings {
		if booking_id == id {
			delete(myBookings, id)
			// Loop over the list of GymClass, looking for
			// an GymClass whose class name value matches the parameter.
			for _, GymClass := range GymClasses {
				if GymClass.ClassName == booking.ClassName {
					GymClass.Capacity++
					break
				}
			}
			c.IndentedJSON(http.StatusOK, "Deleted booking id "+id+" Successfully!")
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Booking not found"})
}
