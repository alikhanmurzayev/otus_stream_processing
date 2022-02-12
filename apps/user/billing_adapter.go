package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type billingAdapter struct {
	billingServiceURL string
}

func NewBillingAdapter(billingServiceURL string) *billingAdapter {
	return &billingAdapter{billingServiceURL: billingServiceURL}
}

func (adapter *billingAdapter) CreateAccount(ctx context.Context, userID int64) error {
	reqBody := map[string]interface{}{"user_id": userID}
	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(reqBody); err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, adapter.billingServiceURL+"/create", body)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return fmt.Errorf("NewRequestWithContext: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not send request: %w", err)
	}
	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("got status code %d. Response from billing service: %s", resp.StatusCode, string(respBody))
	}
	return nil
}
