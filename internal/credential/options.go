package credential

import (
	"github.com/hashicorp/boundary/internal/iam"
)

// getOpts - iterate the inbound Options and return a struct
func getOpts(opt ...Option) (*options, error) {
	opts := getDefaultOptions()
	for _, o := range opt {
		if o == nil {
			continue
		}
		if err := o(opts); err != nil {
			return nil, err
		}
	}
	return opts, nil
}

// Option - how Options are passed as arguments.
type Option func(*options) error

// options = how options are represented
type options struct {
	withIamRepoFn iam.IamRepoFactory
	withUserId    string
}

func getDefaultOptions() *options {
	return &options{}
}

// WithIamRepoFn provides a way to pass in a repo function to use in looking up
// templated data in credential library definitions, if needed
func WithIamRepoFn(with iam.IamRepoFactory) Option {
	return func(o *options) error {
		o.withIamRepoFn = with
		return nil
	}
}

// WithUserId provides a way to pass in the user ID for use in templating (or
// looking up for templating)
func WithUserId(with string) Option {
	return func(o *options) error {
		o.withUserId = with
		return nil
	}
}
