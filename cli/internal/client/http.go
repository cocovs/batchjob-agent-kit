package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	BaseURL string
	Token   string
	Timeout time.Duration
}

type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

func New(cfg Config) (*Client, error) {
	baseURL, err := normalizeBaseURL(cfg.BaseURL)
	if err != nil {
		return nil, err
	}
	timeout := cfg.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	return &Client{
		baseURL: baseURL,
		token:   strings.TrimSpace(cfg.Token),
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func normalizeBaseURL(raw string) (string, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		raw = "http://127.0.0.1:8080"
	}
	if !strings.Contains(raw, "://") {
		if strings.Contains(raw, ":") {
			raw = "http://" + raw
		} else {
			raw = "https://" + raw
		}
	}
	u, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("parse server URL: %w", err)
	}
	if u.Scheme == "" || u.Host == "" {
		return "", fmt.Errorf("invalid server URL: %s", raw)
	}
	path := strings.TrimRight(u.Path, "/")
	switch {
	case path == "":
		u.Path = ""
	case path == "/batch":
		u.Path = "/batch"
	default:
		u.Path = path
	}
	return strings.TrimRight(u.String(), "/"), nil
}

func (c *Client) endpoint(path string) string {
	return c.baseURL + path
}

func (c *Client) GetJSON(ctx context.Context, path string, out any) error {
	return c.GetJSONWithQuery(ctx, path, nil, out)
}

func (c *Client) GetJSONWithQuery(ctx context.Context, path string, query url.Values, out any) error {
	endpoint := c.endpoint(path)
	if len(query) > 0 {
		endpoint += "?" + query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	return decodeResponse(resp, out)
}

func (c *Client) PostJSON(ctx context.Context, path string, in any, out any) error {
	var body io.Reader
	if in != nil {
		payload, err := json.Marshal(in)
		if err != nil {
			return fmt.Errorf("encode request JSON: %w", err)
		}
		body = bytes.NewReader(payload)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint(path), body)
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	return decodeResponse(resp, out)
}

func decodeResponse(resp *http.Response, out any) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response body: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("request failed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}
	if out == nil || len(body) == 0 {
		return nil
	}
	dec := json.NewDecoder(bytes.NewReader(body))
	if err := dec.Decode(out); err != nil {
		return fmt.Errorf("decode response JSON: %w", err)
	}
	return nil
}
