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
	storage := getStorage()
	encoder := getEncoder()
	http.Handle("/encode", handlers.EncodeHandler(storage, encoder))
	http.Handle("/decode/", handlers.DecodeHandler(storage, encoder))
	http.Handle("/", handlers.RedirectHandler(storage, encoder))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func getStorage() storages.IStorage {
	storage := &storages.Filesystem{}
	storageEnv := "SHAWTY_STORAGE_PATH"
	storagePath := getEnv(storageEnv)
	fmt.Println(storageEnv, ":", storagePath)

	err := storage.Init(storagePath)
	if err != nil {
		log.Fatal(err)
	}
	return storage
}

func getEncoder() *hashids.HashID {
	hashData := hashids.NewData()

	saltEnv := "SHAWTY_HASH_SALT"
	salt := getEnv(saltEnv)
	hashData.Salt = salt
	hashData.MinLength = 7
	return hashids.NewWithData(hashData)
}

func getEnv(variableName string) string {
	variable := os.Getenv(variableName)
	if variableName == ""  {
		fmt.Println(variableName, "not defined, exiting.")
		os.Exit(1)
	}
	return variable
}
