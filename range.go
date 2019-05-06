package main

import "fmt"
 

func main(){

	sum := 0
	nums := []int{2,3,4}
	
	for _, ele := range nums{
		sum +=  ele
	}
	fmt.Println("Sum: ", sum)
	
	for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
	
	
	kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s : %s\n", k, v)
	fmt.Println(k,v)
    }
	
	for k := range kvs {
        fmt.Println("key:", k)
    }
	
}
