package main

import (
	"learngo/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

type UserError interface {
	error
	Message() string
}

func errorWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				log.Printf(`Error occured handling %s`, r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK

			if userErr, ok := err.(UserError); ok {
				log.Printf(userErr.Message())
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errorWrapper(filelisting.HandleFileListing))
	http.ListenAndServe(":8080", nil)
}
