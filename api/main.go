package main

import (
	"context"
	"encoding/json"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Post struct {
	Title     string
	CreatedAt time.Time
}

var posts []Post

func homeHandler(w http.ResponseWriter, r *http.Request) {
	responseJSON(w, posts)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Read post
	post := &Post{}
	err = json.Unmarshal(data, post)
	if err != nil {
		responseError(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.CreatedAt = time.Now().UTC()

	cleanPost(post)

	// Don't append, we only need one title at the moment
	//posts = append(posts, post)
	posts = []Post{*post}

	responseJSON(w, post)
}

func cleanPost(p *Post) *Post {
	p.Title = html.EscapeString(p.Title)
	return p
}

func createInitialPost() {
	post := Post{ // b == Student{"Bob", 0}
		Title:     "dynamic string",
		CreatedAt: time.Now().UTC(),
	}

	// Don't append, we only need one title at the moment
	//posts = append(posts, post)
	posts = []Post{post}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	createInitialPost()

	var wait time.Duration

	r := mux.NewRouter()
	r.HandleFunc("/posts", homeHandler).Methods("GET")
	r.HandleFunc("/posts", updateHandler).Methods("POST")

	cor := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081", "http://localhost:8080", "http://string-storer-lb-tf-663390881.eu-west-2.elb.amazonaws.com:80"},
		//AllowCredentials: true,
	})

	handler := cor.Handler(r)

	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      logRequest(handler), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func responseError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	enc.Encode(data)
}
