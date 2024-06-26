package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/diogovalentte/linkwarden-iframe/src/config"
)

func baseRequest(url string, target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+config.LinkwardenToken)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("Error unmarshaling JSON: %s\nReponse text: %s", err.Error(), string(body))
	}

	return nil
}
