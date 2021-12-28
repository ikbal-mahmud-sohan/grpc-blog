package handler

import (
	"fmt"
	"log"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"

	bgvc "myGo/Blogs/gunk/v1/category"
)

type IndexCategory struct{
	Category Category
}

func (h *Handler) IndexCategory (rw http.ResponseWriter, r *http.Request) {
fmt.Println("done")

	res,err:= h.cc.ListCategory(r.Context(), &bgvc.ListCategoryRequest{})
	if err!=nil{
		log.Fatal(err)
	}
	if err:= h.templates.ExecuteTemplate(rw,"indexCategory.html", res); err !=nil{
		http.Error(rw, err.Error(),http.StatusInternalServerError)
		return
	}
}

func (h *Handler) deleteCategory (rw http.ResponseWriter, r *http.Request) {
vars := mux.Vars(r)
	id := vars["id"]
	
	if id == "" {
		http.Error(rw, "invalid update", http.StatusTemporaryRedirect)
		return
	}
		Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	_,err = h.cc.DeleteCategory(r.Context(),&bgvc.DeleteCategoryRequest{
		ID: Id,
	})
	if err!=nil{
		log.Fatal(err)
	}
	http.Redirect(rw,r, "/categories", http.StatusTemporaryRedirect)
}
