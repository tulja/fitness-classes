package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GymClass represents data about a record GymClass.
type GymClass struct {
	Id        string `json:"id"`
	ClassName string `json:"class_name"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Capacity  uint   `json:"capacity"`
}

// GymClasss map to seed record GymClass data.

var id int = 2

var GymClasses = map[string]GymClass{
	"1": {ClassName: "Pilates", StartDate: "25-01-2025", EndDate: "31-01-2025", Capacity: 56},
	"2": {ClassName: "Yoga", StartDate: "01-02-2025", EndDate: "05-02-2025", Capacity: 17},
}

func mapValuesToSliceForGymClass(m map[string]GymClass) []GymClass {
	values := []GymClass{}
	for id, value := range m {
		value.Id = id
		values = append(values, value)
	}
	return values
}

// getGymClasses responds with the list of all GymClasss as JSON.
func GetGymClasses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mapValuesToSliceForGymClass(GymClasses))
}

// createGymClasss adds an GymClass from JSON received in the request body.
func CreateGymClass(c *gin.Context) {
	var newGymClass GymClass

	// Call BindJSON to bind the received JSON to
	// newGymClass.
	if err := c.BindJSON(&newGymClass); err != nil {
		return
	}

	if err := validateDates(newGymClass.StartDate, newGymClass.EndDate); err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	classDateFallsInExistingTimeRange, err := checkDatesOfExistingClasses(newGymClass.StartDate, newGymClass.EndDate)
	if err != nil {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}
	if classDateFallsInExistingTimeRange {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "given Dates overlap with an existing class please reconsider the dates"})
		return
	}

	id++
	// Add the new GymClass to the map.
	newGymClass.Id = strconv.Itoa(id)
	GymClasses[strconv.Itoa(id)] = newGymClass
	c.IndentedJSON(http.StatusCreated, newGymClass)
}

// getGymClassByID locates the GymClass whose ID value matches the id
// parameter sent by the client, then returns that GymClass as a response.
func GetGymClassByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of GymClasss, looking for
	// an GymClass whose ID value matches the parameter.
	for GymClassId, GymClass := range GymClasses {
		if GymClassId == id {
			GymClass.Id = id
			c.IndentedJSON(http.StatusOK, GymClass)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "GymClass not found"})
}

// getGymClassByClassName locates the GymClass whose ID value matches the id
// parameter sent by the client, then returns that GymClass as a response.
func GetGymClassByClassName(c *gin.Context) {
	className := c.Param("class_name")

	// Loop over the list of GymClasss, looking for
	// an GymClass whose ClassName value matches the parameter.
	for GymClassId, GymClass := range GymClasses {
		if GymClass.ClassName == className {
			GymClass.Id = GymClassId
			c.IndentedJSON(http.StatusOK, GymClass)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "GymClass not found"})
}

// deleteGymClassByID locates the GymClass whose ID value matches the id and deletes it
// parameter sent by the client, then returns if deletion is success or failure.
func DeleteGymClassByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of GymClasss, looking for
	// an GymClass whose ID value matches the parameter.
	for GymClassId := range GymClasses {
		if GymClassId == id {
			delete(GymClasses, id)
			c.IndentedJSON(http.StatusOK, "Deleted id "+id+" Successfully!")
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "GymClass not found"})
}

// Function to check if the given date falls within the start and end date range
func checkDatesOfExistingClasses(givenStartDate string, givenEndDate string) (bool, error) {

	for _, GymClass := range GymClasses {
		startDateOfClass := GymClass.StartDate
		endDateOfClass := GymClass.EndDate

		isGivenStartDateInRange, err := checkDateInRange(givenStartDate, startDateOfClass, endDateOfClass)
		if err != nil {
			return false, err
		}

		isGivenEndDateInRange, err := checkDateInRange(givenEndDate, startDateOfClass, endDateOfClass)
		if err != nil {
			return false, err
		}

		if isGivenStartDateInRange && isGivenEndDateInRange {
			return true, nil
		}
	}

	return false, nil
}
