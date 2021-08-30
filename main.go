package main

import (
	"github.com/pluralsight/webservice/models"
	"fmt"
)

func main(){
	u := models.User{
		ID: 2,
		FirstName: "Tricia",
		LastName: "McMillan",
	}	

	fmt.Println(u)
}