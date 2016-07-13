package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gimanzo/shawty/handlers"
	"github.com/gimanzo/shawty/storages"
	"github.com/speps/go-hashids"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	storage := &storages.Filesystem{}
	storageEnv := "SHAWTY_STORAGE_PATH"
	saltEnv := "SHAWTY_HASH_SALT"

	storagePath := os.Getenv(storageEnv)
	if storagePath == ""  {
		fmt.Println(storageEnv, "not defined, exiting.")
		os.Exit(1)
	}
	fmt.Println(storageEnv, ":", storagePath)

	hashData := hashids.NewData()
	salt := os.Getenv(saltEnv)
	if salt == "" {
		fmt.Println(saltEnv, "not defined, exiting.")
		os.Exit(1)
	}
	hashData.Salt = salt
	hashData.MinLength = 7

	encoder := hashids.NewWithData(hashData)


	err := storage.Init(storagePath)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/encode", handlers.EncodeHandler(storage, encoder))
	http.Handle("/decode/", handlers.DecodeHandler(storage, encoder))
	http.Handle("/redirect/", handlers.RedirectHandler(storage, encoder))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
