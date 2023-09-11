package main

import "fmt"

type contantInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contantInfo
	contantInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string){
	p.firstName = newFirstName;
}

func updateAslice( s []string){
	s[0]= "bye"
}

func main() {
	/*
	alex := person{"Alex","Martinez"} // force to follow the order
	alex := person{firstName: "Alex", lastName: "Martinez"} // in this ways you can follow any other
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Martinez"
	fmt.Println(alex)
	fmt.Printf("%+v", alex) // to print each field of the struct
	Jim := person{
		firstName: "Jim \n",
		lastName:  "Doe \n",
		contantInfo: contantInfo{
			email:   "jimmail@mail.com \n",
			zipCode: 39,
		},
	}

	Jim.print()
	fmt.Printf("\n")
	Jim.updateName("Carlos")
	Jim.print()
	*/

	mySlice := []string{"HI","THERE", "I","STILL","LEARNING","GO"}

	updateAslice(mySlice)

	fmt.Print(mySlice)
	/*
		tricky test beacuse until now we said that when we pass a variable,
		we make a copy of that, so why this works?
	
	*/

}
