package cmd

import (
	"fmt"
	"net/http"

	"github.com/url_shortner/handlers"
	"github.com/url_shortner/middleware"
)

type Server struct{}

func (s *Server) Run() {
	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		fmt.Println("Server failed to run due to: ", err)
	} else {
		fmt.Println("server running on Port 4444")
	}
}

func (s *Server) InitializeRoutes() {
	requestShortUrl := http.HandlerFunc(handlers.RequestShortUrl)
	http.Handle("/shortener", middleware.ValidateShortUrlRequest(requestShortUrl))
	http.HandleFunc("/", handlers.RedirectToLongUrl)
}
