package main

import (
	"testing"
)

// Test validateDate function
func TestValidateDate(t *testing.T) {
	tests := []struct {
		dateStr string
		wantErr bool
	}{
		{"26-01-2025", false}, // Valid date
		{"32-01-2025", true},  // Invalid day
		{"01-13-2025", true},  // Invalid month
		{"26-01-25", true},    // Invalid year format
		{"26-01-2025", false}, // Valid date (boundary case)
	}

	for _, tt := range tests {
		t.Run(tt.dateStr, func(t *testing.T) {
			_, err := validateDate(tt.dateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Test validateDates function
func TestValidateDates(t *testing.T) {
	tests := []struct {
		startDateStr string
		endDateStr   string
		wantErr      bool
	}{
		{"01-01-2025", "31-01-2025", false}, // Valid date range
		{"31-01-2025", "01-01-2025", true},  // Invalid date range, start date is after end date
		{"26-01-2025", "26-01-2025", false}, // Valid range where start = end
		{"01-01-2025", "01-01-2025", false}, // Valid range where start = end
		{"invalid", "01-01-2025", true},     // Invalid start date
		{"01-01-2025", "invalid", true},     // Invalid end date
	}

	for _, tt := range tests {
		t.Run(tt.startDateStr+" to "+tt.endDateStr, func(t *testing.T) {
			err := validateDates(tt.startDateStr, tt.endDateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateDates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Test checkDateInRange function
func TestCheckDateInRange(t *testing.T) {
	tests := []struct {
		dateStr      string
		startDateStr string
		endDateStr   string
		want         bool
		wantErr      bool
	}{
		{"15-01-2025", "01-01-2025", "31-01-2025", true, false},  // Date inside range
		{"01-01-2025", "01-01-2025", "31-01-2025", true, false},  // Date is equal to start date
		{"31-01-2025", "01-01-2025", "31-01-2025", true, false},  // Date is equal to end date
		{"01-02-2025", "01-01-2025", "31-01-2025", false, false}, // Date is outside range
		{"invalid", "01-01-2025", "31-01-2025", false, true},     // Invalid date to check
		{"25-01-2025", "01-01-2025", "invalid", false, true},     // Invalid end date
		{"25-01-2025", "invalid", "31-01-2025", false, true},     // Invalid start date
	}

	for _, tt := range tests {
		t.Run(tt.dateStr+" in range", func(t *testing.T) {
			got, err := checkDateInRange(tt.dateStr, tt.startDateStr, tt.endDateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkDateInRange() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("checkDateInRange() got = %v, want %v", got, tt.want)
			}
		})
	}
}
