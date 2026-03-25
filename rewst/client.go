package rewst

import (
	"fmt"
	"strings"
)

type Config struct {
	WebhookSecret  string
	ListOrgsURL    string
	UpsertVarURL   string
	BackupVarsURL  string
}

type Client struct {
	wc            *WebhookClient
	listOrgsURL   string
	upsertVarURL  string
	backupVarsURL string
}

func (c *Client) WebhookClient() *WebhookClient {
	return c.wc
}

func NewClient(cfg Config) (*Client, error) {
	var missing []string
	if cfg.WebhookSecret == "" {
		missing = append(missing, "WebhookSecret")
	}
	if cfg.ListOrgsURL == "" {
		missing = append(missing, "ListOrgsURL")
	}
	if cfg.UpsertVarURL == "" {
		missing = append(missing, "UpsertVarURL")
	}
	if len(missing) > 0 {
		return nil, fmt.Errorf("missing rewst config fields: %s", strings.Join(missing, ", "))
	}

	return &Client{
		wc:            NewWebhookClient(cfg.WebhookSecret),
		listOrgsURL:   cfg.ListOrgsURL,
		upsertVarURL:  cfg.UpsertVarURL,
		backupVarsURL: cfg.BackupVarsURL,
	}, nil
}
