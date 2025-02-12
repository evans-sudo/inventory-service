package product

import "net/http"



const productsPath = "products"


func handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productList := 
	}
}