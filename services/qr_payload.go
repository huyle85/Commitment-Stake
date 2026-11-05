package services

import (
	"encoding/json"
	"errors"
	"time"
)

func EncodeQRContent(payload *CheckInPayload) (string, error) {
	if payload == nil {
		return "", errors.New("payload rỗng")
	}
	data, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func DecodeQRContent(content string) (*CheckInPayload, error) {
	var payload CheckInPayload
	err := json.Unmarshal([]byte(content), &payload)
	if err != nil {
		return nil, err
	}
	if payload.EventID == 0 {
		return nil, errors.New("thiếu eventId")
	}
	if payload.Sig == "" {
		return nil, errors.New("thiếu chữ ký")
	}
	return &payload, nil
}

func IsExpired(payload *CheckInPayload, now time.Time) bool {
	return uint64(now.Unix()) > payload.Deadline
}