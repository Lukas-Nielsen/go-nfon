package nfon

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
)

type method string
type methodSuccess int

const (
	POST   method = "POST"
	PUT    method = "PUT"
	DELETE method = "DELETE"
	GET    method = "GET"

	POST_SUCCESS   methodSuccess = 201
	PUT_SUCCESS    methodSuccess = 204
	DELETE_SUCCESS methodSuccess = 204
	GET_SUCCESS    methodSuccess = 200
)

type apiResponse struct {
	Href   string  `json:"href"`
	Offset int     `json:"offset"`
	Total  int     `json:"total"`
	Size   int     `json:"size"`
	Links  []links `json:"links"`
	Data   []data  `json:"data"`
	Items  []struct {
		Href  string  `json:"href"`
		Data  []data  `json:"data"`
		Links []links `json:"links"`
	} `json:"items"`
}

type ApiRequest struct {
	*Config
	Links []links `json:"links,omitempty"`
	Data  []data  `json:"data,omitempty"`
}

type apiError struct {
	Detail      string           `json:"detail"`
	Title       string           `json:"title"`
	DescribedBy string           `json:"described_by"`
	Errors      []apiErrorErrors `json:"errors"`
}

type apiErrorErrors struct {
	Message string `json:"message"`
	Path    string `json:"path"`
	Value   string `json:"value"`
}

func (e apiError) log() {
	log.Println(e.Title+":", e.Detail)
	if len(e.Errors) > 0 {
		log.Println("details")
		for _, entry := range e.Errors {
			entry := entry
			log.Println(entry.Message, "@", entry.Path, "with value", entry.Value)
		}
	}
}

type data struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type links struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func (a *ApiRequest) send(method method, path string, result *apiResponse) (methodSuccess, apiError) {
	var dataByte []byte
	if len(a.Data) == 0 && len(a.Links) == 0 {
		dataByte = []byte{}
	} else {
		temp, err := json.Marshal(a)
		if err != nil {
			log.Fatalln(err)
		}
		dataByte = temp
	}
	data_md5 := md5.Sum(dataByte)
	data_md5_hex := hex.EncodeToString(data_md5[:])
	content_type := "application/json; charset=utf-8"
	current_time := time.Now().UTC()
	request_date := strings.Replace(current_time.Format(time.RFC1123), "UTC", "GMT", -1)
	string_to_sign := string(method) + "\n" + data_md5_hex + "\n" + content_type + "\n" + request_date + "\n" + path
	h := hmac.New(sha1.New, []byte(a.secret))
	h.Write([]byte(string_to_sign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	contentLength := fmt.Sprint(len(string(dataByte)))

	statusCode := methodSuccess(0)
	var responseError apiError

	client := resty.New().
		SetBaseURL(API_URL).
		SetDebug(true)

	request := client.R().
		SetHeader("Authorization", "NFON-API "+a.key+":"+signature).
		SetHeader("Content-MD5", data_md5_hex).
		SetHeader("Content-Length", contentLength).
		SetHeader("Content-Type", content_type).
		SetHeader("x-nfon-date", request_date).
		SetHeader("User-Agent", "github.com/Lukas-Nielsen/go-nfon").
		SetHeader("Accept", "*/*")

	switch method {

	case DELETE:
		resp, err := request.
			SetError(&responseError).
			Delete(path)
		if err == nil {
			statusCode = methodSuccess(resp.StatusCode())
		} else if strings.Contains(err.Error(), "TLS handshake timeout") {
			return a.send(method, path, result)
		} else {
			log.Println(err)
		}

	case GET:
		resp, err := request.
			SetError(&responseError).
			SetResult(&result).
			Get(path)
		if err == nil {
			statusCode = methodSuccess(resp.StatusCode())
		} else if strings.Contains(err.Error(), "TLS handshake timeout") {
			return a.send(method, path, result)
		} else {
			log.Println(err)
		}

	case POST:
		resp, err := request.
			SetError(&responseError).
			SetResult(&result).
			SetBody(string(dataByte)).
			Post(path)
		if err == nil {
			statusCode = methodSuccess(resp.StatusCode())
			if statusCode == 500 {
				return a.send(method, path, result)
			}
		} else if strings.Contains(err.Error(), "TLS handshake timeout") {
			return a.send(method, path, result)
		} else {
			log.Println(err)
		}

	case PUT:
		resp, err := request.
			SetError(&responseError).
			SetResult(&result).
			SetBody(string(dataByte)).
			Put(path)
		if err == nil {
			statusCode = methodSuccess(resp.StatusCode())
			if statusCode == 500 {
				return a.send(method, path, result)
			}
		} else if strings.Contains(err.Error(), "TLS handshake timeout") {
			return a.send(method, path, result)
		} else {
			log.Println(err)
		}
	}
	return statusCode, responseError
}

func (c *Config) NewRequest() *ApiRequest {
	return &ApiRequest{
		Config: c,
	}
}

func (a *ApiRequest) AddLink(rel string, href string) *ApiRequest {
	a.Links = append(a.Links, links{
		Rel:  rel,
		Href: href,
	})
	return a
}

func (a *ApiRequest) AddData(name string, value any) *ApiRequest {
	a.Data = append(a.Data, data{
		Name:  name,
		Value: value,
	})
	return a
}
