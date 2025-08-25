package shortlinkapi

// [go-shortlink-api]
import (
	"net/http"
)

func ShortenURL(url string) (string, error) {
	resp, err := http.Get("{SHORTENER_API}" + "?url=" + url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// Parse response to get shortened URL
	return "shortened-url", nil
}
