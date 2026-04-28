package joinclient

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gravitational/teleport"
	"github.com/gravitational/teleport/lib/defaults"
)

func TestCandidateProxyAddresses(t *testing.T) {
	t.Skip("no candidate fallback behavior in join path")
}

func TestProxyJoinErrorHint(t *testing.T) {
	t.Parallel()

	msg := proxyJoinErrorHint("proxy.example.com:3080")
	require.Contains(t, msg, "building proxy client using proxy.example.com:3080")
	require.Contains(t, msg, "set proxy_server to proxy.example.com:443")

	msg = proxyJoinErrorHint("proxy.example.com:4443")
	require.Contains(t, msg, "building proxy client using proxy.example.com:4443")
	require.NotContains(t, msg, "set proxy_server to")

	// Keep test tied to defaults explicitly so changes are intentional.
	require.Equal(t, 3080, defaults.HTTPListenPort)
	require.Equal(t, 443, teleport.StandardHTTPSPort)
}
