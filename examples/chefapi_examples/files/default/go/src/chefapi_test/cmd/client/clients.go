package main

import (
	"fmt"
	"os"

	"chefapi_test/testapi"
	"github.com/go-chef/chef"
)

func main() {
	// Use the default test org
	client := testapi.Client()

	// List initial clients
	clientList, err := client.Clients.List()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't list clients: ", err)
	}
	fmt.Printf("List initial clients %+v\n", clientList)

	// Define a Client object
	client1 := chef.ApiNewClient{
		Name: "client1",
		CreateKey: true,
		Validator: false,
	}
	fmt.Printf("Define client1 %+v\n", client1)

	// Define another Client object
	client2 := chef.ApiNewClient{
		Name: "client1",
		CreateKey: false,
		Validator: true,
		PublicKey: "----- BEGIN FAKE PUBLIC KEY -----",
	}
	fmt.Printf("Define client2 %+v\n", client1)

	// Create
	clientResult, err := client.Clients.Create(client1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't create client client1. ", err)
	}
	fmt.Printf("Added client1 %+v\n", clientResult)
	clientResult, err := client.Clients.Create(client2)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't create client client2. ", err)
	}
	fmt.Printf("Added client1 %+v\n", clientResult)

	// List clients
	clientList, err = client.Clients.List()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't list clients: ", err)
	}
	fmt.Printf("List clients after adding %+v\n", clientList)

	// Create a second time expect 409
	clientResult, err = client.Clients.Create(client1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't recreate client client1. ", err)
	}
	fmt.Printf("Added client1 %+v\n", clientResult)

	// Read client1 information
	serverClient, err := client.Clients.Get("client1")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't get client: ", err)
	}
	fmt.Printf("Get client1 %+v\n", serverClient)

	serverClient, err := client.Clients.Get("client2")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't get client2: ", err)
	}
	fmt.Printf("Get client2 %+v\n", serverClient)

	// update client - change the client name
	client1 = chef.ApiNewClient{
		Name: "clientchanged",
		ClientName: "clientchanged",
		Validator: true,
	}
	updateClient, err := client.Clients.Update("client1", client1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't update client: ", err)
	}
	fmt.Printf("Update client1 %+v\n", updateClient)

	// update client - change the validator status
	client2 = chef.ApiNewClient{
		Validator: false,
	}
	updateClient, err := client.Clients.Update("client2", client1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't update client: ", err)
	}
	fmt.Printf("Update client2 %+v\n", updateClient)

	// Info after update
	serverClient, err = client.Clients.Get("clientchanged")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't get client: ", err)
	}
	fmt.Printf("Get client1 after update %+v\n", serverClient)

	// Delete clients ignoring errors :)
	err = client.Clients.Delete("clientchanged")
	fmt.Printf("Delete client1 %+v\n", err)
	err = client.Clients.Delete("client2")
	fmt.Printf("Delete client2 %+v\n", err)

	// List clients
	clientList, err = client.Clients.List()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't list clients: ", err)
	}
	fmt.Printf("List clients after cleanup %+v\n", clientList)
}
