package gopyrest

// [go-py-rest]
import (
	"net/http"
	"io/ioutil"
)

func CallPythonService(url string, data string) (string, error) {
	resp, err := http.Post(url, "application/json", strings.NewReader(data))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
