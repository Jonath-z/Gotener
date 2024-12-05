package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/url_shortner/lib"
	models2 "github.com/url_shortner/models"
	"io"
	"net/http"
	"strconv"
	"time"
)

func RequestShortUrl(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var parseBody models2.RequestShortUrl
	err := json.Unmarshal(body, &parseBody)
	if err != nil {
		fmt.Println(err)
	}

	timeStamp := time.Now().Nanosecond()
	urlHash := lib.GenerateHashWithSha256(parseBody.LongUrl + strconv.Itoa(timeStamp))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models2.ShortenedUrl{
		Hash:     "ljdieks",
		ShortUrl: "https://short/ljdieks",
		Encoding: "base64",
		LongUrl:  urlHash,
	})
}
