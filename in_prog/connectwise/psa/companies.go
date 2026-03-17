package psa

import (
	"context"
	"fmt"
)

func companyIDEndpoint(companyID int) string {
	return fmt.Sprintf("company/companies/%d", companyID)
}

func (c *Client) PostCompany(ctx context.Context, company *Company) (*Company, error) {
	return Post[Company](ctx, c, "company/companies", company)
}

func (c *Client) ListCompanies(ctx context.Context, params map[string]string) ([]Company, error) {
	return GetMany[Company](ctx, c, "company/companies", params)
}

func (c *Client) GetCompany(ctx context.Context, companyID int, params map[string]string) (*Company, error) {
	return GetOne[Company](ctx, c, companyIDEndpoint(companyID), params)
}

func (c *Client) PutCompany(ctx context.Context, companyID int, company *Company) (*Company, error) {
	return Put[Company](ctx, c, companyIDEndpoint(companyID), company)
}

func (c *Client) PatchCompany(ctx context.Context, companyID int, patchOps []PatchOp) (*Company, error) {
	return Patch[Company](ctx, c, companyIDEndpoint(companyID), patchOps)
}

func (c *Client) DeleteCompany(ctx context.Context, companyID int) error {
	return Delete(ctx, c, companyIDEndpoint(companyID))
}
