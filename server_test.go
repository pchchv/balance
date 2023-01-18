package main

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestServerPing(t *testing.T) {
	res, err := http.Get("http://127.0.0.1:" + getEnvValue("PORT") + "/ping")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("status not OK")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Error(err)
		}
	}(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	b := string(body)
	if !strings.Contains(b, "Balance") {
		t.Fatal()
	}
}
