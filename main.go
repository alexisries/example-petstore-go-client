package main

import (
	"context"
	"fmt"
	"petstore/petstore"
)

func main() {
	// initialize petstore api client
	client, err := petstore.NewClientWithResponses("https://petstore3.swagger.io/api/v3")
	if err != nil {
		panic(err)
	}
	// initialize context
	ctx := context.TODO()

	// get first Pet (id: 1) object
	petResponse, err := client.GetPetByIdWithResponse(ctx, 1)
	if err != nil {
		panic(err)
	}

	// check errors
	statuscode := petResponse.StatusCode()
	if statuscode != 200 {
		panic(fmt.Sprintf("error %d from server : %s", statuscode, string(petResponse.Body)))
	}

	// Unmarshall response json to Pet object
	var pet *petstore.Pet = petResponse.JSON200

	// Print 'name' attribute of Pet
	fmt.Println(pet.Name)
}