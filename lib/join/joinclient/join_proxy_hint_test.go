package joinclient

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/gravitational/teleport/lib/utils"
)

func TestJoinNewIncludesProxyHintOnConnectionFailure(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(t.Context(), 500*time.Millisecond)
	defer cancel()

	_, err := joinNew(ctx, JoinParams{
		ProxyServer: utils.NetAddr{Addr: "203.0.113.1:3080", AddrNetwork: "tcp"},
		Log:         slog.Default(),
	})
	require.Error(t, err)

	t.Logf("joinNew error: %v", err)
	require.Contains(t, err.Error(), "set proxy_server to 203.0.113.1:443")
}
