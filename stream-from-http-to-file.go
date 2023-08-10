package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func S2F() {
	response, err := http.Get("https://raw.githubusercontent.com/prust/wikipedia-movie-data/master/movies.json")

	if err != nil {
		panic(err)
	}
	decoder := bufio.NewReader(response.Body)
	file, err := os.Create("./files/S2F.json")
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	if err != nil {
		panic(err)
	}
	count := 0
	readBytes := make([]byte, 1000)
	for {
		if count == 20 {
			break
		}
		n, err := decoder.Read(readBytes)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		writer.Write(readBytes[:n])
	}
}