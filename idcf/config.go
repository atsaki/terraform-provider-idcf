package idcf

import (
	"fmt"
	"log"

	"github.com/atsaki/go-idcf/dns"
)

type Config struct {
	APIKey    string
	SecretKey string
}

type IDCFClient struct {
	dns *dns.Client
}

func (c *Config) Client() (client *IDCFClient, err error) {
	client = new(IDCFClient)

	log.Printf("[INFO] Initializing IDCF Client")
	client.dns, err = dns.NewClient(c.APIKey, c.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("Failed to set up DNS client: %s", err)
	}
	log.Printf("[INFO] IDCF Client successfully initialized")

	return
}
