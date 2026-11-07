package virtualnetwork

import (
	"testing"

	"github.com/containers/gvisor-tap-vsock/pkg/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gvisor.dev/gvisor/pkg/tcpip/stack"
)

func TestCreatePortsForwarderCopiesDisableSourceIPPropagation(t *testing.T) {
	for _, disable := range []bool{false, true} {
		configuration := &types.Configuration{DisableSourceIPPropagation: disable}
		s := stack.New(stack.Options{})
		t.Cleanup(s.Destroy)
		fw, err := createPortsForwarder(configuration, s)
		require.NoError(t, err)
		assert.Equal(t, disable, fw.DisableSourceIPPropagation)
	}
}
