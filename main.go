package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math"
	"microservice/helpers"
	"microservice/src"
	"net/http"
	"strconv"
	"sync"
)

type Patient struct {
	age              int
	weightKg         float64
	heightCm         float64
	gender           string
	physicalActivity float64
	expectedCalories float64
}

type Meal struct {
	name    string
	percent float64
}

func FindMax(products []Recipe) []Recipe {
	max := products[0]
	for _, product := range products {
		if product.Calories > max.Calories {
			max = product
		}

	}
	var recipe []Recipe
	recipe = append(recipe, max)
	return recipe
}

type SubNutrient struct {
	Label        string  `json:"label"`
	Tag          string  `json:"tag"`
	SchemaOrgTag string  `json:"schemaOrgTag"`
	Total        float64 `json:"total"`
	HasRDI       bool    `json:"hasRDI"`
	Daily        float64 `json:"daily"`
	Unit         string  `json:"unit"`
}

type ResponseRecipe struct {
	Uri    string `json:"uri"`
	Label  string `json:"label"`
	Image  string `json:"image"`
	Images struct {
		THUMBNAIL struct {
			Url    string  `json:"url"`
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"THUMBNAIL"`
		SMALL struct {
			Url    string  `json:"url"`
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"SMALL"`
		REGULAR struct {
			Url    string  `json:"url"`
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"REGULAR"`
		LARGE struct {
			Url    string  `json:"url"`
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"LARGE"`
	} `json:"images"`
	Source            string   `json:"source"`
	Url               string   `json:"url"`
	ShareAs           string   `json:"shareAs"`
	Yield             float64  `json:"yield"`
	DietLabels        []string `json:"dietLabels"`
	HealthLabels      []string `json:"healthLabels"`
	Cautions          []string `json:"cautions"`
	IngredientLines   []string `json:"ingredientLines"`
	Calories          float64  `json:"calories"`
	GlycemicIndex     float64  `json:"glycemicIndex"`
	TotalCO2Emissions float64  `json:"totalCO2Emissions"`
	Co2EmissionsClass string   `json:"co2EmissionsClass"`
	TotalWeight       float64  `json:"totalWeight"`
	Ingredients       []src.Food
}

type Recipe struct {
	Uri    string `json:"uri"`
	Label  string `json:"label"`
	Image  string `json:"image"`
	Images struct {
		THUMBNAIL struct {
			Url    string  `json:"url"`
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"THUMBNAIL"`
		SMALL struct {
			Url    string  `json:"url"`
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"SMALL"`
		REGULAR struct {
			Url    string  `json:"url"`
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"REGULAR"`
		LARGE struct {
			Url    string  `json:"url"`
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
		} `json:"LARGE"`
	} `json:"images"`
	Source          string   `json:"source"`
	Url             string   `json:"url"`
	ShareAs         string   `json:"shareAs"`
	Yield           float64  `json:"yield"`
	DietLabels      []string `json:"dietLabels"`
	HealthLabels    []string `json:"healthLabels"`
	Cautions        []string `json:"cautions"`
	IngredientLines []string `json:"ingredientLines"`
	Ingredients     []struct {
		Text     string  `json:"text"`
		Quantity float64 `json:"quantity"`
		Measure  string  `json:"measure"`
		Food     string  `json:"food"`
		Weight   float64 `json:"weight"`
		FoodId   string  `json:"foodId"`
	} `json:"ingredients"`
	Calories          float64  `json:"calories"`
	GlycemicIndex     float64  `json:"glycemicIndex"`
	TotalCO2Emissions float64  `json:"totalCO2Emissions"`
	Co2EmissionsClass string   `json:"co2EmissionsClass"`
	TotalWeight       float64  `json:"totalWeight"`
	CuisineType       []string `json:"cuisineType"`
	MealType          []string `json:"mealType"`
	DishType          []string `json:"dishType"`
	Instructions      []string `json:"instructions"`
	Tags              []string `json:"tags"`
	ExternalId        string   `json:"externalId"`
	TotalNutrients    struct {
	} `json:"totalNutrients"`
	TotalDaily struct {
	} `json:"totalDaily"`
	Digest []struct {
		Label        string        `json:"label"`
		Tag          string        `json:"tag"`
		SchemaOrgTag string        `json:"schemaOrgTag"`
		Total        float64       `json:"total"`
		HasRDI       bool          `json:"hasRDI"`
		Daily        float64       `json:"daily"`
		Unit         string        `json:"unit"`
		SubNutrients []SubNutrient `json:"sub"`
	} `json:"digest"`
}

type RecipeStruct struct {
	From  int `json:"from"`
	To    int `json:"to"`
	Count int `json:"count"`
	Links struct {
		Self struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"self"`
		Next struct {
			Href  string `json:"href"`
			Title string `json:"title"`
		} `json:"next"`
	} `json:"_links"`
	Hits []struct {
		Recipe Recipe `json:"recipe"`
		Links  struct {
			Self struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"self"`
			Next struct {
				Href  string `json:"href"`
				Title string `json:"title"`
			} `json:"next"`
		} `json:"_links"`
	} `json:"hits"`
}

func fetchRecipes(url string, results chan<- []Recipe, wg *sync.WaitGroup, calories float64) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	var recipe RecipeStruct

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body for %s: \n", err)
		return
	}
	err = json.Unmarshal(body, &recipe)
	if err != nil {
		fmt.Printf("Error parsing JSON for %s: \n", err)
		return
	}
	caloriesPerMeal := calories + 0.1*calories
	var recipes []Recipe
	flag := 5
	for _, v := range recipe.Hits {
		if v.Recipe.Calories <= caloriesPerMeal {
			if flag != 0 {
				recipes = append(recipes, v.Recipe)
				flag--
				continue
			}
			break
		}
	}
	results <- FindMax(recipes)
}

func searchProduct(query string, wgi *sync.WaitGroup, results chan<- src.Food) {
	defer wgi.Done()

	response, err := http.Get("https://api.edamam.com/api/food-database/v2/parser?app_id=e3a6a808&app_key=44af244f2f5233960f919629f8f7dfcf&ingr=" + query)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body for %s: \n", err)
		return
	}

	var data src.JSONData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return
	}
	if len(data.Hints) > 0 {
		results <- data.Hints[0].Food

	}
}
func Recipes(ctx *gin.Context) {

	var wg sync.WaitGroup

	//var patient Patient
	//
	//patient.age = 30
	//patient.weightKg = 70.0
	//patient.heightCm = 170.0
	//patient.gender = src.Male
	//patient.physicalActivity = 1.55
	//
	//calories, err := src.CalculateCalories(patient.age, patient.weightKg, patient.heightCm, patient.gender, patient.physicalActivity)
	//
	//if err != nil {
	//	fmt.Println("erorr", err.Error())
	//	return
	//}
	param := ctx.Param("calories")
	calories, _ := strconv.ParseFloat(param, 8)
	results := make(chan []Recipe)

	var meals []Meal

	meals = append(meals,
		Meal{
			name:    "Breakfast",
			percent: 20,
		},
		Meal{
			name:    "Dinner",
			percent: 20,
		},
		Meal{
			name:    "Dinner",
			percent: 20,
		},
		Meal{
			name:    "Lunch",
			percent: 20,
		})

	for _, meal := range meals {
		wg.Add(1)
		caloriesPerMeal := helpers.CalculateValueFromPercent(calories, meal.percent)
		go fetchRecipes("https://api.edamam.com/api/recipes/v2?type=public&app_id=229871f5&app_key=2fbe0ee7d7b5857a57e48e6f58232deb&health=alcohol-free&cuisineType=Central%20Europe&calories=10-"+fmt.Sprintf("%v", math.Round(caloriesPerMeal+0.1*caloriesPerMeal))+"&mealType="+meal.name+"&diet=low-fat&co2EmissionsClass=B&random=true", results, &wg, caloriesPerMeal)
	}

	// Use a separate Goroutine to close the results channel when all Goroutines are done
	go func() {
		// Wait for all Goroutines to finish
		wg.Wait()
		close(results)
	}()

	var allRecipes []Recipe

	// Receive and process recipes
	for recipes := range results {
		allRecipes = append(allRecipes, recipes...)
	}
	var responseRecipe []ResponseRecipe
	for _, v := range allRecipes {
		var r ResponseRecipe
		r.Uri = v.Uri
		r.Images = v.Images
		r.Image = v.Image
		r.Source = v.Source
		r.Url = v.Url
		r.ShareAs = v.ShareAs
		r.Yield = v.Yield
		r.DietLabels = v.DietLabels
		r.HealthLabels = v.HealthLabels
		r.Cautions = v.Cautions
		r.IngredientLines = v.IngredientLines
		r.GlycemicIndex = v.GlycemicIndex
		r.TotalCO2Emissions = v.TotalCO2Emissions
		r.Co2EmissionsClass = v.Co2EmissionsClass
		r.TotalWeight = v.TotalWeight
		r.Label = v.Label
		r.Calories = v.Calories
		var wgi sync.WaitGroup
		ingredients := make(chan src.Food)
		for _, value := range v.Ingredients {
			wgi.Add(1)
			go searchProduct(value.Food, &wgi, ingredients)
		}

		go func() {
			// Wait for all Goroutines to finish
			wgi.Wait()
			close(ingredients)
		}()
		var foodIngredients []src.Food

		// Receive and process recipes
		for recipes := range ingredients {
			foodIngredients = append(foodIngredients, recipes)
		}
		r.Ingredients = foodIngredients

		responseRecipe = append(responseRecipe, r)
	}
	ctx.JSON(http.StatusOK, gin.H{"recipes": responseRecipe})
}

func main() {
	r := gin.Default()
	r.GET("api/v1/recipes/:calories", Recipes)
	r.POST("api/v1/nutrition-calculator", src.CalculateCaloriesGet)
	r.Run(":8088")
}
