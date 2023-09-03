package src

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestCalculateCalorieNeeds(t *testing.T) {
	testCases := []struct {
		name             string
		age              int
		weightKg         float64
		heightCm         float64
		gender           string
		physicalActivity float64
		expectedCalories float64
		expectedError    bool
	}{
		{
			name:             "Male with moderate activity",
			age:              30,
			weightKg:         70.0,
			heightCm:         170.0,
			gender:           Male,
			physicalActivity: 1.55,
			expectedCalories: 2591.09,
			expectedError:    false,
		},
		{
			name:             "Female with low activity",
			age:              25,
			weightKg:         60.0,
			heightCm:         160.0,
			gender:           Female,
			physicalActivity: 1.2,
			expectedCalories: 1667.81,
			expectedError:    false,
		},
		{
			name:             "Child with low activity",
			age:              15,
			weightKg:         60.0,
			heightCm:         160.0,
			gender:           Female,
			physicalActivity: 1.0,
			expectedCalories: 1667.81,
			expectedError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calories, err := CalculateCalories(tc.age, tc.weightKg, tc.heightCm, tc.gender, tc.physicalActivity)

			if tc.expectedError && err == nil {
				t.Error("Expected an error but got none.")
			}

			if !tc.expectedError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if assert.IsEqual(calories, tc.expectedCalories) {
				t.Errorf("Expected %.2f calories but got %.2f", tc.expectedCalories, calories)
			}
		})
	}
}

func TestCalculateCalorieInvalidGender(t *testing.T) {
	testCases := []struct {
		name             string
		age              int
		weightKg         float64
		heightCm         float64
		gender           string
		physicalActivity float64
		expectedCalories float64
		expectedError    bool
	}{
		{
			name:             "Invalid gender",
			age:              40,
			weightKg:         80.0,
			heightCm:         175.0,
			gender:           "invalid",
			physicalActivity: 1.6,
			expectedCalories: 0,
			expectedError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calories, err := CalculateCalories(tc.age, tc.weightKg, tc.heightCm, tc.gender, tc.physicalActivity)

			if tc.expectedError && err == nil {
				t.Error("Expected an error but got none.")
			}

			if !tc.expectedError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}

			if assert.IsEqual(calories, tc.expectedCalories) && assert.IsEqual(tc.expectedError, false) {
				t.Errorf("Expected %.2f calories but got %.2f", tc.expectedCalories, calories)
			}
		})
	}
}
