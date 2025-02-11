package main

import "fmt"

type people struct {
	name    string
	phNo    int
	address string
}

func (p *people) updateName(updatedName string) people {
	p.name = updatedName
	return *p
}
func (p *people) printInfo() {
	fmt.Println("Name:", p.name)
	fmt.Println("Phone Number:", p.phNo)
	fmt.Println("Address:", p.address)
}
func main() {
	p := people{
		name:    "ArShar",
		phNo:    990990990,
		address: "JK06",
	}
	p.updateName("Ryan")
	p.printInfo()
}
