package main

import (
	"fmt"
	"time"
)

// Function to validate date string in format DD-MM-YYYY
func validateDate(dateStr string) (time.Time, error) {
	// Parse date with DD-MM-YYYY format
	layout := "02-01-2006"

	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format. Expected DD-MM-YYYY")
	}

	return date, nil
}

// Function to check if start_date is before end_date
func validateDates(startDateStr, endDateStr string) error {
	// Validate start date
	startDate, err := validateDate(startDateStr)
	if err != nil {
		return err
	}

	// Validate end date
	endDate, err := validateDate(endDateStr)
	if err != nil {
		return err
	}

	// Check if start date is before end date
	if startDate.After(endDate) {
		return fmt.Errorf("start date cannot be after end date")
	}

	return nil
}

func checkDateInRange(dateStr, startDateStr, endDateStr string) (bool, error) {
	// Validate and parse the start date
	startDate, err := validateDate(startDateStr)
	if err != nil {
		return false, err
	}

	// Validate and parse the end date
	endDate, err := validateDate(endDateStr)
	if err != nil {
		return false, err
	}

	// Validate and parse the date to be checked
	givenDate, err := validateDate(dateStr)
	if err != nil {
		return false, err
	}

	// Check if the given date falls within the range [start_date, end_date]
	if givenDate.After(startDate) && givenDate.Before(endDate) {
		return true, nil
	}
	// If the given date is equal to the start or end date, it's still valid
	if givenDate.Equal(startDate) || givenDate.Equal(endDate) {
		return true, nil
	}
	return false, nil
}
