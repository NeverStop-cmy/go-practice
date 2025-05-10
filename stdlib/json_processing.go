package stdlib

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func RunJSONProcessingDemo() {
	// 序列化
	book := Book{Title: "Go编程", Author: "Gopher", Year: 2023}
	jsonData, _ := json.Marshal(book)
	fmt.Println(string(jsonData))

	// 反序列化
	var book2 Book
	jsonStr := `{"title":"Go高级编程","author":"Advanced Gopher","year":2024}`
	json.Unmarshal([]byte(jsonStr), &book2)
	fmt.Println(book2)
}
