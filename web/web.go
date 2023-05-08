package web

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/aronkst/go-file-hosting/data"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func HandlerURL(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	body, err := getBody(request)
	if err != nil {
		httpError(writer, err)

		return
	}

	fileName := body.Name + "." + body.Extension

	file, err := os.Create("static/" + fileName)
	if err != nil {
		httpError(writer, err)

		return
	}
	defer file.Close()

	downloadFile, err := http.Get(body.URL)
	if err != nil {
		httpError(writer, err)

		return
	}
	defer downloadFile.Body.Close()

	_, err = io.Copy(file, downloadFile.Body)
	if err != nil {
		httpError(writer, err)

		return
	}

	output := data.Output{URL: "/" + fileName}
	outputBody, err := json.Marshal(output)

	if err != nil {
		httpError(writer, err)

		return
	}

	_, err = writer.Write(outputBody)
	if err != nil {
		httpError(writer, err)

		return
	}
}

func HandlerFile(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	err := request.ParseMultipartForm(10 << 20)
	if err != nil {
		httpError(writer, err)

		return
	}

	uploadFile, info, err := request.FormFile("file")
	if err != nil {
		httpError(writer, err)

		return
	}
	defer uploadFile.Close()

	fileName := info.Filename
	if fileName == "" {
		httpError(writer, err)

		return
	}

	file, err := os.Create("static/" + fileName)
	if err != nil {
		httpError(writer, err)

		return
	}
	defer file.Close()

	_, err = io.Copy(file, uploadFile)
	if err != nil {
		httpError(writer, err)

		return
	}

	output := data.Output{URL: "/" + fileName}
	outputBody, err := json.Marshal(output)

	if err != nil {
		httpError(writer, err)

		return
	}

	_, err = writer.Write(outputBody)
	if err != nil {
		httpError(writer, err)

		return
	}
}

func getBody(request *http.Request) (data.Body, error) {
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		return data.Body{}, err
	}

	body := data.Body{}
	err = json.Unmarshal(requestBody, &body)

	if err != nil {
		return data.Body{}, err
	}

	if body.URL == "" {
		return data.Body{}, errors.New("invalid url")
	}

	if body.Name == "" {
		body.Name = generateName()
	}

	if body.Extension == "" {
		body.Extension = getExtension(body.URL)
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
