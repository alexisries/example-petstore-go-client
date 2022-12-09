package main

import (
	"context"
	"fmt"
	"petstore/petstore"
)

func main() {
	client, err := petstore.NewClientWithResponses("http://localhost:8080/api/v3")
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	pet, err := client.GetPetByIdWithResponse(ctx, 100)
	if err != nil {
		panic(err)
	}
	statuscode := pet.StatusCode()
	if statuscode != 200 {
		panic(fmt.Sprintf("error %d from server : %s", statuscode, string(pet.Body)))
	}
	fmt.Println(pet.JSON200.Name)

}