package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type response struct {
	Text string `json:"text"`
}

func Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{Text: "This is fine."})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler)

	tlsConf := &tls.Config{
		// all tls versions below 1.2 are insecure and deprecated now
		MinVersion: tls.VersionTLS12,
		// use server ciphersuites listed below
		PreferServerCipherSuites: true,
		// disable weak algorithms such as DSA/DSS/MD5/SHA1
		CipherSuites: []uint16{
			// {key-exchange-algorithm}_{public-key-algorithm}_WITH_
			// {symmetric-encryption-algorithm}_{key-bits}_{pseudo-random-func-MAC}
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		},
	}

	srv := &http.Server{
		Addr:         ":443",
		Handler:      mux,
		TLSConfig:    tlsConf,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	fmt.Println("server started at localhost:443...")
	err := srv.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		panic(err)
	}
}
