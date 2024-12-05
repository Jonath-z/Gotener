package middleware

import (
	"bytes"
	"encoding/json"
	models2 "github.com/url_shortner/models"
	"io"
	"net/http"
)

func ValidateShortUrlRequest(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(models2.ErrorResponse{Message: "Only post request is allowed for this route"})
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(models2.ErrorResponse{Message: "Cannot process the body"})
			return
		}

		var reqBody models2.RequestShortUrl
		parsingErr := json.Unmarshal(body, &reqBody)
		if parsingErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models2.ErrorResponse{Message: "Invalid body"})
			return
		}

		if reqBody.LongUrl == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models2.ErrorResponse{Message: "The url cannot be empty"})
			return
		}
		// Re-encode the body to be used in next handler functions
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		h.ServeHTTP(w, r)
	})
}
