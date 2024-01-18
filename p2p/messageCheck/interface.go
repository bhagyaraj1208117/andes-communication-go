package messagecheck

import "github.com/bhagyaraj1208117/andes-abc-1/core"

type p2pSigner interface {
	Verify(payload []byte, pid core.PeerID, signature []byte) error
	IsInterfaceNil() bool
}
