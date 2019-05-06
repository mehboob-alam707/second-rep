package main

import "fmt"
 

func main(){

    m := make (map[string]interface{})
    m["k1"] = 7
    m["k2"] = ""
    fmt.Println("mapa: ", m)
	
    v1 := m["k2"]
    fmt.Println("v1: ", v1)
	
    _, prs := m["k2"]
    fmt.Println("prs:", prs)
	
    n := map[string]interface{}{"foo": 1, "bar": "some name", "doo": 7.35}
    fmt.Println("map:", n)
    fmt.Println("doo val:", n["doo"])
	
}
