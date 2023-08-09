package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type TopLevel struct{
	id int
	What string
	Final string 
}
type Movie struct{
	Title string `json:"title"`
	Year int `json:"year"`
	Cast []string `json:"cast"`
	Genres []string `json:"genres"`

}
func main() {
	response, err := http.Get("https://raw.githubusercontent.com/prust/wikipedia-movie-data/master/movies.json")

	if err != nil{
		panic(err)
	}
	decoder := json.NewDecoder(response.Body)
	startDelim , err := decoder.Token()
	if err != nil{
		panic(err)
	}
	if startDelim != "["{
		fmt.Println(fmt.Errorf("expected `[' got %v"), startDelim)
	}
	count := 0
	for decoder.More(){
		if count == 20{
			break
		}
		movie := Movie{}
		if err != nil{
			fmt.Println("error")
			panic(err)
		}
		decoder.Decode(&movie)

		fmt.Println(movie)
		count++
	}

	

}