package main

import "sync"

type TopLevel struct {
	id    int
	What  string
	Final string
}
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Cast   []string `json:"cast"`
	Genres []string `json:"genres"`
}

func main() {
	var wg sync.WaitGroup
	go LargeFile(&wg)
	wg.Add(1)
	go StreaminJson(&wg)
	wg.Add(1)
	go F2F(&wg)
	wg.Add(1)
	go S2F(&wg)
	wg.Add(1)
	wg.Wait()

}
