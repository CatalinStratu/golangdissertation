package src

// Constants for calorie calculation
const (
	Male   = "male"
	Female = "female"
)

// Harris-Benedict constants
const (
	BMRMale   = 88.362
	BMRFemale = 447.593
)

type Patient struct {
	Age              int     `json:"age"`
	Weight           float64 `json:"weight"`
	Height           float64 `json:"height"`
	Gender           string  `json:"gender"`
	PhysicalActivity float64 `json:"physical_activity"`
}

type Recipe struct {
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
		Recipe struct {
			Uri    string `json:"uri"`
			Label  string `json:"label"`
			Image  string `json:"image"`
			Images struct {
				THUMBNAIL struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"THUMBNAIL"`
				SMALL struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"SMALL"`
				REGULAR struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"REGULAR"`
				LARGE struct {
					Url    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"LARGE"`
			} `json:"images"`
			Source          string   `json:"source"`
			Url             string   `json:"url"`
			ShareAs         string   `json:"shareAs"`
			Yield           int      `json:"yield"`
			DietLabels      []string `json:"dietLabels"`
			HealthLabels    []string `json:"healthLabels"`
			Cautions        []string `json:"cautions"`
			IngredientLines []string `json:"ingredientLines"`
			Ingredients     []struct {
				Text     string `json:"text"`
				Quantity int    `json:"quantity"`
				Measure  string `json:"measure"`
				Food     string `json:"food"`
				Weight   int    `json:"weight"`
				FoodId   string `json:"foodId"`
			} `json:"ingredients"`
			Calories          int      `json:"calories"`
			GlycemicIndex     int      `json:"glycemicIndex"`
			TotalCO2Emissions int      `json:"totalCO2Emissions"`
			Co2EmissionsClass string   `json:"co2EmissionsClass"`
			TotalWeight       int      `json:"totalWeight"`
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
				Label        string `json:"label"`
				Tag          string `json:"tag"`
				SchemaOrgTag string `json:"schemaOrgTag"`
				Total        int    `json:"total"`
				HasRDI       bool   `json:"hasRDI"`
				Daily        int    `json:"daily"`
				Unit         string `json:"unit"`
				Sub          struct {
				} `json:"sub"`
			} `json:"digest"`
		} `json:"recipe"`
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
	} `json:"hits"`
}

//Products
type Food struct {
	FoodID        string `json:"foodId"`
	Label         string `json:"label"`
	KnownAs       string `json:"knownAs"`
	Nutrients     Nutrients
	Category      string `json:"category"`
	CategoryLabel string `json:"categoryLabel"`
	Image         string `json:"image"`
}

type Nutrients struct {
	ENERCKCAL float64 `json:"ENERC_KCAL"`
	PROCNT    float64 `json:"PROCNT"`
	FAT       float64 `json:"FAT"`
	CHOCDF    float64 `json:"CHOCDF"`
	FIBTG     float64 `json:"FIBTG"`
}

type Measure struct {
	URI       string  `json:"uri"`
	Label     string  `json:"label"`
	Weight    float64 `json:"weight"`
	Qualified []Qualified
}

type Qualified struct {
	Qualifiers []Qualifier
	Weight     float64 `json:"weight"`
}

type Qualifier struct {
	URI   string `json:"uri"`
	Label string `json:"label"`
}

type Hints struct {
	Food     Food
	Measures []Measure
}

type JSONData struct {
	Text  string  `json:"text"`
	Hints []Hints `json:"hints"`
}