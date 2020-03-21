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
	}
	fmt.Printf("Define client1 %+v\n", client1)

	// Delete client1 ignoring errors :)
	err = client.Clients.Delete(client1.Name)

	// Create
	clientResult, err := client.Clients.Create(client1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't create client client1. ", err)
	}
	fmt.Printf("Added client1 %+v\n", clientResult)

	// List clients
	clientList, err = client.Clients.List()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't list clients: ", err)
	}
	fmt.Printf("List clients after adding client1 %+v\n", clientList)

	// Create a second time
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

	// update client
	// TODO - update something about the client
	updateClient, err := client.Clients.Update("client1", client1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't update client: ", err)
	}
	fmt.Printf("Update client1 %+v\n", updateClient)

	// Info after update
	serverClient, err = client.Clients.Get("client1")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't get client: ", err)
	}
	fmt.Printf("Get client1 after update %+v\n", serverClient)

	// Delete client ignoring errors :)
	err = client.Clients.Delete(client1.Name)
	fmt.Printf("Delete client1 %+v\n", err)

	// List clients
	clientList, err = client.Clients.List()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't list clients: ", err)
	}
	fmt.Printf("List clients after cleanup %+v\n", clientList)
}
