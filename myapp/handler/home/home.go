package home

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"myapp/database/praiseData"
	"myapp/database/userData"
)

type Data struct {
	Users        []UserWithSelected
	TargetUserId string
	Praise       praiseData.Praise
	Error        string
}

type UserWithSelected struct {
	User     userData.User
	Selected bool
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {

	data := Data{
		Users:        getUserWithSelectedSlice(userData.GetAllUsers(), 0),
		TargetUserId: "",
	}
	renderTemplate(w, "home", &data)
}

func DisplayPraiseHandler(w http.ResponseWriter, r *http.Request) {
	// validate
	targetUserId, _ := strconv.Atoi(r.FormValue("targetUserId"))
	data := Data{
		Users:        getUserWithSelectedSlice(userData.GetAllUsers(), targetUserId),
		TargetUserId: r.FormValue("targetUserId"),
	}

	praise := praiseData.GetPraiseRandomly(targetUserId)

	log.Println(praise)
	data.Praise = praise
	renderTemplate(w, "home", &data)
}

func getUserWithSelectedSlice(users []userData.User, selectedId int) []UserWithSelected {
	var result []UserWithSelected
	for _, v := range users {
		value := UserWithSelected{
			User:     v,
			Selected: false,
		}
		if selectedId == v.ID {
			value.Selected = true
		}
		result = append(result, value)
	}
	return result
}

func renderTemplate(w http.ResponseWriter, tmpl string, d *Data) { // TODO interfaceで共通化したい
	t, _ := template.ParseFiles("template/" + tmpl + ".html")
	err := t.Execute(w, d)
	if err != nil {
		panic(err)
	}
}
