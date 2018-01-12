package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/oschwald/geoip2-golang"
	"github.com/pariz/gountries"
)

func main() {

	dbLocation := os.Getenv("GEO_DB")

	db, err := geoip2.Open(dbLocation)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := gountries.New()

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {

		keys, ok := r.URL.Query()["client_ip"]

		if !ok || len(keys) < 1 {
			println("Url Param 'key' is missing")
		}

		clientIP := keys[0]
		fmt.Println(clientIP)

		if clientIP != "" {
			ip := net.ParseIP(clientIP)

			record, err := db.Country(ip)

			if err != nil {
				println(err)
			}

			fmt.Println(record.Country)

			country := record.Country.Names["en"]

			fmt.Println(country)

			count, err := query.FindCountryByName(country)

			if err != nil {

			}

			res, err := json.Marshal(count)

			if err != nil {

			}

			w.Write(res)

		}
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		return
	})

	err = http.ListenAndServe(":1989", nil)

	if err != nil {
		panic(err)
	}
}
