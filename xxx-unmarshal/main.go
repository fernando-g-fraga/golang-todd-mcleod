package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Recipe struct {
	Meals []struct {
		IDMeal                      string `json:"idMeal"`
		StrMeal                     string `json:"strMeal"`
		StrDrinkAlternate           any    `json:"strDrinkAlternate"`
		StrCategory                 string `json:"strCategory"`
		StrArea                     string `json:"strArea"`
		StrInstructions             string `json:"strInstructions"`
		StrMealThumb                string `json:"strMealThumb"`
		StrTags                     string `json:"strTags"`
		StrYoutube                  string `json:"strYoutube"`
		StrIngredient1              string `json:"strIngredient1"`
		StrIngredient2              string `json:"strIngredient2"`
		StrIngredient3              string `json:"strIngredient3"`
		StrIngredient4              string `json:"strIngredient4"`
		StrIngredient5              string `json:"strIngredient5"`
		StrIngredient6              string `json:"strIngredient6"`
		StrIngredient7              string `json:"strIngredient7"`
		StrIngredient8              string `json:"strIngredient8"`
		StrIngredient9              string `json:"strIngredient9"`
		StrIngredient10             string `json:"strIngredient10"`
		StrIngredient11             string `json:"strIngredient11"`
		StrIngredient12             string `json:"strIngredient12"`
		StrIngredient13             string `json:"strIngredient13"`
		StrIngredient14             string `json:"strIngredient14"`
		StrIngredient15             string `json:"strIngredient15"`
		StrIngredient16             any    `json:"strIngredient16"`
		StrIngredient17             any    `json:"strIngredient17"`
		StrIngredient18             any    `json:"strIngredient18"`
		StrIngredient19             any    `json:"strIngredient19"`
		StrIngredient20             any    `json:"strIngredient20"`
		StrMeasure1                 string `json:"strMeasure1"`
		StrMeasure2                 string `json:"strMeasure2"`
		StrMeasure3                 string `json:"strMeasure3"`
		StrMeasure4                 string `json:"strMeasure4"`
		StrMeasure5                 string `json:"strMeasure5"`
		StrMeasure6                 string `json:"strMeasure6"`
		StrMeasure7                 string `json:"strMeasure7"`
		StrMeasure8                 string `json:"strMeasure8"`
		StrMeasure9                 string `json:"strMeasure9"`
		StrMeasure10                string `json:"strMeasure10"`
		StrMeasure11                string `json:"strMeasure11"`
		StrMeasure12                string `json:"strMeasure12"`
		StrMeasure13                string `json:"strMeasure13"`
		StrMeasure14                string `json:"strMeasure14"`
		StrMeasure15                string `json:"strMeasure15"`
		StrMeasure16                any    `json:"strMeasure16"`
		StrMeasure17                any    `json:"strMeasure17"`
		StrMeasure18                any    `json:"strMeasure18"`
		StrMeasure19                any    `json:"strMeasure19"`
		StrMeasure20                any    `json:"strMeasure20"`
		StrSource                   any    `json:"strSource"`
		StrImageSource              any    `json:"strImageSource"`
		StrCreativeCommonsConfirmed any    `json:"strCreativeCommonsConfirmed"`
		DateModified                any    `json:"dateModified"`
	} `json:"meals"`
}

func main() {

	res, err := http.Get("https://www.themealdb.com/api/json/v1/1/search.php?s=Arrabiata")

	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Printf("Ocorreu um erro! Código %d \n body:%s\n", res.StatusCode, body)

	}

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s", body)
	fmt.Println("------------------")

	var receita Recipe

	err = json.Unmarshal(body, &receita)

	fmt.Println(receita.Meals)
}
