package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
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

func TestLoadPing(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://127.0.0.1:" + getEnvValue("PORT") + "/ping",
	})

	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics

	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}

	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestDeposit(t *testing.T) {
	body := []byte(`{
		"id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15",
		"funds" : 33.3
	}`)

	req, err := http.NewRequest(http.MethodPatch, "http://127.0.0.1:"+getEnvValue("PORT")+"/deposit", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
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
}

func TestLoadDeposit(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "PATCH",
		URL:    "http://127.0.0.1:" + getEnvValue("PORT") + "/deposit",
		Body: []byte(`{
			"id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15",
			"funds" : 33.3
		}`),
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestAddUser(t *testing.T) {
	body := []byte(`{
		"funds" : 33.3
	}`)

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:"+getEnvValue("PORT")+"/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
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
}

func TestLoadAddUser(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    "http://127.0.0.1:" + getEnvValue("PORT") + "/user",
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestDelete(t *testing.T) {
	body := []byte(`{
		"id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15"
	}`)

	req, err := http.NewRequest(http.MethodDelete, "http://127.0.0.1:"+getEnvValue("PORT")+"/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
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
}

func TestLoadDelete(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "DELETE",
		URL:    "http://127.0.0.1:" + getEnvValue("PORT") + "/user",
		Body: []byte(`{
			"id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15"
		}`),
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestReserve(t *testing.T) {
	body := []byte(`{
		"userID" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15",
		"serviceID" : "ec6741fa-4b02-4e03-a303-0fa96eb15d15",
		"orderID" : "ec6705fa-4b00-0e11-a013-8fa88eb74d35",
		"cost" : 3.5
	}`)

	req, err := http.NewRequest(http.MethodPatch, "http://127.0.0.1:"+getEnvValue("PORT")+"/reserve", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
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
}

func TestLoadReserve(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "PATCH",
		URL:    "http://127.0.0.1:" + getEnvValue("PORT") + "/reserve",
		Body: []byte(`{
			"userID" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15",
			"serviceID" : "ec6741fa-4b02-4e03-a303-0fa96eb15d15",
			"orderID" : "ec6705fa-4b00-0e11-a013-8fa88eb74d35",
			"cost" : 3.5
		}`),
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestReceipt(t *testing.T) {
	body := []byte(`{
		"userID" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15",
		"serviceID" : "ec6741fa-4b02-4e03-a303-0fa96eb15d15"
	}`)

	req, err := http.NewRequest(http.MethodPatch, "http://127.0.0.1:"+getEnvValue("PORT")+"/receipt", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
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
}

func TestLoadReceipt(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "PATCH",
		URL:    "http://127.0.0.1:" + getEnvValue("PORT") + "/receipt",
		Body: []byte(`{
			"userID" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15",
			"serviceID" : "ec6741fa-4b02-4e03-a303-0fa96eb15d15"
		}`),
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}

func TestBalance(t *testing.T) {
	body := []byte(`{
		"id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15"
	}`)

	req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:"+getEnvValue("PORT")+"/balance", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
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
}

func TestLoadBalance(t *testing.T) {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://127.0.0.1:" + getEnvValue("PORT") + "/balance",
		Body: []byte(`{
			"id" : "ec6761fa-4b02-4e93-a213-8fa96eb44d15"
		}`),
	})
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
		metrics.Add(res)
	}
	metrics.Close()
	log.Printf("99th percentile: %s\n", metrics.Latencies.P99)
}
