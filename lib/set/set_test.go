package set

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	s := NewSet[int](1, 2, 3, 1)
	unique := s.Values()
	require.Equal(t, len(unique), 3)
}
