package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/url_shortner/db"
	models2 "github.com/url_shortner/models"
)

func RedirectToLongUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}
	path := r.URL.Path
	pathComponents := strings.Split(path, "/")
	hash := pathComponents[len(pathComponents)-1]

	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models2.ErrorResponse{
			Message: "Invalid Url",
		})
		return
	}

	database := db.Database{}
	if database.Db == nil {
		database.Connect()
	}

	longUrl := new(models2.ShortUrls)
	err := database.Db.Model(longUrl).Where("hash = ?", hash).Select()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models2.ErrorResponse{
			Message: "Url not found",
		})
		return
	}

	http.Redirect(w, r, longUrl.LongUrl, http.StatusMovedPermanently)
}
