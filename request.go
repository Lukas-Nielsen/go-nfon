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
type MethodSuccess int

const (
	POST   method = "POST"
	PUT    method = "PUT"
	DELETE method = "DELETE"
	GET    method = "GET"

	POST_SUCCESS   MethodSuccess = 201
	PUT_SUCCESS    MethodSuccess = 204
	DELETE_SUCCESS MethodSuccess = 204
	GET_SUCCESS    MethodSuccess = 200
)

type ApiResponse struct {
	Href   string  `json:"href"`
	Offset int     `json:"offset"`
	Total  int     `json:"total"`
	Size   int     `json:"size"`
	Links  []Links `json:"links"`
	Data   []Data  `json:"data"`
	Items  []struct {
		Href  string  `json:"href"`
		Data  []Data  `json:"data"`
		Links []Links `json:"links"`
	} `json:"items"`
}

type ApiRequest struct {
	*Config
	Links []Links `json:"links,omitempty"`
	Data  []Data  `json:"data,omitempty"`
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

type Data struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

func DataToMap(data []Data) map[DataName]any {
	result := make(map[DataName]any)
	for _, entry := range data {
		result[DataName(entry.Name)] = entry.Value
	}
	return result
}

type Links struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func LinksToMap(data []Links) map[LinkRel]string {
	result := make(map[LinkRel]string)
	for _, entry := range data {
		result[LinkRel(entry.Rel)] = entry.Href
	}
	return result
}

func (a *ApiRequest) Send(method method, path string, result *ApiResponse) (MethodSuccess, apiError) {
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

	statusCode := MethodSuccess(0)
	var responseError apiError

	client := resty.New().
		SetBaseURL(API_URL).SetDebug(a.debug)

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
			statusCode = MethodSuccess(resp.StatusCode())
		} else if strings.Contains(err.Error(), "TLS handshake timeout") {
			return a.Send(method, path, result)
		} else {
			log.Println(err)
		}

	case GET:
		resp, err := request.
			SetError(&responseError).
			SetResult(&result).
			Get(path)
		if err == nil {
			statusCode = MethodSuccess(resp.StatusCode())
		} else if strings.Contains(err.Error(), "TLS handshake timeout") {
			return a.Send(method, path, result)
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
			statusCode = MethodSuccess(resp.StatusCode())
			if statusCode == 500 {
				return a.Send(method, path, result)
			}
		} else if strings.Contains(err.Error(), "TLS handshake timeout") {
			return a.Send(method, path, result)
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
			statusCode = MethodSuccess(resp.StatusCode())
			if statusCode == 500 {
				return a.Send(method, path, result)
			}
		} else if strings.Contains(err.Error(), "TLS handshake timeout") {
			return a.Send(method, path, result)
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
	a.Links = append(a.Links, Links{
		Rel:  rel,
		Href: href,
	})
	return a
}

func (a *ApiRequest) AddData(name string, value any) *ApiRequest {
	a.Data = append(a.Data, Data{
		Name:  name,
		Value: value,
	})
	return a
}
