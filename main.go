package main

import (
	"assignment1/students"
	"fmt"
)

func main() {
	var input int
	fmt.Print("Input ID Students : ")
	fmt.Scanf("%d", &input)

	numId := int(input)

	var students = []students.Students{
		//data-1
		{ID: 1, Name: "Muhammad Febri Andani",
			Address: students.Address{
				Street: "Jl. Jend S Parman No.21",
				City:   "Jakarta Selatan",
			},
			Job:    "Back End Enginner",
			Reason: "Ingin mendalami go programming"},
		//data-2
		{ID: 2, Name: "Syahwa Rahma Andini",
			Address: students.Address{
				Street: "Jl. Jend Sudirman No. 10",
				City:   "Jakarta Pusat",
			},
			Job:    "Barista",
			Reason: "Ingin mendalami go programming"},
		//data-3
		{ID: 3, Name: "Cintya Saras Adilla",
			Address: students.Address{
				Street: "Jl. Gatot Subroto Kuningan",
				City:   "Jakarta Selatan",
			},
			Job:    "IRT",
			Reason: "Ingin membantu suami dalam coding"},
	}

	//call function
	FindStudent(students, numId)

}

func FindStudent(students []students.Students, ID int) {
	var index int

	for i, s := range students {
		if s.ID == ID {
			index = i
			fmt.Println("======DATA STUDENTS======")
			fmt.Printf("Name\t\t: %v\n", students[index].Name)
			fmt.Printf("Address\t\t\n")
			fmt.Printf("\tCity\t: %v\n", students[index].Address.City)
			fmt.Printf("\tStreet\t: %v\n", students[index].Address.Street)
			fmt.Printf("Job\t\t: %v\n", students[index].Job)
			fmt.Printf("Reason\t\t: %v\n", students[index].Reason)
			fmt.Println("========================")
			return
		} else if i == (len(students) - 1) {
			fmt.Printf("ID : %d, Not Registered ", ID)
		}
	}
}
