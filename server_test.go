package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessShouldSucceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleSuccess))
	defer testServer.Close()

	testClient := testServer.Client()

	// fmt.Println(testServer.URL)
	// testClient.Get(testServer.URL)

	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "success!", string(body))
	assert.Equal(t, 200, response.StatusCode)
}

func TestSuccessShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleSuccess))
	defer testServer.Close()

	testClient := testServer.Client()

	body := strings.NewReader("does not metter what the body is!")

	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}

func TestHealthShouldSucceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()

	// fmt.Println(testServer.URL)
	// testClient.Get(testServer.URL)

	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "ok!", string(body))
	assert.Equal(t, 200, response.StatusCode)
}

func TestHealthShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()

	body := strings.NewReader("does not metter what the body is!")

	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}

func TestBrandNewEndpointShouldSucceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleBrandNewEndpoint))
	defer testServer.Close()

	testClient := testServer.Client()

	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "brand-new-endpoint!", string(body))
	assert.Equal(t, 200, response.StatusCode)
}

func TestBrandNewEndpointShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleBrandNewEndpoint))
	defer testServer.Close()

	testClient := testServer.Client()

	body := strings.NewReader("does not metter what the body is!")

	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)
}
