package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type get_documents_input struct {
	TourIds []int
}

func get_documents_handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var input get_documents_input
	err := decoder.Decode(&input)
	if err != nil {
		panic(err)
	}
	var t2 []string
	for _, i := range input.TourIds {
		j := strconv.Itoa(i)
		t2 = append(t2, j)
	}

	ids := strings.Join(t2, ", ")
	reqBody := fmt.Sprintf("{\"ids\":[%s]}", ids)
	log.Println(reqBody)

	client := &http.Client{}
	url := "https://product-search.gygdev.gygtest.com/gyg_activity_en_live/tour/_mget"
	req, _ := http.NewRequest("GET", url, bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "%s", body)
}

func main() {
	http.HandleFunc("/v1/documents", get_documents_handler)
	http.ListenAndServe(":8080", nil)
}
