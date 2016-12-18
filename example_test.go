package travishook

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/travis-webhook", travisHandler)
	http.ListenAndServe(":8080", nil)
}

/* example implementation for a Travis CI Webhook HTTP Handler in Go */
func travisHandler(writer http.ResponseWriter, reader *http.Request) {

	/* check request type (HTTP Method) */
	if reader.Method != "POST" {
		/* 405 Method Not Allowed */
		writer.Header().Set("Allow", "POST")
		http.Error(writer, http.StatusText(405), 405)

		return
	}

	/* check Content-Type (Header) */
	content_type := reader.Header.Get("Content-Type")
	if content_type != "application/x-www-form-urlencoded" {
		/* 415 Unsupported Media Type */
		http.Error(writer, http.StatusText(415), 415)

		return
	}

	/* get payload */
	rawPayload := reader.FormValue("payload")
	if rawPayload == "" {
		/* 400 Bad Request */
		http.Error(writer, http.StatusText(400), 400)

		return
	}

	/* obtain and verify signature */
	signature := reader.Header.Get("Signature")
	if signature == "" {
		http.Error(writer, http.StatusText(400), 400)

		return
	}

	err := CheckSignature(signature, []byte(rawPayload), defaultServer)
	if err != nil {
		http.Error(writer, http.StatusText(400), 400)

		return
	}

	/* json-decode payload */
	payload := Payload{}
	err = json.Unmarshal([]byte(rawPayload), &payload)
	if err != nil {
		http.Error(writer, http.StatusText(400), 400)

		return
	}

	/* check build status */
	if payload.Status == 0 && payload.StatusMessage == "Passed" {
		/* build successful */
		/* do something */
	} else {
		/* build not successful */
		/* ignore? */
	}

	writer.WriteHeader(200)
	writer.Write([]byte("OK"))

	return
}
