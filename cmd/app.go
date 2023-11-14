package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	// Open the file that you want to send
	file, err := os.Open("/tmp/test.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new buffer to store the request body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a form file field
	part, err := writer.CreateFormFile("file", "file.txt")
	if err != nil {
		fmt.Println("Error writing to buffer:", err)
		return
	}

	// Copy the file content to the part
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file content:", err)
		return
	}

	// Close the writer to finish the multipart form
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing writer:", err)
		return
	}

	// Create a new HTTP request
	request, err := http.NewRequest("POST", "http://localhost:8080/upload", body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the Content-Type header to multipart/form-data with a boundary
	request.Header.Set("Content-Type", writer.FormDataContentType())

	// Make the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer response.Body.Close()

	// Check the response status
	fmt.Println("Response Status:", response.Status)
}
