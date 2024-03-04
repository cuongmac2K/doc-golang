package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func AddNewTaskQueue(
	Env string,
	request *http.Request,
	path string,
	param url.Values,
	queueName string,
) (task interface{}, err error) {
	u, _ := url.ParseRequestURI("http://localhost:8080/")
	u.Path = fmt.Sprintf("/add-taskqueue/%s", queueName)
	urlStr := u.String()

	var client = &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(param.Encode())) // URL-encoded payload
	r.Header.Add("Authorization", fmt.Sprintf("auth_token=\"%s\"", Env))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("taskqueue", path)
	r.Header.Add("request", request.Host)

	_, err = client.Do(r)
	return
}

func SetupNewQueue(
	Env string,
	request *http.Request,
	jsonData []byte,
) (task interface{}, err error) {
	u, _ := url.ParseRequestURI("http://localhost:8080/")
	u.Path = fmt.Sprintf("/setup-queue")
	urlStr := u.String()
	var client = &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(jsonData)) // URL-encoded payload
	r.Header.Add("Authorization", fmt.Sprintf("auth_token=\"%s\"", Env))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("request", request.Host)
	_, err = client.Do(r)
	return
}
func DeleteQueue(
	Env string,
	request *http.Request,
	queueName string,
) (task interface{}, err error) {
	u, _ := url.ParseRequestURI("http://localhost:8080/")
	u.Path = fmt.Sprintf("/delete-queue/%s", queueName)
	urlStr := u.String()
	var client = &http.Client{}
	r, _ := http.NewRequest(http.MethodDelete, urlStr, nil) // URL-encoded payload
	r.Header.Add("Authorization", fmt.Sprintf("auth_token=\"%s\"", Env))
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("request", request.Host)
	_, err = client.Do(r)
	return
}
