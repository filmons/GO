package handlers

import (
    "encoding/json"
	"log"
    "net/http"
    "github.com/filmons/chiapi/db"
    // "github.com/filmons/chiapi/models"
)



//  GetUser fetches a user by ID
// func GetUser(w http.ResponseWriter, r *http.Request) {
//     // Example: Fetch a user using db.DB
//     var user models.User
//     // Assume you have a function in db to get a user by ID
//     err := db.GetUserByID(&user)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     json.NewEncoder(w).Encode(user)
// }
// GetUser fetches a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
    // Extract user ID from request parameters, URL path, or elsewhere
    userID := "123" // Replace with actual user ID retrieval logic

    // Example: Fetch a user using db.DB
    user, err := db.GetUserByID(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(user)
}



func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request received: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
