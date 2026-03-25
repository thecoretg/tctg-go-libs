package rewst

import "context"

func (c *Client) BackUpOrgVars(ctx context.Context, orgID string) error {
	p := struct {
		OrgID string `json:"org_id"`
	}{OrgID: orgID}

	return Send(ctx, c.wc, c.backupVarsURL, p)
}
