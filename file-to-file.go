package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func F2F() {
	f, err := os.Open("./files/Large-File.json")
	create, err := os.Create("Large-File2.json")

	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(bufio.NewReader(f))
	bW := bufio.NewWriter(create)
	encoder := json.NewEncoder(bW)

	defer bW.Flush()//Remember this! when using this bufio, even when in a wrapper needs to be flushed!

	t, err := decoder.Token()
	if t != "{" && err != nil {
		panic(err)
	}
	encoder.Encode(t)

	fmt.Println(t)
	topLevel := TopLevel{}
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		if t == "id" {
			encoder.Encode(t)
			t, err := decoder.Token()
			if err != nil {
				fmt.Print("error id")
			}
			topLevel.id = int(t.(float64))
			encoder.Encode(t)
		}
		if t == "what" {
			encoder.Encode(t)
			t, err := decoder.Token()
			if err != nil {
				fmt.Print("error what")
			}
			topLevel.What = t.(string)
			encoder.Encode(topLevel)
		}
		if t == "final" {
			encoder.Encode(t)
			t, err := decoder.Token()
			if err != nil {
				fmt.Print("error what")
			}
			topLevel.Final = t.(string)
			encoder.Encode(t)
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
			encoder.Encode(startDelim)
			// decoder.Token()
			for decoder.More() {
				movie := Movie{}
				err := decoder.Decode(&movie)
				fmt.Println(movie)
				encoder.Encode(movie)
				if err != nil {
					fmt.Println("error")
					panic(err)
				}
			}

		}
	}
	fmt.Println(topLevel)
	
}
