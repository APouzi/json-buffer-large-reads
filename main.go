package main

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
	LargeFile()
	StreaminJson()

}