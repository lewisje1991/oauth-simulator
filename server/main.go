package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func main() {

	http.HandleFunc("POST /token", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request /token")
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error parsing form", err)
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		clientId := r.Form.Get("client_id")
		clientSecret := r.Form.Get("client_secret")
		grantType := r.Form.Get("grant_type")

		if clientId == "" || clientSecret == "" || grantType == "" {
			fmt.Println("Invalid request")
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		privateKey, err := loadECDSAPrivateKey("private.ec.key")
		if err != nil {
			fmt.Println("Error loading private key", err)
			http.Error(w, "Error loading private key", http.StatusInternalServerError)
			return
		}

		t := jwt.New(jwt.SigningMethodES256)
		s, err := t.SignedString(privateKey)
		if err != nil {
			fmt.Println("Error signing token", err)
			http.Error(w, "Error signing token", http.StatusInternalServerError)
			return
		}

		response := struct {
			AccessToken string `json:"access_token"`
			Expiry      int64  `json:"expires_in"`
			TokenType   string `json:"token_type"`
		}{
			AccessToken: s,
			Expiry:      30,
			TokenType:   "Bearer",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	})

	http.HandleFunc("POST /token-echo", func(w http.ResponseWriter, r *http.Request) {
		//print the headers
		for name, values := range r.Header {
			for _, value := range values {
				if name == "Authorization" {
					fmt.Println(name, value)
				}
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func loadECDSAPrivateKey(path string) (*ecdsa.PrivateKey, error) {
	keyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}
