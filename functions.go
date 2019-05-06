package main

import "fmt"
 
func add(a, b int)int{
	return a+b
	}
	
func val(e, f int)(int, string){
	if e > f{
		return (e-f), "Positive"
	}else{
		return (e-f), "Negative"
	}
}

func ret3()(int, int, int){
	return 1199,100,99
}
	
func main(){

	fmt.Println("Enter 2 numbers ")
	
	var c,d int
	fmt.Scanf("%d %d",&c,&d)
	a := add(c,d) 
	fmt.Println("sum: ", a)
	
	e,f := val(c,d)
	fmt.Println(e," ", f)
	
	g, _ := val(d,c)
	fmt.Println(g)
	
	_, h := val(d,c)
	fmt.Println(h)

	p,q,r := ret3()
	fmt.Println(p,q,r)

	
}
