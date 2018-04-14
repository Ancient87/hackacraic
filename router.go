package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strings"
	"encoding/json"
	"strconv"
)

type Route struct {
	name string
	path string
	f SevHandler
}

var routes = []*Route{
	&Route{name:"Dash", path:"/", f: renderDash},
	&Route{name:"Category", path:"/category/{id}", f:renderCategory},
}

func NewRouter() *mux.Router {
	fmt.Println("Building a beautiful router")
    r := mux.NewRouter()
	for _, route := range routes {
		path := route.path
		fun := route.f
		r.HandleFunc(path, func(wr http.ResponseWriter, req *http.Request) {
			render(wr, req, r, fun)
        }).Methods("GET")
	}
	_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})
	return r
}

type SevHandler func(wr http.ResponseWriter, req *http.Request, r *mux.Router, c *Customer)

// Wrapper which builds the user object for the query then gets ds and returns marshalled json
func render(wr http.ResponseWriter, req *http.Request, r *mux.Router, fun SevHandler ) {
	v := req.URL.Query()
	bearer := v.Get("bearer")
	//id := vars["id"]
	// make the user
	cust := MakeCustomer("A customer", bearer, 10)
	fun(wr, req, r, cust)
	// Serve JSON
}

// Serve the json needed for the dashboard
func renderDash(wr http.ResponseWriter, req *http.Request, r *mux.Router, c *Customer) {
	// Cust Result
	dbv := new(DashBoardView)
	cr := c.getCustomerInfo()
	cats := new([]*CategoryView)
	// Loop through all categories and build the struct
	for _, cat := range categories {
		cv  := new(CategoryView)
		cv.Name = cat.Name
		var vs = make([]string, 0, 20)
		for _, v := range cat.Vendors {
			vs = append(vs, v.Name)
		}
		cv.Vendors = vs
		*cats = append(*cats, cv)
	}
	dbv.Info = cr
	dbv.Categories = *cats
	js, _ := json.MarshalIndent(dbv, "","    ")
	wr.Write(js)
}

type DashBoardView struct {
	Info *CustomerResult `json:"info"`
	Categories []*CategoryView `json:"categories"`
}

// Contains all we need to render our Categories for the Dash
type CategoryView struct {
	Name string `json:"name"`
	Path	string `json:"path"`
	Vendors []string `json:"vendors"`
}

// Serve the category judgment extracting bearer and and id
func renderCategory(wr http.ResponseWriter, req *http.Request, r *mux.Router, c *Customer) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	// Get the category
	cat := categories[id]
	// Get the category judgment for a user	
	j := cat.JudgeCategory(c)
	js, _ := json.MarshalIndent(j, "","    ")
	wr.Write(js)
}
