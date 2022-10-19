package credential

import (
	"testing"

	"github.com/hashicorp/boundary/internal/iam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GetOpts(t *testing.T) {
	t.Parallel()
	t.Run("WithUserId", func(t *testing.T) {
		opts := getDefaultOptions()
		assert.Empty(t, opts.withUserId)
		opts, err := getOpts(WithUserId("foo"))
		require.NoError(t, err)
		assert.Equal(t, "foo", opts.withUserId)
	})
	t.Run("WithIamRepoFn", func(t *testing.T) {
		opts := getDefaultOptions()
		assert.Empty(t, opts.withIamRepoFn)
		opts, err := getOpts(WithIamRepoFn(func() (*iam.Repository, error) { return nil, nil }))
		require.NoError(t, err)
		assert.NotEmpty(t, opts.withIamRepoFn)
	})
}
