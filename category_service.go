package main

import (

)

// Our categories
var c1 = &Category{
    Name: "Lending",
    Vendors:vendors,
}

var c2 = &Category{
    Name: "Foreign Exchange",
    Vendors:vendors,
}

var c3 = &Category{
    Name: "Auditing",
    Vendors:vendors,
}

// Out categories - for now
var categories = []*Category{c1, c2, c3}

// Represents a category
type Category struct {
    Name    string  `json:"name"`
    Vendors []*Vendor `json:"vendors"`
}

// Represents a Judgement Category ready to render
type CategoryJudgement struct {
    Name    string `json:"name"`
    Judgements   []*VendorJudgement `json:"judgements"`
}

// Returns all the cateogies we have
func GetCategories() []*Category{
    return categories
}

// Judges the category and returns the category to render
func (cat *Category) JudgeCategory(c *Customer) *CategoryJudgement {
    // Get Judgements for all vendors against this customer
    cj := new(CategoryJudgement)
    cj.Name = cat.Name
    for _, v := range cat.Vendors {
        j := v.Judge(c)
        cj.Judgements = append(cj.Judgements, j)
    }
    return cj
}

// Returns marshalled JSON
//func GetJson() 
