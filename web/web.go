package web

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/aronkst/go-file-hosting/data"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func HandlerURL(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := getBody(r)
	if err != nil {
		httpError(w, err)
		return
	}

	fileName := body.Name + "." + body.Extension

	file, err := os.Create("static/" + fileName)
	if err != nil {
		httpError(w, err)
		return
	}
	defer file.Close()

	downloadFile, err := http.Get(body.Url)
	if err != nil {
		httpError(w, err)
		return
	}
	defer downloadFile.Body.Close()

	_, err = io.Copy(file, downloadFile.Body)
	if err != nil {
		httpError(w, err)
		return
	}

	output := data.Output{Url: "/" + fileName}
	outputBody, err := json.Marshal(output)
	if err != nil {
		httpError(w, err)
		return
	}

	w.Write(outputBody)
}

func HandlerFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseMultipartForm(10 << 20)

	uploadFile, info, err := r.FormFile("file")
	if err != nil {
		httpError(w, err)
		return
	}
	defer uploadFile.Close()

	fileName := info.Filename
	if fileName == "" {
		httpError(w, err)
		return
	}

	file, err := os.Create("static/" + fileName)
	if err != nil {
		httpError(w, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, uploadFile)
	if err != nil {
		httpError(w, err)
		return
	}

	output := data.Output{Url: "/" + fileName}
	outputBody, err := json.Marshal(output)
	if err != nil {
		httpError(w, err)
		return
	}

	w.Write(outputBody)
}

func getBody(r *http.Request) (data.Body, error) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return data.Body{}, err
	}

	body := data.Body{}
	err = json.Unmarshal(requestBody, &body)
	if err != nil {
		return data.Body{}, err
	}

	if body.Url == "" {
		return data.Body{}, errors.New("invalid url")
	}

	if body.Name == "" {
		body.Name = generateName()
	}

	if body.Extension == "" {
		body.Extension = getExtension(body.Url)
	}

	return body, nil
}

func generateName() string {
	uuid := uuid.New()
	value := uuid.String()

	value = strings.ReplaceAll(value, "-", "")
	return value
}

func getExtension(url string) string {
	values := strings.Split(url, ".")
	extension := values[len(values)-1]

	if strings.Contains(extension, "?") {
		values = strings.Split(extension, "?")
		extension = values[0]
	}

	return extension
}

func httpError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
