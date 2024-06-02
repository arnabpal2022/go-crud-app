package main

type Movie struct {
	ID string `json:"id"`
	ISBN string `json:"isbn"`
	Name string `json:"name"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
}

