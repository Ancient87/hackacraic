package main

import (
    "fmt"
    "net/http"
    "time"
    "io/ioutil"
    "log"
    "encoding/json"
)

var customerIDCounter int = 1

const apimKey string = "03fafdb0a21247c78c59898904a9ed79"
const apiLocation string = "https://bluebank.azure-api.net/api/v0.7"

// Interface OpenBanakingCustomer can get from the API`
type OpenBankingCustomer interface {
    GetAnnualTurnOver() int
    getTurnOver(days int) int
    GetBalance() int
    GetTransactionHistoryDays() int
}

// Something to parse our transaction into based on JSON
/*
type Transaction struct {
   
}
*/

type Account struct {
    ID int `json: "id"`
}

// Customer struct
type Customer struct {
    ID int `json: "id"`
    Name string `json: "name"`
    Requester *Requester
    Lending int `json: "lending"`
}

type customerID struct {
    ID int `json: "id"`
}

type Requester struct {
	bearer string
}


// Constructor for customers
func MakeCustomer(name string, token string, lending int) *Customer {
    reqHeader := makeRequester(token)
    cust := &Customer {
        ID: customerIDCounter,
        Name: name,
        Requester: reqHeader,
        Lending: lending,
    }
	customerIDCounter += 1
    return cust
}
// Get all accounts for the customer
func (c *Customer) getAccounts() *[]Account {
    accounts := []Account{}
    endpoint := fmt.Sprintf("%s/accounts", apiLocation)
    body := c.makeAPICall(endpoint)
    fmt.Printf("%s", body)
    jsonErr := json.Unmarshal(body, &accounts)
    if jsonErr != nil {
        log.Fatal(jsonErr)
	}
	return &accounts
}

// Get all Customer IDs authorized for
func (c* Customer) getCustomerID() int {
    var dat map[string]interface{}
    endpoint := fmt.Sprintf("%s/customers", apiLocation)
    body := c.makeAPICall(endpoint)
    fmt.Printf("The RESPONSE %s \n\n\n", body)
    jsonErr := json.Unmarshal(body, &dat)
    result := dat["results"].(map[string] interface{})
    fmt.Println(result)
     if jsonErr != nil {
        log.Fatal(jsonErr)
	}
    return 0
}

// Make APICall and return json
func (c *Customer) makeAPICall(endpoint string) []byte {
	// create the request
	req, err := c.Requester.newRequest(endpoint)
	client := http.Client{
        Timeout: time.Second *10,
    }
    if err != nil {
        log.Fatal(err)
	}

	res, getErr := client.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }
    return body
}

// Makes a new requester setting the customers bearer
func makeRequester(token string) *Requester {
	return &Requester{
		bearer: token,
	}
}

// Returns a request with bearer set for given endpoint
func (r *Requester) newRequest(endpoint string) (*http.Request, error) {
    req, err := http.NewRequest("GET", endpoint, nil)
    if err != nil {
        return nil, err
    }
	bearer := fmt.Sprintf("Bearer %s", r.bearer)
    req.Header.Add("Authorization", bearer)
	req.Header.Add("Ocp-Apim-Subscription-Key", apimKey)
    req.Header.Set("User-Agent", "hackacraic")
	return req, nil
}
