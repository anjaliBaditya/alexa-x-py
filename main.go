package main

import (
    "encoding/json"
    "log"
    "math/rand"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)




type Item struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

var items []Item


func createItem(w http.ResponseWriter, r *http.Request) {
    var item Item
    _ = json.NewDecoder(r.Body).Decode(&item)
    item.ID = strconv.Itoa(rand.Intn(1000000))
    items = append(items, item)
    json.NewEncoder(w).Encode(item)
}


func getItems(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(items)
}
func getItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range items {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.Error(w, "Item not found", http.StatusNotFound)
}



func updateItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range items {
        if item.ID == params["id"] {
            items = append(items[:index], items[index+1:]...)
            var updatedItem Item
            _ = json.NewDecoder(r.Body).Decode(&updatedItem)
            updatedItem.ID = params["id"]
            items = append(items, updatedItem)
            json.NewEncoder(w).Encode(updatedItem)
            return
        }
    }
    http.Error(w, "Item not found", http.StatusNotFound)
}
func deleteItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range items {
        if item.ID == params["id"] {
            items = append(items[:index], items[index+1:]...)
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.Error(w, "Item not found", http.StatusNotFound)
}

func main() {
    router := mux.NewRouter()

    items = append(items, Item{ID: "1", Name: "Item One", Price: 10.99})
    items = append(items, Item{ID: "2", Name: "Item Two", Price: 20.99})

    router.HandleFunc("/items", createItem).Methods("POST")
    router.HandleFunc("/items", getItems).Methods("GET")
    router.HandleFunc("/items/{id}", getItem).Methods("GET")
    router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
    router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":8000", router))
}
