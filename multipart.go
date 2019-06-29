package snippets

import (
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
)

func getMultipartValues() (ret [][]byte) {
	// TODO: just placeholder
	return
}

// handler
func handleMultipart(w http.ResponseWriter, r *http.Request) {
	mediatype, _, err := mime.ParseMediaType(r.Header.Get("Accept"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	if mediatype != "multipart/form-data" {
		http.Error(w, "set Accept: multipart/form-data", http.StatusMultipleChoices)
		return
	}

	mw := multipart.NewWriter(w)
	w.Header().Set("Content-Type", mw.FormDataContentType())

	for _, value := range getMultipartValues() {
		fw, err := mw.CreateFormField("value")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if _, err := fw.Write(value); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	if err := mw.Close(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// client side
func requestgetMultipart() {
	req, _ := http.NewRequest("GET", "http://localhost:8080/foo", nil)
	req.Header.Set("Accept", "multipart/form-data; charset=utf-8")
	resp, _ := http.DefaultClient.Do(req)

	_, params, _ := mime.ParseMediaType(resp.Header.Get("Content-Type"))

	mr := multipart.NewReader(resp.Body, params["boundary"])

	for {
		part, err := mr.NextPart()
		if err != nil {
			break
		}
		value, _ := ioutil.ReadAll(part)
		log.Printf("Value: %s", value)
	}
}
