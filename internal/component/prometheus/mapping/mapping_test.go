package mapping

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidator(t *testing.T) {
	args := Arguments{}
	err := args.Validate()
	require.NoError(t, err)
}
