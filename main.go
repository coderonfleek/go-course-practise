package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tmpl, _ = template.ParseGlob("templates/*.html")

// Property names of any struct that will be accessible in templates must start with a capital letter
// This is our Go distinguish property access i.e public vs private properties
type PageData struct {
	PageTitle string
	PageBody  string
}

type Employee struct {
	Name string
	Sex  string
	Role string
}

func main() {

	//Parse Templates

	/* if err != nil {
		panic(err)
	} */

	gRouter := mux.NewRouter()

	gRouter.HandleFunc("/", Homepage)
	gRouter.HandleFunc("/about", AboutPage)
	gRouter.HandleFunc("/services", ServicesPage)
	gRouter.HandleFunc("/team", Teampage)

	//fmt.Println("Writing Go now o")

	http.ListenAndServe(":3000", gRouter)
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome to the Homepage")
	tmpl.ExecuteTemplate(w, "home.html", nil)
}

func AboutPage(w http.ResponseWriter, r *http.Request) {

	tmpl.ExecuteTemplate(w, "about.html", &PageData{
		"About Page",
		"When you know, you know. Go on fleek",
	})
}

func ServicesPage(w http.ResponseWriter, r *http.Request) {

	data := struct {
		PageInfo    PageData
		ServiceList []string
	}{
		PageInfo: PageData{
			"Services Page",
			"Here is a List of Services we offer",
		},
		ServiceList: []string{"Web Development", "Graphic Design", "Branding", "DevOps"},
	}

	tmpl.ExecuteTemplate(w, "services.html", data)

}

func Teampage(w http.ResponseWriter, r *http.Request) {

	data := struct {
		PageInfo PageData
		TeamList []Employee
	}{
		PageInfo: PageData{
			"Team Page",
			"Here is our staff",
		},
		TeamList: []Employee{
			{
				Name: "Fikayo Adepoju",
				Sex:  "male",
				Role: "Educator",
			},
			{
				Name: "Kemi Adepoju",
				Sex:  "female",
				Role: "CFO",
			},
			{
				Name: "Abolade Adepoju",
				Sex:  "male",
				Role: "Producer",
			},
		},
	}

	tmpl.ExecuteTemplate(w, "team.html", data)

}
