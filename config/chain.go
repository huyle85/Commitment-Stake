package config

import (
	"crypto/ecdsa"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type ChainConfig struct {
	RPCURL        string
	ChainID       int64
	PoolAddress   common.Address
	USDCAddress   common.Address
	OrganiserKey  *ecdsa.PrivateKey
	OrganiserAddr common.Address
}

// Chain được gán sau khi gọi LoadChainConfig()
var Chain *ChainConfig

func env(key string) (string, error) {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return "", fmt.Errorf("thiếu biến %s trong .env", key)
	}
	return v, nil
}

func addrEnv(key string) (common.Address, error) {
	v, err := env(key)
	if err != nil {
		return common.Address{}, err
	}
	if !common.IsHexAddress(v) {
		return common.Address{}, fmt.Errorf("%s không phải địa chỉ hợp lệ", key)
	}
	return common.HexToAddress(v), nil
}

func LoadChainConfig() error {
	rpc, err := env("RPC_URL")
	if err != nil {
		return err
	}
	chainStr, err := env("CHAIN_ID")
	if err != nil {
		return err
	}
	chainID, err := strconv.ParseInt(chainStr, 10, 64)
	if err != nil {
		return fmt.Errorf("CHAIN_ID không hợp lệ: %w", err)
	}
	pool, err := addrEnv("POOL_ADDRESS")
	if err != nil {
		return err
	}
	usdc, err := addrEnv("USDC_ADDRESS")
	if err != nil {
		return err
	}
	keyHex, err := env("ORGANISER_PRIVATE_KEY")
	if err != nil {
		return err
	}
	key, err := crypto.HexToECDSA(strings.TrimPrefix(keyHex, "0x"))
	if err != nil {
		return fmt.Errorf("ORGANISER_PRIVATE_KEY không hợp lệ: %w", err)
	}

	Chain = &ChainConfig{
		RPCURL:        rpc,
		ChainID:       chainID,
		PoolAddress:   pool,
		USDCAddress:   usdc,
		OrganiserKey:  key,
		OrganiserAddr: crypto.PubkeyToAddress(key.PublicKey),
	}
	return nil
}