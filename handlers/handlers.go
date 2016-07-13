package handlers

import (
	"net/http"
	"github.com/gimanzo/shawty/storages"
	"github.com/gimanzo/shawty/analytics"
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
			hash, _ := encoder.Encode(intArray)
			w.Write([]byte(hash))
			analytics.Log(analytics.CategoryEncode, url, hash, analytics.StatusSuccess)
		}
	}

	return http.HandlerFunc(handleFunc)
}

func DecodeHandler(storage storages.IStorage, encoder *hashids.HashID) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		hash := r.URL.Path[len("/decode/"):]
		url, err := Decode(hash, encoder, storage)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			analytics.Log(analytics.CategoryDecode, "", hash, analytics.StatusMiss)
			w.Write([]byte("URL Not Found"))
			return
		}
		analytics.Log(analytics.CategoryDecode, url, hash, analytics.StatusHit)
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
			analytics.Log(analytics.CategoryRedirect, "", hash, analytics.StatusMiss)
			w.Write([]byte("URL Not Found"))
			return
		}
		analytics.Log(analytics.CategoryRedirect, url, hash, analytics.StatusSuccess)
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
