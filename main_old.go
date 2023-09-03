package main

type Keyword struct {
	Name     string
	Category string
}

type EcoscoreData struct {
	Adjustments struct {
		OriginsOfIngredients struct {
			AggregatedOrigins []struct {
				EpiScore            int    `json:"epi_score"`
				Origin              string `json:"origin"`
				Percent             int    `json:"percent"`
				TransportationScore int    `json:"transportation_score"`
			} `json:"aggregated_origins"`
			EpiScore                int      `json:"epi_score"`
			EpiValue                int      `json:"epi_value"`
			OriginsFromOriginsField []string `json:"origins_from_origins_field"`
			TransportationScore     int      `json:"transportation_score"`
			TransportationScores    struct {
				World int `json:"world"`
				Xk    int `json:"xk"`
			} `json:"transportation_scores"`
			TransportationValue  int `json:"transportation_value"`
			TransportationValues struct {
				World int `json:"world"`
				Xk    int `json:"xk"`
			} `json:"transportation_values"`
			Value  int `json:"value"`
			Values struct {
				World int `json:"world"`
				Xk    int `json:"xk"`
			} `json:"values"`
			Warning string `json:"warning"`
		} `json:"origins_of_ingredients"`
		Packaging struct {
			NonRecyclableAndNonBiodegradableMaterials int    `json:"non_recyclable_and_non_biodegradable_materials"`
			Value                                     int    `json:"value"`
			Warning                                   string `json:"warning"`
		} `json:"packaging"`
		ProductionSystem struct {
			Labels  []interface{} `json:"labels"`
			Value   int           `json:"value"`
			Warning string        `json:"warning"`
		} `json:"production_system"`
		ThreatenedSpecies struct {
			Warning string `json:"warning"`
		} `json:"threatened_species"`
	} `json:"adjustments"`
	Agribalyse struct {
		Warning string `json:"warning"`
	} `json:"agribalyse"`
	Missing struct {
		Categories  int `json:"categories"`
		Ingredients int `json:"ingredients"`
		Labels      int `json:"labels"`
		Origins     int `json:"origins"`
		Packagings  int `json:"packagings"`
	} `json:"missing"`
	MissingAgribalyseMatchWarning int `json:"missing_agribalyse_match_warning"`
	MissingKeyData                int `json:"missing_key_data"`
	Scores                        struct {
	} `json:"scores"`
	Status string `json:"status"`
}
type Product struct {
	ID                       string       `json:"_id"`
	Keywords                 []string     `json:"_keywords"`
	AddedCountriesTags       []string     `json:"added_countries_tags"`
	Allergens                string       `json:"allergens"`
	AllergensFromIngredients string       `json:"allergens_from_ingredients"`
	AllergensFromUser        string       `json:"allergens_from_user"`
	AllergensHierarchy       []string     `json:"allergens_hierarchy"`
	AllergensTags            []string     `json:"allergens_tags"`
	CategoriesProperties     struct{}     `json:"categories_properties"`
	CategoriesPropertiesTags []string     `json:"categories_properties_tags"`
	CheckersTags             []string     `json:"checkers_tags"`
	Code                     string       `json:"code"`
	CodesTags                []string     `json:"codes_tags"`
	Complete                 int          `json:"complete"`
	Completeness             float64      `json:"completeness"`
	CorrectorsTags           []string     `json:"correctors_tags"`
	Countries                string       `json:"countries"`
	CountriesHierarchy       []string     `json:"countries_hierarchy"`
	CountriesTags            []string     `json:"countries_tags"`
	CreatedT                 int          `json:"created_t"`
	Creator                  string       `json:"creator"`
	DataQualityBugsTags      []string     `json:"data_quality_bugs_tags"`
	DataQualityErrorsTags    []string     `json:"data_quality_errors_tags"`
	DataQualityInfoTags      []string     `json:"data_quality_info_tags"`
	DataQualityTags          []string     `json:"data_quality_tags"`
	DataQualityWarningsTags  []string     `json:"data_quality_warnings_tags"`
	EcoscoreData             EcoscoreData `json:"ecoscore_data"`
	EcoscoreGrade            string       `json:"ecoscore_grade"`
	EcoscoreTags             []string     `json:"ecoscore_tags"`
	EditorsTags              []string     `json:"editors_tags"`
	EntryDatesTags           []string     `json:"entry_dates_tags"`
	ExpirationDate           string       `json:"expiration_date"`
	FoodGroupsTags           []string     `json:"food_groups_tags"`
	ImageFrontSmallURL       string       `json:"image_front_small_url"`
	ImageFrontThumbURL       string       `json:"image_front_thumb_url"`
	ImageFrontURL            string       `json:"image_front_url"`
	ImageIngredientsSmallURL string       `json:"image_ingredients_small_url"`
	ImageIngredientsThumbURL string       `json:"image_ingredients_thumb_url"`
	ImageIngredientsURL      string       `json:"image_ingredients_url"`
	ImageSmallURL            string       `json:"image_small_url"`
	ImageThumbURL            string       `json:"image_thumb_url"`
	ImageURL                 string       `json:"image_url"`
	Images                   struct {
		// ... nested image structure ...
	} `json:"images"`
	InformersTags            []string `json:"informers_tags"`
	InterfaceVersionCreated  string   `json:"interface_version_created"`
	InterfaceVersionModified string   `json:"interface_version_modified"`
	Lang                     string   `json:"lang"`
	Languages                struct {
		// ... nested languages structure ...
	} `json:"languages"`
	LanguagesCodes struct {
		En int `json:"en"`
		Ro int `json:"ro"`
	} `json:"languages_codes"`
	LanguagesHierarchy []string `json:"languages_hierarchy"`
	LanguagesTags      []string `json:"languages_tags"`
	LastEditDatesTags  []string `json:"last_edit_dates_tags"`
	LastEditor         string   `json:"last_editor"`
	LastImageDatesTags []string `json:"last_image_dates_tags"`
	LastImageT         int      `json:"last_image_t"`
	LastModifiedBy     string   `json:"last_modified_by"`
	LastModifiedT      int      `json:"last_modified_t"`
	LC                 string   `json:"lc"`
	MainCountriesTags  []string `json:"main_countries_tags"`
	MaxImgID           string   `json:"max_imgid"`
	MiscTags           []string `json:"misc_tags"`
	NovaGroupDebug     string   `json:"nova_group_debug"`
	NovaGroupError     string   `json:"nova_group_error"`
	NovaGroupsTags     []string `json:"nova_groups_tags"`
	NutrientLevels     struct {
		// ... nested nutrient levels structure ...
	} `json:"nutrient_levels"`
	NutrientLevelsTags []string `json:"nutrient_levels_tags"`
	Nutriments         struct {
		// ... nested nutriments structure ...
	} `json:"nutriments"`
	NutritionDataPer             string   `json:"nutrition_data_per"`
	NutritionDataPreparedPer     string   `json:"nutrition_data_prepared_per"`
	NutritionGradesTags          []string `json:"nutrition_grades_tags"`
	NutritionScoreBeverage       int      `json:"nutrition_score_beverage"`
	NutritionScoreDebug          string   `json:"nutrition_score_debug"`
	NutritionScoreWarningNoFiber int      `json:"nutrition_score_warning_no_fiber"`
	PackagingMaterialsTags       []string `json:"packaging_materials_tags"`
	PackagingRecyclingTags       []string `json:"packaging_recycling_tags"`
	PackagingShapesTags          []string `json:"packaging_shapes_tags"`
	Packagings                   []struct {
		// ... nested packagings structure ...
	} `json:"packagings"`
	PhotographersTags    []string `json:"photographers_tags"`
	PnnsGroups1          string   `json:"pnns_groups_1"`
	PnnsGroups1Tags      []string `json:"pnns_groups_1_tags"`
	PnnsGroups2          string   `json:"pnns_groups_2"`
	PnnsGroups2Tags      []string `json:"pnns_groups_2_tags"`
	PopularityKey        int      `json:"popularity_key"`
	PopularityTags       []string `json:"popularity_tags"`
	ProductNameEn        string   `json:"product_name_en"`
	RemovedCountriesTags []string `json:"removed_countries_tags"`
	Rev                  int      `json:"rev"`
	ScansN               int      `json:"scans_n"`
	SelectedImages       struct {
		// ... nested selected images structure ...
	} `json:"selected_images"`
	Sortkey               int      `json:"sortkey"`
	States                string   `json:"states"`
	StatesHierarchy       []string `json:"states_hierarchy"`
	StatesTags            []string `json:"states_tags"`
	Traces                string   `json:"traces"`
	TracesFromIngredients string   `json:"traces_from_ingredients"`
	TracesFromUser        string   `json:"traces_from_user"`
	TracesHierarchy       []string `json:"traces_hierarchy"`
	TracesTags            []string `json:"traces_tags"`
	UniqueScansN          int      `json:"unique_scans_n"`
	UnknownNutrientsTags  []string `json:"unknown_nutrients_tags"`
	UpdateKey             string   `json:"update_key"`
	URL                   string   `json:"url"`
	WeighersTags          []string `json:"weighers_tags"`
}

