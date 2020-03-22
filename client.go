package chef

import "fmt"

type ApiClientService struct {
	client *Client
}

// ApiClient represents the native Go version of the deserialized Client type
// TODO: Doubt very many of these fields are there now, not in the doc
//
type ApiClient struct {
	Name       string `json:"name"`
	ClientName string `json:"clientname"`
	OrgName    string `json:"orgname"`
	Validator  bool   `json:"validator"`
	Uri        string `json:"uri,omitempty"`
	JsonClass  string `json:"json_class"`
	ChefType   string `json:"chef_type"`
}

// ApiNewClient structure to request a new client
type ApiNewClient struct {
	Name       string `json:"name"`
	ClientName string `json:"clientname,omitempty"`
	Validator  bool   `json:"validator,omitempty"`
	Admin      bool   `json:"admin,omitempty"` // not supported and ignored as of 12.1.0
	CreateKey  bool   `json:"create_key,omitempty"`
	PublicKey  bool   `json:"public_key,omitempty"`
}

// ApiNewClientresult
type ApiClientCreateResult struct {
	Uri     string  `json:"uri,omitempty"`
	ChefKey ChefKey `json:"chef_key,omitempty"`
}

// TODO this should probably be ???
type ApiClientListResult map[string]string

type ApiClientKey struct {
	Name           string `json:"name"`
	PublicKey      string `json:"public_key"`
	ExpirationDate string `json:"expiration_date"`
}

type ApiClientKeyListResultItem struct {
	Name    string `json:"name"`
	Expired bool   `json:"expired"`
}

type ApiClientKeyListResult []ApiClientKeyListResultItem

// String makes ApiClientListResult implement the string result
func (c ApiClientListResult) String() (out string) {
	for k, v := range c {
		out += fmt.Sprintf("%s => %s\n", k, v)
	}
	return out
}

// List lists the clients in the Chef server.
//
// Chef API docs: https://docs.chef.io/api_chef_server/#get-11
func (e *ApiClientService) List() (data ApiClientListResult, err error) {
	err = e.client.magicRequestDecoder("GET", "clients", nil, &data)
	return
}

// Create makes a Client on the chef server
//
// Chef API docs: https://docs.chef.io/api_chef_server.html#clients
// TODO pass in the structure instead of the fields
// TODO breaking change
func (e *ApiClientService) Create(client ApiNewClient) (data *ApiClientCreateResult, err error) {
	body, err := JSONReader(client)
	if err != nil {
		return
	}
	err = e.client.magicRequestDecoder("POST", "clients", body, &data)
	return
}

// Delete removes a client on the Chef server
//
// Chef API docs: https://docs.chef.io/api_chef_server.html#clients-name
// TODO doc says name and validator flag are returned
func (e *ApiClientService) Delete(name string) (err error) {
	url := fmt.Sprintf("clients/%s", name)
	err = e.client.magicRequestDecoder("DELETE", url, nil, nil)
	return
}

// Get gets a client from the Chef server.
//
// Chef API docs: https://docs.chef.io/api_chef_server.html#clients-name
func (e *ApiClientService) Get(name string) (client ApiClient, err error) {
	url := fmt.Sprintf("clients/%s", name)
	err = e.client.magicRequestDecoder("GET", url, nil, &client)
	return
}

// Put updates a client on the Chef server.
//
// Chef API docs: https://docs.chef.io/api_chef_server.html#clients-name
// TODO: Add a go test
// TODO: Update the other go tests
func (e *ApiClientService) Update(name string, client ApiNewClient) (data *ApiClientCreateResult, err error) {
	body, err := JSONReader(client)
	url := fmt.Sprintf("clients/%s", name)
	if err != nil {
		return
	}
	err = e.client.magicRequestDecoder("PUT", url, body, &data)
	return
}

// ListKeys lists the keys associated with a client on the Chef server.
//
// Chef API docs: https://docs.chef.io/api_chef_server.html#clients-client-keys
func (e *ApiClientService) ListKeys(clientName string) (data *ApiClientKeyListResult, err error) {
	url := fmt.Sprintf("clients/%s/keys", clientName)
	err = e.client.magicRequestDecoder("GET", url, nil, &data)
	return
}

// POST /client/CLIENT/keys TODO

// DELETE /client/CLIENT/keys TODO
// TODO fix the orger of the client/keys doc in chef api

// PUT /client/CLIENT/keys TODO

// GetKey gets a client key from the Chef server.
//
// Chef API docs: https://docs.chef.io/api_chef_server.html#clients-client-keys-key
func (e *ApiClientService) GetKey(clientName string, keyName string) (data *ApiClientKey, err error) {
	url := fmt.Sprintf("clients/%s/keys/%s", clientName, keyName)
	err = e.client.magicRequestDecoder("GET", url, nil, &data)
	return
}
