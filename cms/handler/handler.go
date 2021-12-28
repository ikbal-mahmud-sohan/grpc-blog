package handler

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"

	bgvc "myGo/Blogs/gunk/v1/category"
	bgv "myGo/Blogs/gunk/v1/post"
)

type Handler struct {
	templates *template.Template
	decoder   *schema.Decoder
	sess      *sessions.CookieStore
	tc        bgv.PostServiceClient
	cc        bgvc.CategoryServiceClient
}


func New(decoder *schema.Decoder, sess *sessions.CookieStore,
	tc bgv.PostServiceClient, cc bgvc.CategoryServiceClient) *mux.Router {
	h := &Handler{
		decoder: decoder,
		sess:    sess,
		tc:      tc,
		cc:      cc,
	}

	h.parseTemplate()
	r := mux.NewRouter()
	r.HandleFunc("/", h.Index)
	r.HandleFunc("/categories", h.IndexCategory)
	s := r.NewRoute().Subrouter()
	//post
	s.HandleFunc("/posts/create", h.createPost)
	s.HandleFunc("/posts/store", h.storePost)
	s.HandleFunc("/posts/{id:[0-9]+}/edit", h.editPost)
	r.HandleFunc("/singlePost/{id:[0-9]+}/view", h.SinglePage)

	s.HandleFunc("/posts/{id:[0-9]+}/update", h.updatePost)
	s.HandleFunc("/posts/{id:[0-9]+}/delete", h.deletePost)
	//catCategory
	s.HandleFunc("/categories/create", h.createCategory)
	s.HandleFunc("/categories/store", h.storeCategory)
	s.HandleFunc("/categories/{id:[0-9]+}/edit", h.editCategory)
	s.HandleFunc("/categories/{id:[0-9]+}/update", h.updateCategory)
	s.HandleFunc("/categories/{id:[0-9]+}/delete", h.deleteCategory)

	s.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("./"))))
	s.Use(h.middleWare)

	r.NotFoundHandler = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := h.templates.ExecuteTemplate(rw, "404.html", nil); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	return r
}
func (h *Handler) middleWare(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(rw, r)
	})
}
func (h *Handler) parseTemplate() {
	h.templates = template.Must(template.ParseFiles(
		"cms/assets/templates/posts/createPost.html",
		"cms/assets/templates/posts/indexPost.html",
		"cms/assets/templates/posts/editPost.html",
		"cms/assets/templates/posts/SinglePost.html",
		"cms/assets/templates/categories/createCategory.html",
		"cms/assets/templates/categories/indexCategory.html",
		"cms/assets/templates/categories/editCategory.html",
		"cms/assets/templates/404.html",
	))
}
