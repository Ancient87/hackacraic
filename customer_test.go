package main

import "testing"

const token1 string = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ilg1ZVhrNHh5b2pORnVtMWtsMll0djhkbE5QNC1jNTdkTzZRR1RWQndhTmsifQ.eyJleHAiOjE1MjM3MTQzNzUsIm5iZiI6MTUyMzcxMDc3NSwidmVyIjoiMS4wIiwiaXNzIjoiaHR0cHM6Ly9sb2dpbi5taWNyb3NvZnRvbmxpbmUuY29tL2Q1Zjg1NjgyLWY2N2EtNDQ0NC05MzY5LTJjNWVjMWEwZThjZC92Mi4wLyIsInN1YiI6IjIyMzJjOWMyLTJmYzktNDU3OS04YzJlLWE2OTYyMDc5MWUwNiIsImF1ZCI6IjQwOTU3YjljLTYzYmMtNGFiNC05ZWNiLTY3YjU0M2M4ZTRjYSIsIm5vbmNlIjoiZGVmYXVsdE5vbmNlIiwiaWF0IjoxNTIzNzEwNzc1LCJhdXRoX3RpbWUiOjE1MjM3MTA3NzUsIm9pZCI6IjIyMzJjOWMyLTJmYzktNDU3OS04YzJlLWE2OTYyMDc5MWUwNiIsIm5hbWUiOiJHcmVnb3J5VENhcnJvbGwiLCJmYW1pbHlfbmFtZSI6IkNhcnJvbGwiLCJnaXZlbl9uYW1lIjoiR3JlZ29yeSIsImVtYWlscyI6WyJncmVnc2NhcnJvbGxAZ21haWwuY29tIl0sInRmcCI6IkIyQ18xX0JsdWVCYW5rU1VTSSJ9.ISQQU8xgamyBRykYT2xQ1smhrzE111-62x7ENIqi4JfpizcaP1K7Cdt3pezC7p1HfN7bQv-ah39BbYv6AUr1XAkjz-nk03MnlvXRV5e_QsdpqmKnrU_0RPPxXGAJkkBwDPeMISuxG8xwlftqp6GyT7lZ51eVO5AdDO2S9V3niEojICc_JNUsHIsG_jeOdL8ra9Lo8i55yrIA4zAbDWetz2zD4rBkF6fsIg3OSNM66QijEejAhg2SGhSZT1irhSxpjXOsPHWZNPF2q6Eef6ywA7XF0mFPgaio-K83EyeDZmlUYMnHVWRIGpIIYVg9yw17Cq8uckD0XlRdXg1nKMK6pQ"
const token2 string = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6Ilg1ZVhrNHh5b2pORnVtMWtsMll0djhkbE5QNC1jNTdkTzZRR1RWQndhTmsifQ.eyJleHAiOjE1MjM3MDYyNDQsIm5iZiI6MTUyMzcwMjY0NCwidmVyIjoiMS4wIiwiaXNzIjoiaHR0cHM6Ly9sb2dpbi5taWNyb3NvZnRvbmxpbmUuY29tL2Q1Zjg1NjgyLWY2N2EtNDQ0NC05MzY5LTJjNWVjMWEwZThjZC92Mi4wLyIsInN1YiI6Ijc3ZDQxOWNjLWZhMmItNDlmMS04NTJmLWUzMjBhMTNkZmI5NCIsImF1ZCI6IjQwOTU3YjljLTYzYmMtNGFiNC05ZWNiLTY3YjU0M2M4ZTRjYSIsIm5vbmNlIjoiZGVmYXVsdE5vbmNlIiwiaWF0IjoxNTIzNzAyNjQ0LCJhdXRoX3RpbWUiOjE1MjM3MDI2NDQsIm9pZCI6Ijc3ZDQxOWNjLWZhMmItNDlmMS04NTJmLWUzMjBhMTNkZmI5NCIsIm5hbWUiOiJTZXZlcmluIiwiZmFtaWx5X25hbWUiOiJHbHV0ZXhvIiwiZ2l2ZW5fbmFtZSI6IlNldiIsImVtYWlscyI6WyJzZXZlcmluLmdhc3NhdWVyQGdtYWlsLmNvbSJdLCJ0ZnAiOiJCMkNfMV9CbHVlQmFua1NVU0kifQ.cP0vRlA_Pk4U2VATc7ffffhtB4otZs7JZydFWNr95tlGO8eBAgQ9cldWg23nm8zVy5uUTl2e-aJytriHVm5Ti4PL_ZS3LMgo0y6zNitNy7wXuxpkkwbGDDQIHGMm4-_WsnphhgQV8CaNddxWF0JgjCT_e5gpy4CksCQcjfJAQ_945qZXSWWTY33wijYG_81Kto5TL0I_HpGJWlhRauo-TxWGPlpJ6_Q735hjKKvshDNSdwnKZK4DPfSyZ5-eXE36XCiITLGUQu-bZWg5wAj-Lqzfx0Y-hyoo-L3Iec-YdlRwNX_yAodPoTrd1M5DoFG4ekn6gMFjGK1mXikBLvJ77g"

func TestCustomer(t *testing.T) {

	// Make my two test customers
	cust1 := MakeCustomer("Charmander", token1, 100)
	cust2 := MakeCustomer("Bulbasaur", token2, 1000)

	/*
    customers := Customer{
		cust1,
		cust2,
	}
	*/
	
    // Test we can get the customer logged in
    custID := cust1.getCustomerID()
    t.Logf("Customer ID acc 1 %s", custID)
    custID = cust2.getCustomerID()
    t.Logf("Customer ID acc 2 %s", custID)
    // Test we can get all accounts
	/*
    accounts := cust1.getAccounts()
	t.Logf("Account for 1 %s", accounts)
	accounts = cust2.getAccounts()
	t.Logf("Account for 2 %s", accounts)
    */
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

