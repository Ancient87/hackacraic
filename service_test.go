package main

import (
    "testing"
    "encoding/json"
)

const token1 string = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ilg1ZVhrNHh5b2pORnVtMWtsMll0djhkbE5QNC1jNTdkTzZRR1RWQndhTmsifQ.eyJleHAiOjE1MjM3NDc5NDEsIm5iZiI6MTUyMzc0NDM0MSwidmVyIjoiMS4wIiwiaXNzIjoiaHR0cHM6Ly9sb2dpbi5taWNyb3NvZnRvbmxpbmUuY29tL2Q1Zjg1NjgyLWY2N2EtNDQ0NC05MzY5LTJjNWVjMWEwZThjZC92Mi4wLyIsInN1YiI6Ijc3ZDQxOWNjLWZhMmItNDlmMS04NTJmLWUzMjBhMTNkZmI5NCIsImF1ZCI6IjQwOTU3YjljLTYzYmMtNGFiNC05ZWNiLTY3YjU0M2M4ZTRjYSIsIm5vbmNlIjoiZGVmYXVsdE5vbmNlIiwiaWF0IjoxNTIzNzQ0MzQxLCJhdXRoX3RpbWUiOjE1MjM3NDQzNDEsIm9pZCI6Ijc3ZDQxOWNjLWZhMmItNDlmMS04NTJmLWUzMjBhMTNkZmI5NCIsIm5hbWUiOiJTZXZlcmluIiwiZmFtaWx5X25hbWUiOiJHbHV0ZXhvIiwiZ2l2ZW5fbmFtZSI6IlNldiIsImVtYWlscyI6WyJzZXZlcmluLmdhc3NhdWVyQGdtYWlsLmNvbSJdLCJ0ZnAiOiJCMkNfMV9CbHVlQmFua1NVU0kifQ.ns7cT7dUwXUkPPOcdybFbCZU8rLb9l4ajT2lnB1DmMe45EUgGNbGFZ6lcDLyk1sIfK68O1G97049kWeEQPIbaMKR1GlR9A5H6axa5X0jD6UM7_FPzu1FbXPQaQr8m8btULuVBXGFs7CdmIrEtxxQu_HCfi03VJtQKBIBeU60RUqNczPRp3yhOeFyEc0Kl-ECYehXgKvsH4wso45-R4epYYR2RBZGztWpQmkFW6aw24h4On8Vzf8adOGYOLYDG7VuMkqIjOrxXFhtewlnH0TRxbp46_sYFAc9NVbSggwZQZEG8EkK9E8X2ZX7QUSulRT18cmtuI-zVYND1wG7LIL3Sw"
//const token2 string = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ilg1ZVhrNHh5b2pORnVtMWtsMll0djhkbE5QNC1jNTdkTzZRR1RWQndhTmsifQ.eyJleHAiOjE1MjM3MDYyNDQsIm5iZiI6MTUyMzcwMjY0NCwidmVyIjoiMS4wIiwiaXNzIjoiaHR0cHM6Ly9sb2dpbi5taWNyb3NvZnRvbmxpbmUuY29tL2Q1Zjg1NjgyLWY2N2EtNDQ0NC05MzY5LTJjNWVjMWEwZThjZC92Mi4wLyIsInN1YiI6Ijc3ZDQxOWNjLWZhMmItNDlmMS04NTJmLWUzMjBhMTNkZmI5NCIsImF1ZCI6IjQwOTU3YjljLTYzYmMtNGFiNC05ZWNiLTY3YjU0M2M4ZTRjYSIsIm5vbmNlIjoiZGVmYXVsdE5vbmNlIiwiaWF0IjoxNTIzNzAyNjQ0LCJhdXRoX3RpbWUiOjE1MjM3MDI2NDQsIm9pZCI6Ijc3ZDQxOWNjLWZhMmItNDlmMS04NTJmLWUzMjBhMTNkZmI5NCIsIm5hbWUiOiJTZXZlcmluIiwiZmFtaWx5X25hbWUiOiJHbHV0ZXhvIiwiZ2l2ZW5fbmFtZSI6IlNldiIsImVtYWlscyI6WyJzZXZlcmluLmdhc3NhdWVyQGdtYWlsLmNvbSJdLCJ0ZnAiOiJCMkNfMV9CbHVlQmFua1NVU0kifQ.cP0vRlA_Pk4U2VATc7ffffhtB4otZs7JZydFWNr95tlGO8eBAgQ9cldWg23nm8zVy5uUTl2e-aJytriHVm5Ti4PL_ZS3LMgo0y6zNitNy7wXuxpkkwbGDDQIHGMm4-_WsnphhgQV8CaNddxWF0JgjCT_e5gpy4CksCQcjfJAQ_945qZXSWWTY33wijYG_81Kto5TL0I_HpGJWlhRauo-TxWGPlpJ6_Q735hjKKvshDNSdwnKZK4DPfSyZ5-eXE36XCiITLGUQu-bZWg5wAj-Lqzfx0Y-hyoo-L3Iec-YdlRwNX_yAodPoTrd1M5DoFG4ekn6gMFjGK1mXikBLvJ77g"

func TestCustomer(t *testing.T) {

	// Make my two test customers
	cust1 := MakeCustomer("Charmander", token1, 100)
	//v1 := vendors[0]
	//cust2 := MakeCustomer("Bulbasaur", token1, 1000)

	/*
    customers := Customer{
		cust1,
		cust2,
	}
	*/
    // Test we can get the customer logged in
    custID := cust1.getCustomerID()
    t.Logf("Customer ID acc 1 %s", custID)

    // Test we can get the accounts for the customerID
    accs := cust1.GetAccounts(custID)
    t.Logf("Customer accounts for %s: %s", custID, accs)
	//facc := (*accs)[0]
    //aid := facc.Cid
	// Get the balance for the last year
	//ly := cust1.GetNetLastYearForAccount(aid)
	//t.Logf("Customer balance for last year %d", ly)
	// Test judgment
	//testJudgment(t, v1, cust1)

    // Test the Category
    testCategory(t, c1, cust1)
}

func testJudgment(t *testing.T, v *Vendor, c *Customer) {
    // Judge customer by vendor
    judgement := v.Judge(c)
    t.Logf("Judgement for %s against %s : %s", c.Name, v.Name, judgement)
}

func testCategory(t *testing.T, cat *Category, c *Customer) {
    cj := cat.JudgeCategory(c)
    marshalled, _ := json.MarshalIndent(cj, "", "    ")
    t.Logf("Marshalled json %s", marshalled)
}


/*
func testGetPokemonByName(t *testing.T, s *pokeService, names []string) *Pokemons {
    ps := Pokemons{}
    for _, name := range names {
        p := s.getPokemonByName(name)
        t.Logf("I got myself a %s", p)
        if p.Name != name {
            t.Errorf(" Expected %s but got %s", name, p.Name)
        }
        ps = append(ps, &p)
    }
    t.Logf("Gotta catch them all %s", ps)
    return &ps
}
*/

