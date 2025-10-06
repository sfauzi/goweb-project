package categorycontroller

import (
	"goweb-project/models/categorymodel"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {

}
func Edit(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}
