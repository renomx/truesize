package main_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"."
)

var (
	server   *httptest.Server
	reader   io.Reader //Ignore this for now
	shoesUrl string
)

var a main.App

func init() {

	a = main.App{}

	a.SetConfig()

	a.Initialize(
		a.Config.Tests.Host,
		a.Config.Tests.DbPort,
		a.Config.Tests.User,
		a.Config.Tests.Password,
		a.Config.Tests.DbName)

	server = httptest.NewServer(a.Router)
}

func TestCreateShoe(t *testing.T) {

	shoesUrl = fmt.Sprintf("%s/shoe", server.URL)

	shoeJson := `{"name": "adidas Yeezy", 
    "sizes": [1, 2, 2, 3, 2, 3, 2, 2, 3, 4, 2, 5, 2, 3]}`

	reader = strings.NewReader(shoeJson)

	request, err := http.NewRequest("POST", shoesUrl, reader)

	res, err := http.DefaultClient.Do(request)

	bodyBytes, err := ioutil.ReadAll(res.Body)

	shoe := a.View
	json.Unmarshal(bodyBytes, &shoe)

	if shoe == nil {
		t.Errorf("It is expected that shoe has a value")
	}

	if shoe.Name != "adidas Yeezy" {
		t.Errorf("Expected shoe name is adidas Yeezy %s", shoe.Name)
	}

	if len(shoe.Sizes) == 0 {
		t.Errorf("The number of sizes is expected to be greater than zero, got %d", len(shoe.Sizes))
	}

	// TODO: Test TrueToSize value

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP Code expected: 200, got %d", res.StatusCode)
	}
}

func TestGetShoe(t *testing.T) {
	shoesUrl = fmt.Sprintf("%s/shoe/adidas Yeezy", server.URL)

	request, err := http.NewRequest("GET", shoesUrl, nil)

	res, err := http.DefaultClient.Do(request)

	bodyBytes, err := ioutil.ReadAll(res.Body)

	shoe := a.View
	json.Unmarshal(bodyBytes, &shoe)

	if shoe.Name != "adidas Yeezy" {
		t.Errorf("Expected shoe name is adidas Yeezy %s", shoe.Name)
	}

	if len(shoe.Sizes) == 0 {
		t.Errorf("The number of sizes is expected to be greater than zero, got %d", len(shoe.Sizes))
	}

	if res.StatusCode != 200 {
		t.Errorf("HTTP Code expected: 200, got %d", res.StatusCode)
	}

	if err != nil {
		t.Errorf("Oops something went wrong %s", err)
	}

}

func TestAddTrueToSizeValue(t *testing.T) {
	shoesUrl = fmt.Sprintf("%s/shoe/truetosize/adidas Yeezy", server.URL)

	trueToSizeJson := `{"size": 4}`

	reader = strings.NewReader(trueToSizeJson)

	request, err := http.NewRequest("PUT", shoesUrl, reader)

	res, err := http.DefaultClient.Do(request)

	bodyBytes, err := ioutil.ReadAll(res.Body)

	shoe := a.View
	json.Unmarshal(bodyBytes, &shoe)

	if res.StatusCode != 200 {
		t.Errorf("HTTP Code expected: 200, got %d", res.StatusCode)
	}

	if err != nil {
		t.Errorf("Oops something went wrong %s", err)
	}

}
