package threatdown

import (
	"context"
	"os"
	"testing"
)

func TestSearchEndpoints(t *testing.T) {
	c := newTestClient(t)
	ctx := context.Background()

	accountID := os.Getenv("THREATDOWN_NEBULA_ACCOUNT_ID")
	if accountID == "" {
		t.Skip("skipping: THREATDOWN_NEBULA_ACCOUNT_ID must be set")
	}

	endpoints, err := c.SearchEndpoints(ctx, accountID, map[string]string{
		"is_deleted": "false",
	})
	if err != nil {
		t.Fatalf("SearchEndpoints: %v", err)
	}
	t.Logf("got %d endpoints", len(endpoints))
}
