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

type CustomerResults struct {
    //M []Meta `json:"meta"`
	Results []CustomerResult `json:"results"`
}

type TransactionResults struct {
	Results []TransactionResult `json:"results"`
}

type TransactionResult struct {
	Id string `json:"id"`
	Description string `json:"transactionDescription"`
	Amount float64 `json:"transactionAmount"`
}

type Meta struct {
	total int `json:"total"`
	count int `json:"count"`
}

// Represent details of customers that are useful
type CustomerResult struct {
    Id string `json:"id"`
    Name string `json:"givenName"`
}

// Struct to hold customer account API results
type CustomerAccountResults struct {
	Results []CustomerAccount `json:"results"`
}
// Struct to hold a customer account
type CustomerAccount struct {
	Cid string `json:"id"`
	Number string `json:"accountNumber"`
	Name string `json:"accountFriendlyName"`
	Balance float64 `json:"accountBalance"`
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
// Get all accounts for the customerID
func (c *Customer) GetAccounts(cid string) *[]CustomerAccount {
    var accounts = new(CustomerAccountResults)
    endpoint := fmt.Sprintf("%s/customers/%s/accounts", apiLocation, cid)
    //fmt.Println("Endpoint queries for GetAccounts %s", endpoint)
	body := c.makeAPICall(endpoint)
    fmt.Printf("Get accounts response %s \n", body)
    jsonErr := json.Unmarshal(body, &accounts)
    if jsonErr != nil {
        log.Fatal(jsonErr)
	}
	fmt.Println("Account Result:", accounts)
	//accs := accounts.Results
	//fmt.Println(accs[0])
	return &accounts.Results
}

// Get all Customer IDs authorized for
func (c* Customer) getCustomerID() string {
    var cr = new(CustomerResults)
    endpoint := fmt.Sprintf("%s/customers", apiLocation)
    body := c.makeAPICall(endpoint)
    fmt.Printf("The RESPONSE %s \n\n\n", body)
    jsonErr := json.Unmarshal(body, &cr)
    if jsonErr != nil {
        log.Fatal(jsonErr)
	}
	fmt.Println("Customer Result ",cr)
    return cr.Results[0].Id
}


// Get Transactions for last year
func (c *Customer) GetTransactionsLastYear(account string) []TransactionResult {
	var tr = new(TransactionResults)
	endpoint := fmt.Sprintf("%s/accounts/%s/transactions/?filter=transactionDateTime.gt.2017-04-01T00:00:00.000Z", apiLocation, account)
	body := c.makeAPICall(endpoint)
	fmt.Printf("URL endpoint for GetTransactionsLastYear %s", endpoint)
    //fmt.Printf("The Transactions since 2017 today: %s\n\n\n", body)
    jsonErr := json.Unmarshal(body, &tr)
    if jsonErr != nil {
        log.Fatal(jsonErr)
	}
	//fmt.Println("Customer Result ",tr)
	return tr.Results
}

// Get net result for last year
func (c *Customer) GetNetLastYearForAccount(account string) int {
	tr := c.GetTransactionsLastYear(account)
	net := 0.0
	// Sum transactions of last year
	for _, t := range tr {
		amount := t.Amount
		//desc := t.Description
		//id := t.Id
		net += amount
	}
	return int(net)
}

// Get Net Last Year for all accounts
func (c *Customer) GetNetLastYear() int {
	totalNet := 0
	// Gest customer ID
	cid := c.getCustomerID()
	// Get all accounts for me with my ID
	accs := c.GetAccounts(cid)
	accsP := *accs
	for _, acc := range accsP {
		yearNet := c.GetNetLastYearForAccount(acc.Cid)
		totalNet += yearNet
	}
	// Test 
	return totalNet
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
