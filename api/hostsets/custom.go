package hostsets

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/hashicorp/boundary/api"
	"github.com/kr/pretty"
)

func (c *Client) Create2(ctx context.Context, hostCatalogId string, opt ...Option) (*HostSet, *api.Error, error) {
	if hostCatalogId == "" {
		return nil, nil, fmt.Errorf("empty hostCatalogId value passed into Create request")
	}
	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts.postMap["host_catalog_id"] = hostCatalogId
	req, err := c.client.NewRequest(ctx, "POST", "host-sets", opts.postMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating Create request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during Create call: %w", err)
	}

	target := new(HostSet)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding Create response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *Client) Read2(ctx context.Context, hostSetId string, opt ...Option) (*HostSet, *api.Error, error) {
	if hostSetId == "" {
		return nil, nil, fmt.Errorf("empty hostSetId value passed into Read request")
	}

	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("host-sets/%s", hostSetId), nil, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating Read request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during Read call: %w", err)
	}

	target := new(HostSet)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding Read response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *Client) Update2(ctx context.Context, hostSetId string, version uint32, opt ...Option) (*HostSet, *api.Error, error) {
	if hostSetId == "" {
		return nil, nil, fmt.Errorf("empty hostSetId value passed into Update request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, nil, errors.New("zero version number passed into Update request and automatic versioning not specified")
		}
		existingTarget, existingApiErr, existingErr := c.Read2(ctx, hostSetId, opt...)
		if existingErr != nil {
			return nil, nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingApiErr != nil {
			return nil, nil, fmt.Errorf("error from controller when performing initial check-and-set read: %s", pretty.Sprint(existingApiErr))
		}
		if existingTarget == nil {
			return nil, nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Version
	}

	opts.postMap["version"] = version

	req, err := c.client.NewRequest(ctx, "PATCH", fmt.Sprintf("host-sets/%s", hostSetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating Update request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during Update call: %w", err)
	}

	target := new(HostSet)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding Update response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *Client) Delete2(ctx context.Context, hostSetId string, opt ...Option) (bool, *api.Error, error) {
	if hostSetId == "" {
		return false, nil, fmt.Errorf("empty hostSetId value passed into Delete request")
	}

	if c.client == nil {
		return false, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("host-sets/%s", hostSetId), nil, apiOpts...)
	if err != nil {
		return false, nil, fmt.Errorf("error creating Delete request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return false, nil, fmt.Errorf("error performing client request during Delete call: %w", err)
	}

	type deleteResponse struct {
		Existed bool
	}
	target := &deleteResponse{}
	apiErr, err := resp.Decode(target)
	if err != nil {
		return false, nil, fmt.Errorf("error decoding Delete response: %w", err)
	}
	if apiErr != nil {
		return false, apiErr, nil
	}
	return target.Existed, apiErr, nil
}

func (c *Client) List2(ctx context.Context, hostCatalogId string, opt ...Option) ([]*HostSet, *api.Error, error) {
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	req, err := c.client.NewRequest(ctx, "GET", "host-sets", nil, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating List request: %w", err)
	}

	opts.queryMap["host_catalog_id"] = hostCatalogId

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during List call: %w", err)
	}

	type listResponse struct {
		Items []*HostSet
	}
	target := &listResponse{}
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding List response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target.Items, apiErr, nil
}

func (c *Client) AddHosts2(ctx context.Context, hostSetId string, version uint32, hostIds []string, opt ...Option) (*HostSet, *api.Error, error) {
	if hostSetId == "" {
		return nil, nil, fmt.Errorf("empty hostSetId value passed into AddHosts request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, nil, errors.New("zero version number passed into AddHosts request")
		}
		existingTarget, existingApiErr, existingErr := c.Read2(ctx, hostSetId, opt...)
		if existingErr != nil {
			return nil, nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingApiErr != nil {
			return nil, nil, fmt.Errorf("error from controller when performing initial check-and-set read: %s", pretty.Sprint(existingApiErr))
		}
		if existingTarget == nil {
			return nil, nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Version
	}

	opts.postMap["version"] = version

	if len(hostIds) > 0 {
		opts.postMap["host_ids"] = hostIds
	}

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("host-sets/%s:add-hosts", hostSetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating AddHosts request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during AddHosts call: %w", err)
	}

	target := new(HostSet)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding AddHosts response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *Client) SetHosts2(ctx context.Context, hostSetId string, version uint32, hostIds []string, opt ...Option) (*HostSet, *api.Error, error) {
	if hostSetId == "" {
		return nil, nil, fmt.Errorf("empty hostSetId value passed into SetHosts request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, nil, errors.New("zero version number passed into SetHosts request")
		}
		existingTarget, existingApiErr, existingErr := c.Read2(ctx, hostSetId, opt...)
		if existingErr != nil {
			return nil, nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingApiErr != nil {
			return nil, nil, fmt.Errorf("error from controller when performing initial check-and-set read: %s", pretty.Sprint(existingApiErr))
		}
		if existingTarget == nil {
			return nil, nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Version
	}

	opts.postMap["version"] = version

	if len(hostIds) > 0 {
		opts.postMap["host_ids"] = hostIds
	} else if hostIds != nil {
		// In this function, a non-nil but empty list means clear out
		opts.postMap["host_ids"] = nil
	}

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("host-sets/%s:set-hosts", hostSetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating SetHosts request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during SetHosts call: %w", err)
	}

	target := new(HostSet)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding SetHosts response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}

func (c *Client) RemoveHosts2(ctx context.Context, hostSetId string, version uint32, hostIds []string, opt ...Option) (*HostSet, *api.Error, error) {
	if hostSetId == "" {
		return nil, nil, fmt.Errorf("empty hostSetId value passed into RemoveHosts request")
	}
	if c.client == nil {
		return nil, nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	apiOpts = append(apiOpts, api.WithNewStyle())

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, nil, errors.New("zero version number passed into RemoveHosts request")
		}
		existingTarget, existingApiErr, existingErr := c.Read2(ctx, hostSetId, opt...)
		if existingErr != nil {
			return nil, nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingApiErr != nil {
			return nil, nil, fmt.Errorf("error from controller when performing initial check-and-set read: %s", pretty.Sprint(existingApiErr))
		}
		if existingTarget == nil {
			return nil, nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Version
	}

	opts.postMap["version"] = version

	if len(hostIds) > 0 {
		opts.postMap["host_ids"] = hostIds
	}

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("host-sets/%s:remove-hosts", hostSetId), opts.postMap, apiOpts...)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating RemoveHosts request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("error performing client request during RemoveHosts call: %w", err)
	}

	target := new(HostSet)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, nil, fmt.Errorf("error decoding RemoveHosts response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr, nil
	}
	return target, apiErr, nil
}
