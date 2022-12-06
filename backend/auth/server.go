package main

import (
// standard library
	"net/http"
	"context"
// 3rd party library
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/gorilla/mux"
// my modules
	"authservice/modules/auth"
)

func main() {
	r := mux.NewRouter()
	db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to gorm-sqlite3 database")
	}
	r.Use(InstallDatabaseMiddleware(db))

	r.HandleFunc("/api/auth/users", auth.UserListHandler)
	r.HandleFunc("/api/auth/login", auth.UserLoginHandler)
	r.HandleFunc("/api/auth/verify-token", auth.VerifyTokenHandler)
	
	fmt.Println("Running at localhost:8000")
	http.ListenAndServe(":8000", r)
}

func InstallDatabaseMiddleware(db *gorm.DB) mux.MiddlewareFunc {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			ctx := context.WithValue(context.Background(), "database", db)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}