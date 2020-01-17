package main

import (
	"FullStack-Challenge/challengeApi/Boot"
	AuthorsRepo "FullStack-Challenge/challengeApi/Domains/authors/entity/repository/mongodb"
	AuthorsInterface "FullStack-Challenge/challengeApi/Domains/authors/interface"
	AuthorsUsecase "FullStack-Challenge/challengeApi/Domains/authors/usecase"
	"FullStack-Challenge/challengeApi/presenters"

	BooksRepo "FullStack-Challenge/challengeApi/Domains/books/entity/repository/mongodb"
	BooksInterface "FullStack-Challenge/challengeApi/Domains/books/interface"
	BooksUsecase "FullStack-Challenge/challengeApi/Domains/books/usecase"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var startTime time.Time
var db *mgo.Database

func main() {

	db = MongoStart()
	router := loadRouter()
	Boot.Boot(db.Session)
	log.Fatal(http.ListenAndServe(":8000", router))

}

func MongoStart() *mgo.Database {
	session, err :=
		mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("challengeDB")

	return db

}

func loadRouter() *mux.Router {

	router := mux.NewRouter()

	// Repos
	authorsRepo := AuthorsRepo.NewMongoDBRepository(db.Session)
	booksRepo := BooksRepo.NewMongoDBRepository(db.Session)

	// Usecases
	authorsUsecase := AuthorsUsecase.NewUsecase(authorsRepo)
	booksUsecase := BooksUsecase.NewUsecase(booksRepo)

	// Interfaces
	authorsInterface := AuthorsInterface.AuthorInterface(authorsUsecase)
	booksInterface := BooksInterface.BooksInterface(booksUsecase)

	//Authors Router
	router.HandleFunc("/authors", authorsInterface.Create).Methods("POST")
	router.HandleFunc("/authors/{id}", authorsInterface.GetOne).Methods("GET")
	router.HandleFunc("/authors", authorsInterface.GetAll).Methods("GET")
	router.HandleFunc("/authors/{id}", authorsInterface.Delete).Methods("DELETE")
	router.HandleFunc("/authors", authorsInterface.Update).Methods("PATCH")

	//Books Router
	router.HandleFunc("/books", booksInterface.Create).Methods("POST")
	router.HandleFunc("/books/{id}", booksInterface.GetOne).Methods("GET")
	router.HandleFunc("/books", booksInterface.GetAll).Methods("GET")
	router.HandleFunc("/books/{id}", booksInterface.Delete).Methods("DELETE")
	router.HandleFunc("/books", booksInterface.Update).Methods("PATCH")

	//Other Routes ================================
	router.HandleFunc("/", HealthCheckHandler)

	return router
}

//Check the status and uptime of the server
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	lifeTime := bson.M{"started": startTime.Format("2006-01-02T15:04:05.999999-07:00"), "uptime": fmt.Sprint(math.Floor(time.Since(startTime).Seconds()*1000) / 1000)}
	presenters.ReturnHttpPayload(lifeTime, w)
}

//Empty Response for OPTIONS Request
func OptionsResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "limit, includes, skip, sort, filter, select, self, Accept, Accept-Language, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Methods", "PATCH, POST, GET, OPTIONS, PUT, DELETE")

	w.WriteHeader(200)
}
