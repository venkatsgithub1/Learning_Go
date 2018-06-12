package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// create a new gorilla router
	r := mux.NewRouter()

	// route to home.
	r.HandleFunc("/", HomeHandler)

	// multiple routes under a main route
	// main route: authentication.
	authRouter := r.PathPrefix("/authentication/").Subrouter()
	// /authentication/ login or logout or singup.
	authRouter.Path("/login").HandlerFunc(LoginHandler)
	authRouter.Path("/signup").HandlerFunc(SignupHandler)
	authRouter.Path("/logout").HandlerFunc(LogoutHandler)

	// subroute on posts.
	postsRouter := r.Path("/posts").Subrouter()
	postsRouter.Methods("GET").HandlerFunc(PostsIndexHandler)
	postsRouter.Methods("POST").HandlerFunc(PostsCreateHandler)

	// on posts
	// particular post by id:
	post := r.PathPrefix("/posts/{id}").Subrouter()
	// call a different handler on get request.
	post.Methods("GET").Path("/edit").HandlerFunc(EditHandler)
	// just show all the posts.
	post.Methods("GET").HandlerFunc(ShowPostsHandler)
	post.Methods("PUT", "POST").HandlerFunc(PostUpdateHandler)
	post.Methods("DELETE").HandlerFunc(PostDeleteHandler)

	http.ListenAndServe(":8000", r)
}

// PostDeleteHandler function is used to delete posts.
func PostDeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete posts")
}

// PostUpdateHandler function is used to updated a post.
func PostUpdateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update posts")
}

// ShowPostsHandler function is used for displaying posts.
func ShowPostsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "showing posts")
}

// EditHandler function is called to edit posts.
func EditHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edit post")
}

// PostsIndexHandler function is called to check posts.
func PostsIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Posts index")
}

// PostsCreateHandler function is called to check posts.
func PostsCreateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Posts create")
}

// LoginHandler function opens login page.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login")
}

// LogoutHandler function opens logout page.
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "logout")
}

// SignupHandler function opens signup page.
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Singup")
}

// HomeHandler function prints response on page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page")
}
