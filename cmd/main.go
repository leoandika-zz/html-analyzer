package main

import (
	"HTMLAnalyzer/model"
	"HTMLAnalyzer/service"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var htmlAnalyzerService service.Service

func checkHTML(w http.ResponseWriter, r *http.Request) {
	var response model.Response
	url, ok := r.URL.Query()["url"]
	if !ok || len(url[0]) < 1 {
		log.Println("Request Param 'url' is missing")
		return
	}

	err := validateURL(url[0])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	analyzeResult, err := htmlAnalyzerService.CheckHTMLFromURL(url[0])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response = model.Response{
		HTMLVersion:            analyzeResult.HTMLVersion,
		PageTitle:              analyzeResult.PageTitle,
		HeadingCount:           analyzeResult.HeadingCount,
		InternalLinksCount:     analyzeResult.InternalLinkCount,
		ExternalLinksCount:     analyzeResult.ExternalLinkCount,
		InaccessibleLinksCount: analyzeResult.InaccessibleLinkCount,
		LoginFormExist:         analyzeResult.LoginFormExist,
	}

	respondWithJSON(w, http.StatusOK, response)
}

func main() {

	router := mux.NewRouter()
	http.Handle("/", router)

	client := http.Client{}
	htmlAnalyzerService = service.NewHTMLAnalyzerService(&client)

	router.HandleFunc("/analyzehtml", checkHTML).Methods("GET")

	fmt.Println("Connected to port 8087")
	log.Fatal(http.ListenAndServe(":8087", router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc.Encode(payload)
}

func validateURL(url string) error {
	if url[:12] != "https://www." && url[:11] != "http://www." {
		return errors.New("invalid url. Please use https:// or http:// and www")
	}
	return nil
}