type ProductData struct {
	Count     int       `json:"count"`
	Page      int       `json:"page"`
	PageCount int       `json:"page_count"`
	PageSize  int       `json:"page_size"`
	Products  []Product `json:"products"`
	Skip      int       `json:"skip"`
	// ... add more fields as needed ...
}

func isBreakfastFood(category string) bool {
	// Definim aici categoriile potrivite pentru micul dejun
	breakfastCategories := map[string]bool{
		"Protein": true,
		"Carbs":   true,
		"Fruits":  true,
	}

	return breakfastCategories[category]
}

//const openFoodFactsAPI = "https://ro.openfoodfacts.org/cgi/search.pl"
//
//func fetchData(url string, wg *sync.WaitGroup, results chan<- []map[string]interface{}) {
//	defer wg.Done()
//
//	req, err := http.NewRequest(http.MethodGet, url, nil)
//	if err != nil {
//		fmt.Printf("client: could not create request: %s\n", err)
//		return
//	}
//
//	res, err := http.DefaultClient.Do(req)
//	if err != nil {
//		fmt.Printf("client: error making http request: %s\n", err)
//		return
//	}
//	defer res.Body.Close()
//
//	resBody, err := io.ReadAll(res.Body)
//	if err != nil {
//		fmt.Printf("client: could not read response body: %s\n", err)
//		return
//	}
//
//	var data map[string]interface{}
//
//	err = json.Unmarshal(resBody, &data)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	results <- data
//}
//

