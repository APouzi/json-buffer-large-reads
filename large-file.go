package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LargeFile() {
	f, err := os.Open("Large-File.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(bufio.NewReader(f))

	t, err := decoder.Token()
	if t != "{" && err != nil {
		panic(err)
	}

	fmt.Println(t)
	topLevel := TopLevel{}
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		if t == "id" {
			t, err := decoder.Token()
			if err != nil {
				fmt.Print("error id")
			}
			topLevel.id = int(t.(float64))
		}
		if t == "what" {
			t, err := decoder.Token()
			if err != nil {
				fmt.Print("error what")
			}
			topLevel.What = t.(string)
		}
		if t == "final" {
			t, err := decoder.Token()
			if err != nil {
				fmt.Print("error what")
			}
			topLevel.Final = t.(string)
		}

		if t == "movies" {
			startDelim, err := decoder.Token()
			fmt.Println("startDelim", startDelim)
			if err != nil {
				panic(err)
			}
			if startDelim.(json.Delim) != '[' {
				fmt.Println(fmt.Errorf("expected `[' got %v"), startDelim)
			}
			// decoder.Token()
			for decoder.More() {
				movie := Movie{}
				err := decoder.Decode(&movie)
				fmt.Println(movie)
				if err != nil {
					fmt.Println("error")
					panic(err)
				}
			}

		}
	}
	fmt.Println(topLevel)
}
