package handlers

import (
	"net/http"
	"github.com/gimanzo/shawty/storages"
	"github.com/gimanzo/shawty/analytics"
	"strconv"
	"github.com/speps/go-hashids"
	"fmt"
	"os"
	"errors"
)

func EncodeHandler(storage storages.IStorage, encoder *hashids.HashID) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if url := r.PostFormValue("url"); url != "" {
			id := storage.Save(url)
			intId, _ := strconv.Atoi(id)
			intArray := []int{intId}
			hash, _ := encoder.Encode(intArray)

			addVersion(w)
			w.Write([]byte(hash))
			analytics.Log(analytics.CategoryEncode, url, hash, analytics.StatusSuccess, r)
		}
	}

	return http.HandlerFunc(handleFunc)
}
func addVersion(writer http.ResponseWriter) {
	writer.Header().Set("x-version", os.Getenv("APP_VERSION"))
}

func DecodeHandler(storage storages.IStorage, encoder *hashids.HashID) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		hash := r.URL.Path[len("/decode/"):]
		url, err := Decode(hash, encoder, storage)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			analytics.Log(analytics.CategoryDecode, "", hash, analytics.StatusMiss, r)
			addVersion(w)
			w.Write([]byte("URL Not Found"))
			return
		}
		analytics.Log(analytics.CategoryDecode, url, hash, analytics.StatusHit, r)
		addVersion(w)
		w.Write([]byte(url))
	}
	return http.HandlerFunc(handleFunc)
}

func RedirectHandler(storage storages.IStorage, encoder *hashids.HashID) http.Handler {
	handleFunc := func(w http.ResponseWriter, request *http.Request) {
		hash := request.URL.Path[len("/"):]
		url, err := Decode(hash, encoder, storage)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusNotFound)
			analytics.Log(analytics.CategoryRedirect, "", hash, analytics.StatusMiss, request)
			addVersion(w)
			w.Write([]byte("URL Not Found"))
			return
		}
		analytics.Log(analytics.CategoryRedirect, url, hash, analytics.StatusSuccess, request)
		addVersion(w)
		http.Redirect(w, request, string(url), 301)
	}
	return http.HandlerFunc(handleFunc)
}

func Decode(hash string, encoder *hashids.HashID, storage storages.IStorage) (string, error) {
	code, _ := encoder.DecodeInt64WithError(hash)
	if(len(code) == 0) {
		return "", errors.New("No hash to decode")
	}
	codeString := strconv.FormatUint(uint64(code[0]), 10)
	url, err := storage.Load(codeString)
	return url, err
}
