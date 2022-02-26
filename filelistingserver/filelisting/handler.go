package filelisting

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type userError string

func (err userError) Error() string {
	return err.Message()
}

func (err userError) Message() string {
	return string(err)
}

func HandleFileListing(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path
	prefix := strings.HasPrefix(path, "/list/")
	if !prefix {
		log.Printf(`path must be start with "/list/"`)
		return userError(`path must be start with "/list/"`)
	}
	filename := path[len("/list/"):]
	open, err := os.Open(filename)
	defer open.Close()
	if err != nil {
		return err
	}
	all, err := ioutil.ReadAll(open)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
