package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	// URL - url for mock payments api
	// See disclaimer in readme.md
	URL = ""
	// DebtsEndpoint - endpoint to retrieve debts
	DebtsEndpoint = "debts"
	// PaymentPlansEndpoint - endpoing to retrieve payment plans
	PaymentPlansEndpoint = "payment_plans"
	// PaymentsEndpoint - endpoint to retrieve payments
	PaymentsEndpoint = "payments"
)

// GetDebts - Returns all debts from debts endpoint
func GetDebts() (Debts, error) {
	fullURL := fmt.Sprintf("%s/%s", URL, DebtsEndpoint)
	rawBody, err := makeGetRequest(fullURL)
	if err != nil {
		return nil, err
	}
	var debts Debts
	err = json.Unmarshal(rawBody, &debts)
	if err != nil {
		return nil, err
	}
	return debts, nil
}

// GetPaymentPlans - Returns all payment plans from payment plans endpoint
func GetPaymentPlans() (PaymentPlans, error) {
	fullURL := fmt.Sprintf("%s/%s", URL, PaymentPlansEndpoint)
	rawBody, err := makeGetRequest(fullURL)
	if err != nil {
		return nil, err
	}
	var plans PaymentPlans
	err = json.Unmarshal(rawBody, &plans)
	if err != nil {
		return nil, err
	}
	return plans, nil
}

// GetPayments - Returns all payment plans from payments endpoint
func GetPayments() (Payments, error) {
	fullURL := fmt.Sprintf("%s/%s", URL, PaymentsEndpoint)
	rawBody, err := makeGetRequest(fullURL)
	if err != nil {
		return nil, err
	}
	var payments Payments
	err = json.Unmarshal(rawBody, &payments)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func makeGetRequest(requestURL string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, requestURL, bytes.NewBuffer(nil))
	if err != nil {
		return nil, fmt.Errorf("unable to create get request err: %w", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to make get request err: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body")
	}
	return body, err
}
