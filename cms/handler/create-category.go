package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/mux"

	bgvc "myGo/Blogs/gunk/v1/category"
)

type formCategoryData struct {
	Category Category
	Errors   map[string]string
}

type Category struct {
	ID    int64
	Title string
}

func (c *Category) validate() error {

	return validation.ValidateStruct(c,
		validation.Field(&c.Title,
			validation.Required.Error("This filed cannot be null"),
			validation.Length(3, 30).Error("The Post name length must be between 3 and 30"),
		),
	)
}

func (h *Handler) createCategory(rw http.ResponseWriter, r *http.Request) {
	Category := Category{}
	Errors := map[string]string{}

	h.loadCreatedCategoryForm(rw, Category, Errors)
}
func (h *Handler) storeCategory(rw http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	var categories Category

	if err := h.decoder.Decode(&categories, r.PostForm); err != nil {

		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
fmt.Println(categories)
	if err := categories.validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[strings.Title(key)] = value.Error()

			}
			h.loadCreatedCategoryForm(rw, categories, vErrs)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return

	}
	_, err := h.cc.CreateCategory(r.Context(), &bgvc.CreateCategoryRequest{
		Category: &bgvc.Category{
			Title: categories.Title,
		},
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(rw, r, "/categories", http.StatusTemporaryRedirect)

}

func (h *Handler) loadCreatedCategoryForm(rw http.ResponseWriter, categories Category, errs map[string]string) {

	form := formCategoryData{
		Category: categories,
		Errors:   errs,
	}
	if err := h.templates.ExecuteTemplate(rw, "createCategory.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *Handler) editCategory(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(rw, "invalid ", http.StatusTemporaryRedirect)
		return
	}

	Id, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	Errors := map[string]string{}

	res, err := h.cc.GetCategory(r.Context(), &bgvc.GetCategoryRequest{
		ID: Id,
	})
	if err != nil {
		log.Fatal(err)
	}
	Category := Category{
		ID:    res.Category.ID,
		Title: res.Category.Title,
	}

	h.loadUpdatedCategoryForm(rw, Category, Errors)
}

func (h *Handler) updateCategory(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(rw, "invalid update", http.StatusTemporaryRedirect)
		return
	}

	Id, err := strconv.ParseInt(id, 10, 64)
	fmt.Println(Id)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	var categories Category

	if err := h.decoder.Decode(&categories, r.PostForm); err != nil {

		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(categories)

	if err := categories.validate(); err != nil {
		vErrors, ok := err.(validation.Errors)
		if ok {
			vErrs := make(map[string]string)
			for key, value := range vErrors {
				vErrs[strings.Title(key)] = value.Error()

			}

			h.loadUpdatedCategoryForm(rw, categories, vErrs)
			return
		}
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return

	}

	_, err = h.cc.UpdateCategory(r.Context(), &bgvc.UpdateCategoryRequest{
		Category: &bgvc.Category{
			ID:    Id,
			Title: categories.Title,
		},
	})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(rw, r, "/categories", http.StatusTemporaryRedirect)
}

func (h *Handler) loadUpdatedCategoryForm(rw http.ResponseWriter, categories Category, errs map[string]string) {

	form := formCategoryData{
		Category: categories,
		Errors:   errs,
	}
	if err := h.templates.ExecuteTemplate(rw, "editCategory.html", form); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}
