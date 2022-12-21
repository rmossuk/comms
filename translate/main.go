package main

import (
	"fmt"
	"net/http"

	"io/ioutil"
	"strings"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {

		response := "No text passed to translate"
		path := r.URL.Path

		if strings.Contains(path, "translate/") {

			pathParts := strings.Split(path, "translate/")

			if len(pathParts) > 0 {

				text := pathParts[1]

				fmt.Println("calling API traslation", "https://api.funtranslations.com/translate/klingon.json?text="+text)

				r1, _ := spinhttp.Get("https://api.funtranslations.com/translate/klingon.json?text=" + text)

				body, err := ioutil.ReadAll(r1.Body)
				if err != nil {
					return
				}

				bodyString := string(body)

				if !(strings.Contains(bodyString, "error")) {
					bodyPart1 := strings.Split(bodyString, "\"translated\": \"")
					bodyParts2 := strings.Split(bodyPart1[1], "\",")
					response = bodyParts2[0]
				} else {
					// response = "Too Many Requests: Rate limit of 5 requests per hour exceeded"
					response = "Je Many Requests: Rate vus of 5 requests per rep exceeded"
				}
			}
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, response)
	})
}

func main() {}
