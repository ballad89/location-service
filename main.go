package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/oschwald/geoip2-golang"
	"github.com/pariz/gountries"
)

func getOrDefault(key, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}

func main() {

	dbLocation := getOrDefault("GEO_DB", "./GeoLite2-Country.mmdb")

	log.Println("Using db file at", dbLocation)

	db, err := geoip2.Open(dbLocation)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := gountries.New()

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {

		clientIpKey, ok := r.URL.Query()["client_ip"]

		if !ok || len(clientIpKey) < 1 {
			log.Println("Url Param 'client_ip' is missing")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Url Param 'client_ip' is missing"))
			return
		}

		clientIP := clientIpKey[0]

		if clientIP != "" {
			ip := net.ParseIP(clientIP)

			record, err := db.Country(ip)

			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Ip not a valid public ip"))
				return
			}

			country := record.Country.Names["en"]

			if country == "" {
				log.Println("Ip did not return a country")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Ip did not return a country"))
				return
			}
			log.Println("Found country", country)

			count, err := query.FindCountryByName(country)

			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Country details not found in db"))
				return
			}

			res, err := json.Marshal(count)

			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			w.Write(res)

		}
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
		return
	})

	port := getOrDefault("PORT", ":1989")

	log.Println("Starting up on ...", port)

	err = http.ListenAndServe(port, nil)

	if err != nil {
		panic(err)
	}
}
