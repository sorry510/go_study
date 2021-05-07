package main

import "fmt"

type Values map[string][]string

func (v Values) Get(key string) string {
	if vs := v[key]; len(vs) > 0 {
		return vs[len(vs)-1]
	}
	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func main() {
	m := Values{"lang": {"cn", "en"}}
	(&m).Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))
	fmt.Println(m.Get("item"))
	fmt.Println(m)

	m = nil
	fmt.Println(m.Get("item")) // ""
	// m.Add("item", "3")         // panic: assignment to entry in nil map
}
