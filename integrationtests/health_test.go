package integrationtests

import (
	router "github.com/charan678/go-rest-example/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func shouldTestHealthAPI(t *testing.T) {
	ts := httptest.NewServer(router.SetupRoutes())
	defer ts.Close()
	resp, err := http.Get("/v1/health")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
