package main

import "fmt"

// 结构体
type Books struct {
   title string
   author string
   subject string
   book_id int
}

func main() {
	var b = Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	fmt.Println(b)
	fmt.Println(Books{title: "haskell 语言", author: "www.runoob.com", subject: "haskell 语言教程", book_id: 6495407})

	var Book1 Books        /* 声明 Book1 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "rust 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "rust 语言教程"
	Book1.book_id = 6495407
	fmt.Println(Book1)
}
