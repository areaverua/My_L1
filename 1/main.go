package main


import "fmt"


type Human struct{
	Name string
	SurName string
	Age int
}


func (h * Human) About(){
	fmt.Println(h.Name, h.SurName, h.Age)
}

type Action struct{
	Human
}


func main(){
	a := Action{Human: 
		Human{
		Name: "Jack", SurName: "Jons", Age: 18,}}
		
	a.About()
}