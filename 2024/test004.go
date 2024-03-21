package main

import "fmt"

type struct1 struct {
    i1  int
    f1  float32
    str string
}

func main() {
	ms := struct1{10, 15.5, "Chris"}
	ms.f1 = 15.5
	
	ms1 := &struct1{10, 15.5, "Chris"}
	
	var ms2 *struct1
    ms2 = &struct1{10, 15.5, "Chris"}
	
	var ms3 = new(struct1)
	ms3.i1 = 123
	ms3.f1 = 32
	ms3.str = "chirs"
	
	fmt.Println(ms, ms1, ms2, ms3)
}