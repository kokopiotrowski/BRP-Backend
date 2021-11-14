/*
 * StockX - Bachelor Project
 *
 * This API is developed for a Bachelor Project Application named StockX.   - It both projects current stock market data and works as a learning tool while doing fake stock trading.    Authors of the project - Aleksander Stefan Bialik - Konrad Piotrowski  
 *
 * API version: 1.0.0
 * Contact: 280053@via.dk
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"log"
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	sw "./go"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
