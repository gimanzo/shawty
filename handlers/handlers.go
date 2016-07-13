package handlers

import (
	"net/http"
	"github.com/gimanzo/shawty/storages"
	"strconv"
	"github.com/speps/go-hashids"
	"fmt"
)

func EncodeHandler(storage storages.IStorage, encoder *hashids.HashID) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if url := r.PostFormValue("url"); url != "" {
			id := storage.Save(url)
			intId, _ := strconv.Atoi(id)
			intArray := []int{intId}
			encoded, _ := encoder.Encode(intArray)
			w.Write([]byte(encoded))
		}
	}

	return http.HandlerFunc(handleFunc)
}

func DecodeHandler(storage storages.IStorage, encoder *hashids.HashID) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		url, err := Decode(r.URL.Path[len("/decode/"):], encoder, storage)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("URL Not Found"))
			return
		}
		w.Write([]byte(url))
	}
	return http.HandlerFunc(handleFunc)
}

func RedirectHandler(storage storages.IStorage, encoder *hashids.HashID) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		hash := r.URL.Path[len("/"):]
		url, err := Decode(hash, encoder, storage)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("URL Not Found"))
			return
		}
		http.Redirect(w, r, string(url), 301)
	}
	return http.HandlerFunc(handleFunc)
}

func Decode(hash string, encoder *hashids.HashID, storage storages.IStorage) (string, error) {
	code, _ := encoder.DecodeInt64WithError(hash)
	codeString := strconv.FormatUint(uint64(code[0]), 10)
	url, err := storage.Load(codeString)
	return url, err
}
