package src

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math"
	"net/http"
)

func CalculateCaloriesGet(ctx *gin.Context) {
	//
	//headerKey := ctx.Request.Header.Get("key")
	//if headerKey != "$2a$04$F1Qu5mcERQFIlQez1B33t.EDlYod4EajUJej6vmXnZCxn8MXffhAy" {
	//	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	//	return
	//}
	patient := Patient{}

	b, err := io.ReadAll(ctx.Request.Body)
	defer ctx.Request.Body.Close()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	err = json.Unmarshal(b, &patient)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	calories, _ := CalculateCalories(patient.Age, patient.Weight, patient.Height, patient.Gender, patient.PhysicalActivity)
	ctx.JSON(http.StatusOK, gin.H{"calories": math.Round(calories)})
}

// calculateCalories calculates daily calorie needs based on input parameters
func CalculateCalories(age int, weightKg float64, heightCm float64, gender int, physicalActivity float64) (float64, error) {
	if gender != 1 && gender != 0 {
		return 0, fmt.Errorf("Invalid gender")
	}

	var bmr float64
	if gender == 1 {
		bmr = BMRMale + (13.397 * weightKg) + (4.799 * heightCm) - (5.677 * float64(age))
	} else if gender == 0 {
		bmr = BMRFemale + (9.247 * weightKg) + (3.098 * heightCm) - (4.330 * float64(age))
	}

	tdee := bmr * physicalActivity
	return tdee, nil
}

//	"sedentary":     1.2,
////	"lightlyActive": 1.375,
////	"moderatelyActive": 1.55,
////	"veryActive":    1.725,
////	"superActive":   1.9,
