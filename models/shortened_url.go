package models

type ShortenedUrl struct {
	LongUrl  string `json:"longUrl"`
	ShortUrl string `json:"shortUrl"`
	Hash     string `json:"hash"`
	Encoding string `json:"encoding"`
}
