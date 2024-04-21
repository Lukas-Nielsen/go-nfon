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

func (e Error) Log() {
	log.Println(e.Title+":", e.Detail)
	if len(e.Errors) > 0 {
		log.Println("details")
		for _, entry := range e.Errors {
			entry := entry
			log.Println(entry.Message, "@", entry.Path, "with value", entry.Value)
		}
	}
}

func (r *Request) Send(method method, path string, result *Response) (int, Error) {
	return r.send(method, path, true, result)
}

func (r *Request) send(method method, path string, first bool, result *Response) (int, Error) {
	if first {
		r.client.requestCount += 1
	}
	var dataByte []byte
	if len(r.Data) == 0 && len(r.Links) == 0 {
		dataByte = []byte{}
	} else {
		temp, err := json.Marshal(r)
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
	h := hmac.New(sha1.New, []byte(r.client.config.secret))
	h.Write([]byte(string_to_sign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	contentLength := fmt.Sprint(len(string(dataByte)))

	statusCode := 0
	var responseError Error
	var response response

	client := resty.New().
		SetBaseURL(r.client.config.uri).SetDebug(r.client.config.debug)

	request := client.R().
		SetHeader("Authorization", "NFON-API "+r.client.config.key+":"+signature).
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
			statusCode = resp.StatusCode()
		} else if strings.Contains(err.Error(), "TLS handshake timeout") {
			return r.send(method, path, false, result)
		} else {
			log.Println(err)
		}

	case GET:
		resp, err := request.
			SetError(&responseError).
			SetResult(&response).
			Get(path)
		if err == nil {
			statusCode = resp.StatusCode()
		} else if strings.Contains(err.Error(), "TLS handshake timeout") {
			return r.send(method, path, false, result)
		} else {
			log.Println(err)
		}

	case POST:
		resp, err := request.
			SetError(&responseError).
			SetResult(&response).
			SetBody(string(dataByte)).
			Post(path)
		if err == nil {
			statusCode = resp.StatusCode()
			if statusCode == 500 {
				return r.send(method, path, false, result)
			}
		} else if strings.Contains(err.Error(), "TLS handshake timeout") {
			return r.send(method, path, false, result)
		} else {
			log.Println(err)
		}

	case PUT:
		resp, err := request.
			SetError(&responseError).
			SetBody(string(dataByte)).
			Put(path)
		if err == nil {
			statusCode = resp.StatusCode()
			if statusCode == 500 {
				return r.send(method, path, false, result)
			}
		} else if strings.Contains(err.Error(), "TLS handshake timeout") {
			return r.send(method, path, false, result)
		} else {
			log.Println(err)
		}
	}

	response.parse(result)
	return statusCode, responseError
}

func (c *Client) NewRequest() *Request {
	return &Request{
		client: c,
	}
}

func (r *Request) AddLink(rel string, href string) *Request {
	r.Links = append(r.Links, links{
		Rel:  rel,
		Href: href,
	})
	return r
}

func (r *Request) AddData(name string, value any) *Request {
	r.Data = append(r.Data, data{
		Name:  name,
		Value: value,
	})
	return r
}
