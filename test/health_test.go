// +build e2e

package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	fmt.Println("Running e2e tests for health check endpoint")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/health")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())

	//fmt.Println(resp.StatusCode())
}
