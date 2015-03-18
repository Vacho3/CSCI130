package main

import (
	"text/template"
	"os"
)

type Person struct {
	Title string
	LName string
	DidAttend bool
	DidDonate bool
	Events []string
}

func main() {

	const mailtemplate = `Dear {{.Title}} {{.LName}},
{{if .DidAttend}}Thanks for coming to the fundraiser.
{{else}}We missed you at the fundraiser.
{{end}}{{if .DidDonate}}We would like to thank you for the generous donation.
{{end}}Reminder of upcoming events:
{{range .Events }}  - {{.}}
{{end}}Sincerely,
Vachagan Vardanyan

`
	var events = []string{	
		"Some Big Huge Event (June 5th, 2015)",
		"My Birthday Party (Bring Gifts) (June 1th, 2015)",
		"Another Fundraiser (Bring Money) (July 7th, 2015)",
	}

	var people = []Person{
		{"Mrs.", "Tyson", 	true, 	true, 	events},
		{"Ms.",  "Jason", 	true, 	false, 	events},
		{"Mr.",  "Vic",     false, 	false, 	events},
		{"Mr.",  "Shawn", 	false, 	true, 	events},
	}

	tmpl := template.Must(template.New("mailtemplate").Parse(mailtemplate))
	for _,person := range people {
		tmpl.Execute(os.Stdout, person)
	}

}