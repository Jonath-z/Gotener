package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/url_shortner/db"
	"github.com/url_shortner/lib"
	models2 "github.com/url_shortner/models"
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
	database := db.Database{}

	if database.Db == nil {
		database.Connect()
	}

	_, saveErr := database.Db.Model(&models2.ShortUrls{
		LongUrl:  parseBody.LongUrl,
		ShortUrl: "https://short/" + urlHash,
		Hash:     urlHash,
	}).Insert()

	if saveErr != nil {
		fmt.Println(saveErr)
		json.NewEncoder(w).Encode(models2.ErrorResponse{
			Message: "Unable to save the shortened Url",
		})
		return
	}

	json.NewEncoder(w).Encode(models2.ShortenedUrl{
		Hash:     urlHash,
		ShortUrl: "https://short/" + urlHash,
		LongUrl:  parseBody.LongUrl,
	})
}
