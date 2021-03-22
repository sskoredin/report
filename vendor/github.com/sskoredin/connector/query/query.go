package query

import (
	"crypto/tls"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Query string

func (q Query) Get() ([]byte, error) {
	resp, err := http.Get(q.string())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (q Query) Post(contentType string, body io.Reader) ([]byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Post(q.string(), contentType, body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if err := checkStatus(resp); err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (q Query) GetWithAttemptsInStruct(result interface{}) error {
	d, err := q.GetWithAttempts()
	if err != nil {
		return err
	}

	return json.Unmarshal(d, result)
}

func (q Query) getWithAttempts(attempt int) ([]byte, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	d, err := q.Get()
	if err != nil {
		if attempt > 10 {
			return nil, err
		}
		attempt++
		time.Sleep(1 * time.Second)
		return q.getWithAttempts(attempt)
	}
	return d, err
}

func (q Query) GetWithAttempts() ([]byte, error) {
	return q.getWithAttempts(0)
}

func (q Query) string() string {
	return string(q)
}

func checkStatus(resp *http.Response) error {
	if resp.StatusCode == 200 {
		return nil
	}
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var respErr iikoCardErrResponse
	if err := json.Unmarshal(d, &respErr); err != nil {
		return err
	}

	return errors.New(respErr.Message)
}

type iikoCardErrResponse struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}
