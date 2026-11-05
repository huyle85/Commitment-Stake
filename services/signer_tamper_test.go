package services

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"commitment-stake/config"
)

func TestSignatureBoundToEventAndChain(t *testing.T) {
	key, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	pool := common.HexToAddress("0x0000000000000000000000000000000000000001")
	organiser := crypto.PubkeyToAddress(key.PublicKey)
	config.Chain = &config.ChainConfig{
		ChainID:       43113,
		PoolAddress:   pool,
		OrganiserKey:  key,
		OrganiserAddr: organiser,
	}

	payload, err := SignCheckIn(1)
	if err != nil {
		t.Fatal(err)
	}

	otherEvent, err := RecoverSigner(43113, pool, 2, payload.Deadline, payload.Sig)
	if err == nil && otherEvent == organiser {
		t.Fatal("đổi eventId mà chữ ký vẫn hợp lệ")
	}

	otherChain, err := RecoverSigner(1, pool, 1, payload.Deadline, payload.Sig)
	if err == nil && otherChain == organiser {
		t.Fatal("đổi chainId mà chữ ký vẫn hợp lệ")
	}

	otherDeadline, err := RecoverSigner(43113, pool, 1, payload.Deadline+1, payload.Sig)
	if err == nil && otherDeadline == organiser {
		t.Fatal("đổi deadline mà chữ ký vẫn hợp lệ")
	}
}