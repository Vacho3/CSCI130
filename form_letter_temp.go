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
Alexander Plummer

`
	var upcomingEvents = []string{	
		"Some Big Huge Event (June 5th, 2015)",
		"My Birthday Party (Bring Gifts) (June 8th, 2015)",
		"Another Fundraiser (Bring Money) (July 15th, 2015)",
	}

	var people = []Person{
		{"Mrs.", "Knope", 	true, 	true, 	upcomingEvents},
		{"Ms.",  "Perkins", 	false, 	false, 	upcomingEvents},
		{"Mr.",  "Swanson", 	true, 	false, 	upcomingEvents},
		{"Mr.",  "Haverford", 	false, 	true, 	upcomingEvents},
	}

	tmpl := template.Must(template.New("mailtemplate").Parse(mailtemplate))
	for _,person := range people {
		tmpl.Execute(os.Stdout, person)
	}

}