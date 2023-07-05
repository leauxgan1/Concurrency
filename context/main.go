package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel()

	// Create http request
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "http://placehold.it/200x200", nil)
	if err != nil {
		panic(err)
	}

	// Send http request and receive result
	result, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	// Close the connection to the body (?)
	defer result.Body.Close()

	// Read the image data from the body
	imageData, err := ioutil.ReadAll(result.Body)

	fmt.Printf("The image was downloaded, its size was %d\n", len(imageData))

}
