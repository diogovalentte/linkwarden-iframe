package routes_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	api "github.com/diogovalentte/linkwarden-iframe/src"
	"github.com/diogovalentte/linkwarden-iframe/src/config"
	"github.com/diogovalentte/linkwarden-iframe/src/models"
)

func setup() error {
	err := config.SetConfigs("../../.env.test")
	if err != nil {
		return err
	}

	return nil
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestRoutes(t *testing.T) {
	router := api.SetupRouter()

	t.Run("Get all links", func(t *testing.T) {
		path := "/v1/links"
		target := &map[string][]*models.Link{}
		err := baseTestingAPIRequest(router, path, target)
		if err != nil {
			t.Error(err)
		}
	})
}

func baseTestingAPIRequest(router *gin.Engine, urlPath string, target interface{}) error {
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		return err
	}
	router.ServeHTTP(w, req)

	body := w.Body.Bytes()
	err = json.Unmarshal(body, target)
	if err != nil {
		return fmt.Errorf("Error unmarshaling JSON: %s\nReponse text: %s", err.Error(), string(body))
	}

	return nil
}
