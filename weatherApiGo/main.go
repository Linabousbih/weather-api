package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

type myjson struct{
	Name string `json:name`

	Weather []struct{
		Description string `json:"description"`
	} `json:"weather"`

	Main struct{
		Temp float64 `json:temp`
	}`json:main`
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}

	API_KEY := os.Getenv("API_KEY")
	BASE_URL := os.Getenv("BASE_URL")


	fmt.Print("Where do you want to check the weather: ")
	var city string
	fmt.Scanln(&city)


	requestURL := fmt.Sprintf("%s?appid=%s&q=%s", BASE_URL, API_KEY, city)


	response, err := http.Get(requestURL)
	if err != nil {
		fmt.Println("An error occurred:", err)
		os.Exit(1)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		var data myjson
		if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
			fmt.Println("Failed to decode JSON response:", err)
			os.Exit(1)
		}
		
		fmt.Printf("To day in %v the weather is %v and the temperature is %2.2fºC\n",data.Name,data.Weather[0].Description,data.Main.Temp-273.15)

		

	} else {
		fmt.Println("An error occurred.")
	}
}
