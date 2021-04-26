package main

import "fmt"

// 结构体
type Books struct {
   title string
   author string
   subject string
   book_id int
}

type Job struct {
    name string
    hangye string
    money float32
}

type Man struct {
    name string
    age int
    sex string
    job Job
    book Books
}



func main() {
    // type1
	var b = Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407}
	fmt.Println(b)

    // type2
	fmt.Println(Books{title: "haskell 语言", author: "www.runoob.com", subject: "haskell 语言教程", book_id: 6495407})

    // type3
	var book1 Books        /* 声明 Book1 为 Books 类型 */
	/* book 1 描述 */
	book1.title = "rust 语言"
	book1.author = "www.runoob.com"
	book1.subject = "rust 语言教程"
	book1.book_id = 6495407
	fmt.Println(book1)

    
    var job1 Job
    str := "hello"
    job1 = Job{"yuan", str, 5555.0}
    job1.name = "yuan2"

    var man1 Man
    man1.name = "simple"
    man1.age = 30
    man1.sex = "man"
    man1.job = job1
    man1.book = book1
    fmt.Println(man1)
}
