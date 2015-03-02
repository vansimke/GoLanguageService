package main

import (
	"net/http"
	"io/ioutil"
	"tokenizer"
	"encoding/json"
)

func main() {
	http.HandleFunc("/tokenize", func(w http.ResponseWriter, req *http.Request) {
			data, _ := ioutil.ReadAll(req.Body)
			tokens := tokenizer.Tokenize(string(data))
			result, _ := json.Marshal(tokens)
			w.Header().Add("Content Type", "application/json")
			w.Write(result)
		});
	
	http.ListenAndServe(":3030", nil)
}

