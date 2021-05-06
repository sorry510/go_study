package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`        // 可以覆盖原始名称
	Color  bool `json:"color,omitempty"` // omitempty选项，表示当Go语言结构体成员为空或零值时不生成该JSON对象（这里false为零值）
	Actors []string
	little string // 只有导出的结构体成员才会被编码,小写字母开头的不会被解析
}

func main() {
	var movies = []Movie{
		{Title: "hello", Year: 12, Color: true, Actors: []string{"zhang", "zhao"}, little: "lit"},
		{Title: "hello2", Year: 122, Color: true, Actors: []string{"li", "zhao"}, little: "lit"},
	}
	data, err := json.Marshal(movies) // 编组
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []Movie
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(titles[0].Title)
}
