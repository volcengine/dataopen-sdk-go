package client

import (
	"fmt"
	"testing"
)

// Test Access Token
func TestGetToken(t *testing.T) {
	app_id := ""
	app_secret := ""

	client := Client(app_id, app_secret, "", "", "")

	if err := client.GetToken(); err != nil {
		t.Fatal(err)
	}

	fmt.Println("is_authenticated", client.IsAuthenticated())

	if !client.IsAuthenticated() {
		t.Errorf("Expected access token to not be null")
	}
}

// Test Request GET
func TestGetRequest(t *testing.T) {
	app_id := ""
	app_secret := ""

	client := Client(app_id, app_secret, "", "", "")

	headers := make(map[string]string)

	params := make(map[string]ParamsValueType)
	params["app"] = "46"
	params["page_size"] = "1"
	params["page"] = "1"

	body := make(map[string]interface{})

	res, err := client.Request("/xxx/openapi/v1/open/flight-list", "GET", headers, params, body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Output: ", res)
}

// Test Request POST
func TestPostRequest(t *testing.T) {
	app_id := ""
	app_secret := ""

	client := Client(app_id, app_secret, "", "", "")

	headers := make(map[string]string)

	params := make(map[string]ParamsValueType)

	body := make(map[string]interface{})
	body["uid_list"] = []string{"1111111110000"}

	res, err := client.Request(
		"/xxx/openapi/v1/open/flight/version/6290880/add-test-user",
		"POST",
		headers,
		params,
		body,
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Output: ", res)
}

// Test material Request POST
func TestPostMaterialRequest(t *testing.T) {
	app_id := ""
	app_secret := ""

	client := Client(app_id, app_secret, "https://analytics.volcengineapi.com",
		"dataopen_staging", "")

	headers := make(map[string]string)

	params := make(map[string]ParamsValueType)

	body := make(map[string]interface{})

	body["name"] = "ccnnodetest"
	body["title"] = "测试title"
	body["type"] = "component"
	body["description"] = "测试description4"
	body["frameworkType"] = "react"

	res, err := client.Request(
		"/material/openapi/v1/material",
		"PUT",
		headers,
		params,
		body,
	)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Output: ", res)
}
