package service

// There are a number of assertions in this test:
// ■ We receive a 200 from the /skus/{sku} resource.
// ■ We can parse the body of that response.
// ■ The body of that response can be converted into a fulfillmentStatus struct.
// ■ The details of that response are the fakes that we expect, since we’re not building a fully
// functioning service.

// "encoding/json"
// "io/ioutil"
// "net/http"
// "net/http/httptest"
// "github.com/codegangsta/negroni"
// "github.com/gorilla/mux"
// "github.com/unrolled/render"

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/unrolled/render"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

func TestGetFullfilmentStatusReturns200ForExistingSKU(t *testing.T) {
	var (
		request  *http.Request
		recorder *httptest.ResponseRecorder
	)

	server := NewServer()

	targetSKU := "THINGAMAJIG12"

	recorder = httptest.NewRecorder()
	// Make a request with get method to /skus/targetSKU
	request, _ = http.NewRequest("GET", "/skus/"+targetSKU, nil)

	server.ServeHTTP(recorder, request)

	var detail fulfillmentStatus

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}

	// Le o conteudo da resposta registada
	payload, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	// Filtrar de certa forma o conteudo da resposta com o conteudo da struct
	// Recebe como 2o argumento o endereco de detail e altera o seu valor com
	// a informacao filtrada de payload
	err = json.Unmarshal(payload, &detail)
	if err != nil {
		t.Errorf("Error unmarshaling response to fulfillment status: %v", err)
	}

	// Testa condicoes da ja alterada variavel detail
	if detail.QuantityInStock != 1000 {
		t.Errorf("Expected 1000 qty in stock, got %d", detail.QuantityInStock)
	}
	if detail.ShipsWithin != 99 {
		t.Errorf("Expected shipswithin 14 days, got %d", detail.ShipsWithin)
	}
	if detail.SKU != "THINGAMAJIG12" {
		t.Errorf("Expected SKU THINGAMAJIG12, got %s", detail.SKU)
	}

}
