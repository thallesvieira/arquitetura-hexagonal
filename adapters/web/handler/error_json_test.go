package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_JsonError(t *testing.T) {
	msg := "Hello json"
	result := JsonError(msg)
	require.Equal(t, []byte(`{"message":"Hello json"}`), result)
}
