package approvePraise

import (
	"fmt"
	"html/template"
	"log"
	"myapp/database/praiseData"
	"net/http"
	"strconv"
)

type Data struct {
	Praises           []praiseData.Praise `validate:"-"`
	ApprovedPraiseIds []string            `validate:"-"`
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	data := Data{
		Praises: praiseData.GetAllPraises(false),
	}

	renderTemplate(w, "approval", &data)
}

func ApproveHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("approve praise start")
	// TODO validate
	r.ParseForm()
	fmt.Printf("%+v\n", r.Form)
	approvedPraiseIds := r.Form["approvedPraiseIds"]

	log.Println(approvedPraiseIds)
	for _, v := range approvedPraiseIds {
		var intVal int
		intVal, _ = strconv.Atoi(v)
		praiseData.ApprovePraises(intVal)
	}

	log.Println("approve praise end")
	http.Redirect(w, r, "/praise/approve", http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, d *Data) {
	t, _ := template.ParseFiles("template/" + tmpl + ".html")
	err := t.Execute(w, d)
	if err != nil {
		panic(err)
	}
}
