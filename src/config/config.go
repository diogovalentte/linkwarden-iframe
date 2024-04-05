package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	LinkwardenAddress string
	LinkwardenToken   string
)

func SetConfigs(filePath string) error {
	if filePath != "" {
		err := godotenv.Load(filePath)
		if err != nil {
			return err
		}
	}

	LinkwardenAddress = os.Getenv("LINKWARDEN_ADDRESS")
	LinkwardenToken = os.Getenv("LINKWARDEN_TOKEN")

	if LinkwardenAddress == "" || LinkwardenToken == "" {
		return fmt.Errorf("LINKWARDEN_ADDRESS and LINKWARDEN_TOKEN variables should be set")
	}

	if strings.HasSuffix(LinkwardenAddress, "/") {
		LinkwardenAddress = LinkwardenAddress[:len(LinkwardenAddress)-1]
	}

	return nil
}
