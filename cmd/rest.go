package cmd

import (
	"gopkg.in/resty.v1"
	"fmt"
	"log"
	"encoding/json"
	"bytes"
)

func prepareRequest() *resty.Request {
	var r = resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("User-Agent", "pulsar-ctl 2.1.0")
	return r
}

func RestPut(path string, content interface{}) {
	resp, err := prepareRequest().
		SetBody(content).
		Put(adminUrl + path)

	if err != nil {
		log.Fatal("REST call failed: ", err)
	}

	if resp.StatusCode() != 204 {
		logErrorReasonAndExit(resp)
	}
}

func RestPost(path string, content interface{}) {
	resp, err := prepareRequest().
		SetBody(content).
		Post(adminUrl + path)

	if err != nil {
		log.Fatal("REST call failed: ", err)
	}

	if resp.StatusCode() != 204 {
		logErrorReasonAndExit(resp)
	}
}

func RestDelete(path string) {
	resp, err := prepareRequest().
		Delete(adminUrl + path)

	if err != nil {
		log.Fatal("REST call failed: ", err)
	}

	if resp.StatusCode() != 204 {
		logErrorReasonAndExit(resp)
	}
}

func RestGet(path string) string {
	resp, err := prepareRequest().Get(adminUrl + path)

	if err != nil {
		log.Fatal("REST call failed: ", err)
	}

	if resp.StatusCode() != 200 {
		logErrorReasonAndExit(resp)
	}

	var out = bytes.Buffer{}
	json.Indent(&out, resp.Body(), "", "   ")
	return out.String()
}

func RestPrint(path string) {
	fmt.Println(RestGet( path))
}

func RestGetStringList(path string) []string {
	response := RestGet(path)

	var list []string
	json.Unmarshal([]byte(response), &list)
	return list
}

func RestPrintStringList(path string) {
	list := RestGetStringList(path)

	for _, item := range list {
		fmt.Println(item)
	}
}

type ErrorReason struct {
	Reason string `json: reason`
}

func logErrorReasonAndExit(response *resty.Response) {
	if response.Body() != nil {
		// Try to parse response as JSON
		var errorReason ErrorReason
		json.Unmarshal(response.Body(), &errorReason)
		log.Fatal("Request failed: ", errorReason.Reason)
	} else {
		log.Fatal("Request failed: ", response.Status())
	}
}

func init() {
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(20))
	resty.SetHostURL(adminUrl)
}
