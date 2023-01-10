package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Producto struct {
	Id          string `json:id`
	Nombre      string `json:nombre`
	Descripcion string `json:descripcion`
	Cantidad    int    `json:cantidad`
}

//Array global de productos:
var Productos []Producto

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h2>Servidor funcionando!</h2></body></html>")
	fmt.Println("Solicitud atendida: homePage")
}

func findAllProductos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: findAllProductos")
	json.NewEncoder(w).Encode(Productos)
}

func findProductoById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: findProductoById")
	vars := mux.Vars(r)
	key := vars["id"]
	for _, producto := range Productos {
		if producto.Id == key {
			json.NewEncoder(w).Encode(producto)
		}
	}
}
func createNewProducto(w http.ResponseWriter, r *http.Request) {
	// Se obtiene el body desde el request y
	// se deserializa en una variable producto:
	reqBody, _ := ioutil.ReadAll(r.Body)
	var producto Producto
	json.Unmarshal(reqBody, &producto)
	// adicionamos en el array el nuevo producto:
	Productos = append(Productos, producto)
	json.NewEncoder(w).Encode(producto)
}

func deleteProducto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: deleteProducto")
	vars := mux.Vars(r)
	key := vars["id"]
	// buscar el producto a eliminar:
	for index, producto := range Productos {
		if producto.Id == key {
			// borrar del array:
			Productos = append(Productos[:index], Productos[index+1:]...)
		}
	}
}

func updateProducto(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Solicitud atendida: updateProducto")
	// Se obtiene el body desde el request y
	// se deserializa en una variable producto:
	reqBody, _ := ioutil.ReadAll(r.Body)
	var producto Producto
	json.Unmarshal(reqBody, &producto)
	key := producto.Id
	// buscar el producto a actualizar:
	for index, p := range Productos {

		if p.Id == key {
			fmt.Println("", producto)
			// actualizar el array:
			Productos[index] = producto
			break
		}
	}
	json.NewEncoder(w).Encode(producto)
}

func iniciarServidor() {
	fmt.Println("API REST simple con lenguaje go.")
	ruteador := mux.NewRouter().StrictSlash(true)
	ruteador.HandleFunc("/", homePage)
	ruteador.HandleFunc("/productos", findAllProductos)
	////el orden de definicion es importante en el manejo de rutas:
	ruteador.HandleFunc("/producto", createNewProducto).Methods("POST")
	ruteador.HandleFunc("/producto/{id}", deleteProducto).Methods("DELETE")
	ruteador.HandleFunc("/producto", updateProducto).Methods("PUT")
	ruteador.HandleFunc("/producto/{id}", findProductoById)
	log.Fatal(http.ListenAndServe(":8080", ruteador))
}
func main() {
	Productos = []Producto{
		Producto{Id: "1", Nombre: "Monitor 17 pulgadas", Descripcion: "Conexión	HDMI - Full HD", Cantidad: 3},
		Producto{Id: "2", Nombre: "Teclado USB", Descripcion: "Color negro y teclas multimedia", Cantidad: 7},
	}
	iniciarServidor()
}
