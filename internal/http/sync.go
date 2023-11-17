package httpServerSync

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func File(file *os.File) {
	defer file.Close()
	// Create a new buffer to store the request body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		panic(err)
	}

	// Copy the file content to the part
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}

	// Close the writer to finish the multipart form
	err = writer.Close()
	if err != nil {
		panic(err)
	}

	// Create a new HTTP request
	request, err := http.NewRequest("POST", "http://localhost:8080/upload", body)
	if err != nil {
		panic(err)
	}

	// Set the Content-Type header to multipart/form-data with a boundary
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Make the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Check the response status
	if response.StatusCode == 200 {
		fmt.Println("Backed up successfully:", file.Name())
	} else {
		fmt.Println("Server could not sync", file.Name())
		panic("")
	}
}
