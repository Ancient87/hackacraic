package main

import (
    //"fmt"
)


// Vendor struct
type Vendor struct {
    Id int `json:"id"`
    Name string `json:"name"`
    MinTurnOver int `json:"minturnover"`
    MaxTurnOver int `json:"maxturnover"`
    ExistingLending int `json:"maxlending"`
    TransactionHistoryMin int `json:"minhistory"`
	Checks []VendorCheck `json:"checks"`
}


// Test Vendor
var Ulster = &Vendor {
Id: 1,
Name: "Ulster Bank",
MinTurnOver: 1000,
MaxTurnOver: 1000000,
TransactionHistoryMin: 1*365,
ExistingLending: 500,
Checks: checks,
}


// Test Vendor
var Santander = &Vendor {
Id: 2,
Name: "Santander",
MinTurnOver: 0,
MaxTurnOver: 10000000000,
TransactionHistoryMin: 3*365,
ExistingLending: 1000,
Checks: checks,
}

// Test Vendor
var LendingWorks = &Vendor {
Id: 3,
Name: "Lending Works",
MinTurnOver: 1000,
MaxTurnOver: 1000000,
TransactionHistoryMin: 1*365,
ExistingLending: 20,
Checks: checks,
}


// Test Vendor
var ArchOver = &Vendor {
Id: 4,
Name: "Arch Over",
MinTurnOver: 1000,
MaxTurnOver: 5000000,
TransactionHistoryMin:3*365,
ExistingLending: 100,
Checks: checks,
}

// Define an array of ponters to vendors
type Vendors []*Vendor

// Set of our test vendors
var vendors = []*Vendor{Ulster, Santander, LendingWorks, ArchOver}


// Suitability Judgement by Vendor returned by Vendors
type VendorJudgement struct {
	Name string `json:"vendorname"`
	CheckTotal int `json:"totalChecks"`
	CheckPassed int `json:"checksPassed"`
	CheckResults []CheckResult `json:"testResults"`
}


// vendors satisfy the Check interface run through all checks and pass judgement
func (v *Vendor) Judge(c *Customer) *VendorJudgement {
	// Create Judgement
	// Total checks run
	tc := 0
	// Checks passed
	tp := 0
	j := &VendorJudgement {
		Name: v.Name,
	}
	// Loop through all checks and see where we are
	for _, check := range v.Checks {
		// Create Check result struct
		cr := &CheckResult{}
		// Copy Check Name into Judgement Name
		cr.CheckName = check.Name
		// Get the check functiion
		cf := check.C
		// Run the check function store result is and should
		result, is, should := cf(c, v)
		cr.Is = is
		cr.Should = should
		if result {
			// Record test success
			tp += 1
		}
		cr.Pass = result
		// Increase number of checks performed
		tc += 1
		// Add checkResult to the Judgement
		j.CheckResults = append(j.CheckResults, *cr)
	}
	// Write final check metrics
	j.CheckTotal = tc
	j.CheckPassed = tp
	return j
}

// The Result of a check that can be rendered by the UI
type CheckResult struct {
	CheckName	string	`json:"checkname"`
	Pass		bool	`json:"pass"`
	Is			int		`json:"is"`
	Should		int		`json:"should"`
}

// Checks expose the check method which take a struct of type customer and vendor and pass judgement Returns result is and should
type Judger interface {
    Judge(c *Customer) *VendorJudgement
}

type Check func (c *Customer, v *Vendor) (bool, int, int)

// Holds check interface and gives it a name
type VendorCheck struct {
    Name string
    C Check
}

// Static list of all our current checks - hacky AF
var checks = []VendorCheck{
	VendorCheck { Name: "Annual TurnOver" , C: AnnualTurnOver },
	VendorCheck { Name: "Max Lending", C: MaxExistingLending },
}

// Function of type Judge which compares customer turnover for a year to vendor min
func AnnualTurnOver (c *Customer, v *Vendor) (bool, int, int) {
	//is := customer.turnoverForDays(365)
	is := c.GetNetLastYear()
	should := v.MinTurnOver
    return is >= should, is, should
}

func MaxExistingLending (c *Customer, v *Vendor) (bool, int, int) {
	is := c.Lending
	should := v.ExistingLending
	return is < should, is, should
}


