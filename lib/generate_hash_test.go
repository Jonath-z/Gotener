package lib

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateHashWithSha256(t *testing.T) {
	longUrl := "https://medium.com/@jonathanzihindula95/automating-commit-messages-with-commit-ai-a-guide-to-building-a-github-tool-in-go-33878ec22754"
	hash := GenerateHashWithSha256(longUrl)

	fmt.Println(hash)
	assert.Equal(t, hash, "5a7f3a99")
}
