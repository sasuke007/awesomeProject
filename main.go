package main

import (
	"errors"
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	//	dog := Dog{
	//		Name:  "Samson",
	//		Breed: "German Shepard",
	//	}
	//	PrintInfo(&dog)
	//	myJson := `
	//[
	//	{
	//		"first_name":"Clark",
	//		"last_name":"Kent",
	//		"hair_color":"black",
	//		"has_dog":true
	//	},
	//	{
	//		"first_name":"Bruce",
	//		"last_name":"Wayne",
	//		"hair_color":"black",
	//		"has_dog":false
	//	}
	//]`
	//	var unmarshalled []Person
	//	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	//	if err != nil {
	//		log.Println("Error in unmarshalling ", err)
	//	}
	//	log.Println(unmarshalled)
	//	// write json from a struct
	//	var mySlice []Person
	//	var m1 Person
	//	m1.FirstName = "Walley"
	//	m1.LastName = "West"
	//	m1.HairColor = "red"
	//	m1.HasDog = false
	//	mySlice = append(
	//		mySlice,
	//		m1,
	//	)
	//
	//	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	log.Println(string(newJson))

}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func PrintInfo(a Animal) {
	log.Println(a.Says(), a.NumberOfLegs())
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		log.Println("Division not possible")
		return 0, errors.New("division by 0 not supported")
	}
	return x / y, nil
}
