package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	_ "github.com/luraproject/lura/v2/transport/http/server/plugin"
)

var pluginName = "token-validator"

var HandlerRegisterer = registerer(pluginName)

type registerer string

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

type TokenValidationResponse struct {
	UserID string `json:"userID"`
}

func (r registerer) registerHandlers(_ context.Context, extra map[string]interface{}, h http.Handler) (http.Handler, error) {
	config, ok := extra[pluginName].(map[string]interface{})
	if !ok {
		return h, errors.New("configuration not found")
	}

	validationURL, ok := config["validation_url"].(string)
	if !ok {
		return h, errors.New("validation_url is required in configuration")
	}

	logger.Debug(fmt.Sprintf("Token validator plugin initialized with validation URL: %s", validationURL))

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("X-Access-Token")
		if token == "" {
			logger.Warning("Missing X-Access-Token header")
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		userID, err := validateToken(validationURL, token)
		if err != nil {
			logger.Error(fmt.Sprintf("Token validation failed: %v", err))
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), "userID", userID)
		newReq := req.WithContext(ctx)
		newReq.Header.Add("X-User-ID", userID)

		h.ServeHTTP(w, newReq)
	}), nil
}

func validateToken(validationURL, token string) (string, error) {
	req, err := http.NewRequest("POST", validationURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create validation request: %v", err)
	}

	req.Header.Add("X-Access-Token", token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("validation request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("validation service returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read validation response: %v", err)
	}

	var validationResp TokenValidationResponse
	if err := json.Unmarshal(body, &validationResp); err != nil {
		return "", fmt.Errorf("failed to parse validation response: %v", err)
	}

	if validationResp.UserID == "" {
		return "", errors.New("empty userID in validation response")
	}

	return validationResp.UserID, nil
}

func main() {
	fmt.Println("Token validator plugin loaded successfully!")
}

var logger Logger = noopLogger{}

func (registerer) RegisterLogger(v interface{}) {
	l, ok := v.(Logger)
	if !ok {
		return
	}
	logger = l
	logger.Debug(fmt.Sprintf("[PLUGIN: %s] Logger loaded", HandlerRegisterer))
}

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})
	Critical(v ...interface{})
	Fatal(v ...interface{})
}

type noopLogger struct{}

func (n noopLogger) Debug(_ ...interface{})    {}
func (n noopLogger) Info(_ ...interface{})     {}
func (n noopLogger) Warning(_ ...interface{})  {}
func (n noopLogger) Error(_ ...interface{})    {}
func (n noopLogger) Critical(_ ...interface{}) {}
func (n noopLogger) Fatal(_ ...interface{})    {}
