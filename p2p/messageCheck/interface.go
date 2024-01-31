package messagecheck

import "github.com/bhagyaraj1208117/andes-core-go/core"

type p2pSigner interface {
	Verify(payload []byte, pid core.PeerID, signature []byte) error
	IsInterfaceNil() bool
}
