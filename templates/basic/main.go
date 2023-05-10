package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "one.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}

	// -------------------------------------------------------------

	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad"}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}

	// -------------------------------------------------------------

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief is no beliefs",
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", buddha)
	if err != nil {
		log.Fatalln(err)
	}

	// -------------------------------------------------------------

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	sagesSlice := []sage{buddha, gandhi, mlk, jesus, muhammad}
	err = tpl.ExecuteTemplate(os.Stdout, "four.gohtml", sagesSlice)
	if err != nil {
		log.Fatalln(err)
	}

	// -------------------------------------------------------------

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sagesSlice,
		cars,
	}

	err = tpl.ExecuteTemplate(os.Stdout, "five.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

	// -------------------------------------------------------------
	err = tpl.ExecuteTemplate(os.Stdout, "six.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
