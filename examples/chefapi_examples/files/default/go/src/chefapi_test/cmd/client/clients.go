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
	fmt.Println("List initial clients", clientList)

	// Define an Client object
	client1 := chef.ApiNewClient("client1")
	fmt.Println("Define client1", client1)

	// Delete client1 ignoring errors :)
	err = client.Clients.Delete(client1.Name)

	// Create
	clientResult, err := client.Clients.Post(client1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't create client client1. ", err)
	}
	fmt.Println("Added client1", clientResult)

	// List clients
	clientList, err = client.Clients.List()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't list clients: ", err)
	}
	fmt.Println("List clients after adding client1", clientList)

	// Create a second time
	clientResult, err = client.Clients.Post(client1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't recreate client client1. ", err)
	}
	fmt.Println("Added client1", clientResult)

	// Read client1 information
	serverClient, err := client.Clients.Get("client1")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't get client: ", err)
	}
	fmt.Printf("Get client1 %+v\n", serverClient)

	// update client
	// TODO - update something about the client
	updateClient, err := client.Clients.Put(client1)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't update client: ", err)
	}
	fmt.Println("Update client1", updateClient)

	// Info after update
	serverClient, err = client.Clients.Get("client1")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't get client: ", err)
	}
	fmt.Printf("Get client1 after update %+v\n", serverClient)

	// Delete client ignoring errors :)
	err = client.Clients.Delete(client1.Name)
	fmt.Println("Delete client1", err)

	// List clients
	clientList, err = client.Clients.List()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't list clients: ", err)
	}
	fmt.Println("List clients after cleanup", clientList)
}
