package services

import (
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"commitment-stake/config"
)

const (
	domainName    = "CommitmentPool"
	domainVersion = "1"
	QRValidity    = 5 * time.Minute
)

var (
	domainTypeHash  = crypto.Keccak256([]byte("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"))
	checkInTypeHash = crypto.Keccak256([]byte("CheckIn(uint256 eventId,uint64 deadline)"))
)

// Payload này được mã hoá thành QR cho người tham dự quét
type CheckInPayload struct {
	EventID  uint64 `json:"eventId"`
	Deadline uint64 `json:"deadline"`
	Sig      string `json:"sig"` // 0x + 65 byte r||s||v, v = 27/28
}

func pad32(b []byte) []byte { return common.LeftPadBytes(b, 32) }

func domainSeparator(chainID int64, pool common.Address) []byte {
	return crypto.Keccak256(
		domainTypeHash,
		crypto.Keccak256([]byte(domainName)),
		crypto.Keccak256([]byte(domainVersion)),
		pad32(big.NewInt(chainID).Bytes()),
		pad32(pool.Bytes()),
	)
}

// CheckInDigest phải trùng với hàm checkInDigest(id, deadline) trong contract
func CheckInDigest(chainID int64, pool common.Address, eventID, deadline uint64) []byte {
	structHash := crypto.Keccak256(
		checkInTypeHash,
		pad32(new(big.Int).SetUint64(eventID).Bytes()),
		pad32(new(big.Int).SetUint64(deadline).Bytes()),
	)
	return crypto.Keccak256(
		[]byte{0x19, 0x01},
		domainSeparator(chainID, pool),
		structHash,
	)
}

// SignCheckIn ký payload check-in có hiệu lực 5 phút kể từ bây giờ
func SignCheckIn(eventID uint64) (*CheckInPayload, error) {
	c := config.Chain
	if c == nil {
		return nil, errors.New("chưa gọi config.LoadChainConfig()")
	}
	deadline := uint64(time.Now().Add(QRValidity).Unix())
	digest := CheckInDigest(c.ChainID, c.PoolAddress, eventID, deadline)

	sig, err := crypto.Sign(digest, c.OrganiserKey)
	if err != nil {
		return nil, err
	}
	sig[64] += 27 // Solidity (ECDSA.recover) cần v = 27/28

	return &CheckInPayload{
		EventID:  eventID,
		Deadline: deadline,
		Sig:      hexutil.Encode(sig),
	}, nil
}

// RecoverSigner khôi phục địa chỉ người ký, dùng để tự kiểm tra
func RecoverSigner(chainID int64, pool common.Address, eventID, deadline uint64, sigHex string) (common.Address, error) {
	sig, err := hexutil.Decode(sigHex)
	if err != nil {
		return common.Address{}, err
	}
	if len(sig) != 65 {
		return common.Address{}, errors.New("chữ ký phải dài 65 byte")
	}
	if sig[64] >= 27 {
		sig[64] -= 27
	}
	pub, err := crypto.SigToPub(CheckInDigest(chainID, pool, eventID, deadline), sig)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pub), nil
}