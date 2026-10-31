package services

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"commitment-stake/config"
)

func TestSignAndRecover(t *testing.T) {
	key, _ := crypto.GenerateKey()
	pool := common.HexToAddress("0x0000000000000000000000000000000000000001")
	config.Chain = &config.ChainConfig{
		ChainID:       43113,
		PoolAddress:   pool,
		OrganiserKey:  key,
		OrganiserAddr: crypto.PubkeyToAddress(key.PublicKey),
	}

	p, err := SignCheckIn(1)
	if err != nil {
		t.Fatal(err)
	}
	got, err := RecoverSigner(43113, pool, p.EventID, p.Deadline, p.Sig)
	if err != nil {
		t.Fatal(err)
	}
	if got != config.Chain.OrganiserAddr {
		t.Fatalf("sai người ký: %s != %s", got, config.Chain.OrganiserAddr)
	}
	t.Logf("deadline=%d digest=%s", p.Deadline, hexutil.Encode(CheckInDigest(43113, pool, p.EventID, p.Deadline)))
}