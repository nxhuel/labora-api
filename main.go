package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
    ID int `json:ID`
    Name string `json:Name`
}

var items = []Item {
    {
        ID: 1,
        Name: "Item one",
    },
}


func getItems(w http.ResponseWriter, r *http.Request) {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page < 1 {
        page = 1
    }

    itemsPerPage, _ := strconv.Atoi(r.URL.Query().Get("itemsPerPage"))
    if itemsPerPage < 1 {
        itemsPerPage = 10
    }

    start := (page - 1) * itemsPerPage

    end := start + itemsPerPage

    if start > len(items) {
        start = len(items)
    }

    if end > len(items) {
        end = len(items)
    }

    // Envio item en formato json
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items[start:end])
}

func getItem(w http.ResponseWriter, r *http.Request) {

}

func createItem(w http.ResponseWriter, r *http.Request) {
    var newItem Item
    // Informacion que el cliente le envia al servidor
    reqBody, err := ioutil.ReadAll(r.Body)

    // Control de error
    if err != nil {
        fmt.Fprintf(w, "Inserte datos validos")
    }

    // Asigno el reqBody(o sea la informacion que estoy recibiendo) al nuevo item "newItem"
    json.Unmarshal(reqBody, &newItem)

    // Se limita el uso a 10 items maximo
    if len(items) < 50 {
        // Se suma uno al ID c/v que se cree un item
        newItem.ID = len(items) + 1
        items = append(items, newItem)
    
        // Le mostramos al usuario su tipo de dato
        w.Header().Set("Content-Type", "application/json")

        // Le mostramos al usuario si su tipo de dato fue correcto
        w.WriteHeader(http.StatusCreated)

        // Le enviamos el dato creado
        json.NewEncoder(w).Encode(newItem)
    } else {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Te sobrepasaste de items, el maximo es 10")
    }
}

func updateItem(w http.ResponseWriter, r *http.Request) {

}

func deleteItem(w http.ResponseWriter, r *http.Request) {

}

func indexRoute(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Bienvenido a mi primera labora-API")
}

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/", indexRoute)

    router.HandleFunc("/items", getItems).Methods("GET")
    router.HandleFunc("/items/{id}", getItem).Methods("GET")
    router.HandleFunc("/items", createItem).Methods("POST")
    router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
    router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

    http.ListenAndServe(":8000", router)
}