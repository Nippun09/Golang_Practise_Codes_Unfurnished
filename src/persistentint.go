package main

import "fmt"

type Studentmarks struct {
	Marks        int
	Subject      string
	Student_name string
}

type StudentsSlice struct {
	Students []Studentmarks
}

func main() {

	s1 := Studentmarks{
		Marks:        5,
		Subject:      "english",
		Student_name: "whatever",
	}

	s2 := Studentmarks{
		Marks:        7,
		Subject:      "english",
		Student_name: "whatever2",
	}

	s3 := Studentmarks{
		Marks:        3,
		Subject:      "english",
		Student_name: "whatever3",
	}

	s4 := Studentmarks{
		Marks:        5,
		Subject:      "maths",
		Student_name: "whatever3",
	}

	s5 := Studentmarks{
		Marks:        8,
		Subject:      "history",
		Student_name: "whatever2",
	}

	studslice := []Studentmarks{s1, s2, s3, s4, s5}

	Studentslicedef := StudentsSlice{
		Students: studslice,
	}

	//var Finalmap map[string]int

	for studentmarkvar := range Studentslicedef.Students {

		fmt.Printf("%T", studentmarkvar)
		// if student,found:=Finalmap[studentmarkvar.Student_name]; found{

		// }
	}
}
