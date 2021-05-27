package main

import (
	"text/template"
	"os"
)

type fan struct{
	Rpm int
	Num_of_blades int
}


func main() {
	tlate:= template.Must(template.ParseFiles("pipe2argexp.gohtml"))

	fstruct:=fan{56, 4}

	tlate.Execute(os.Stdout, fstruct)
}

func (f fan) Summer(a int) int{
	return a
}