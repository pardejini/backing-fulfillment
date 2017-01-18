package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// getFulfillmentStatusHandler simulates actual fulfillment by supplying
// bogus values for QuantityInStock and ShipsWithin any give SKU.
// Used to demonstrate a backing service supporting a primary service
func getFulfillmentStatusHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		sku := vars["sku"]

		// Transforma o objecto de resposta(w) em JSON e envia para o cliente
		// com StatusOK

		formatter.JSON(w, http.StatusOK, fulfillmentStatus{
			SKU:             sku,
			ShipsWithin:     99,
			QuantityInStock: 1000,
		})
	}
}

func rootHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.Text(w, http.StatusOK, "Fulfillment Service, see http://github.com/cloudnativego/backing-fulfillment for API.")
	}
}