//func NutritionPlan(ctx *gin.Context) {
//	var wg sync.WaitGroup
//	results := make(chan []map[string]interface{})
//
//	products := []Keyword{
//		{Name: "oua", Category: "Protein"},
//		{Name: "iaurt", Category: "Dairy"},
//		{Name: "ciocolata", Category: "Dairy"},
//		{Name: "Cereale", Category: "Carbs"},
//		{Name: "Fructe proaspete", Category: "Fruits"},
//		{Name: "Cârnați", Category: "Protein"},
//	}
//
//	// Afișam produsele pentru micul dejun
//	fmt.Println("Produse pentru micul dejun:")
//	for _, p := range products {
//		//if isBreakfastFood(p.Category) {
//		//fmt.Printf("%s\n", p.Name)
//		wg.Add(1)
//		go fetchData(fmt.Sprintf("%s?action=process&search_terms=%s&page_size=5&json=1", openFoodFactsAPI, p.Name), &wg, results)
//		//}
//	}
//
//	go func() {
//		wg.Wait()
//		close(results)
//	}()
//
//	var allData []map[string]interface{}
//	for data := range results {
//		allData = append(allData, data)
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"products": len(allData)})
//}

//func samain() {
//	//asdasd
//	r := gin.Default()
//	//	r.GET("/nutrition-plan", NutritionPlan)
//	r.POST("api/v1/nutrition-calculator", src.CalculateCalories)
//	r.Run(":8088")
//}

//package main
//
//import (
//"encoding/json"
//"fmt"
//"io"
//"net/http"
//"sync"
//)
//
//type Product struct {
//	ProductName string `json:"product_name"`
//}
//
//type SearchResult struct {
//	Products []Product `json:"products"`
//}
//
//func searchAndCollect(searchTerm string, ch chan<- []Product, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	// Construct the API URL
//	apiURL := fmt.Sprintf("https://ro.openfoodfacts.org/cgi/search.pl?search_terms=%s&json=1", searchTerm)
//
//	// Make the API request
//	response, err := http.Get(apiURL)
//	if err != nil {
//		fmt.Printf("Error making API request for %s: %v\n", searchTerm, err)
//		return
//	}
//	defer response.Body.Close()
//
//	// Read the response body
//	body, err := io.ReadAll(response.Body)
//	if err != nil {
//		fmt.Printf("Error reading response body for %s: %v\n", searchTerm, err)
//		return
//	}
//
//	// Parse the JSON response
//	var searchResult SearchResult
//	err = json.Unmarshal(body, &searchResult)
//	if err != nil {
//		fmt.Printf("Error parsing JSON for %s: %v\n", searchTerm, err)
//		return
//	}
//
//	// Send the products to the channel
//	ch <- searchResult.Products
//}
//
//func main() {
//	searchTerms := []string{"lapte", "oua", "branza", "peste"} // The search terms you want to use
//
//	var wg sync.WaitGroup
//	productChannel := make(chan []Product)
//
//	// Start Goroutines to search and collect results
//	for _, term := range searchTerms {
//		wg.Add(1)
//		go searchAndCollect(term, productChannel, &wg)
//	}
//
//	// Wait for all Goroutines to finish
//	go func() {
//		wg.Wait()
//		close(productChannel)
//	}()
//	var allProducts []Product
//	for products := range productChannel {
//		allProducts = append(allProducts, products...)
//	}
//	// Collect and print the products from the channel
//	fmt.Println("number of product", len(allProducts))
//	//for products := range productChannel {
//	//	fmt.Printf("Search results, nr %d:\n", len(products))
//	//	for i, product := range products {
//	//		fmt.Printf("%d. %s\n", i+1, product.ProductName)
//	//	}
//	//}
//}
//
//{
//"age":18,
//"weight":90,
//"height":180,
//"gender":1,
//"physical_activity":90
//}
