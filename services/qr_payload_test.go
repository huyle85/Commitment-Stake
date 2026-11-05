package services

import (
	"testing"
	"time"
)

func TestEncodeDecodeQRContent(t *testing.T) {
	original := &CheckInPayload{EventID: 7, Deadline: 1793456807, Sig: "0xabc"}
	content, err := EncodeQRContent(original)
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := DecodeQRContent(content)
	if err != nil {
		t.Fatal(err)
	}
	if decoded.EventID != original.EventID {
		t.Fatal("eventId không khớp")
	}
	if decoded.Deadline != original.Deadline {
		t.Fatal("deadline không khớp")
	}
	if decoded.Sig != original.Sig {
		t.Fatal("chữ ký không khớp")
	}
}

func TestDecodeQRContentRejectsMissingFields(t *testing.T) {
	_, err := DecodeQRContent(`{"eventId":0,"deadline":1,"sig":""}`)
	if err == nil {
		t.Fatal("phải báo lỗi khi thiếu trường")
	}
}

func TestIsExpired(t *testing.T) {
	payload := &CheckInPayload{EventID: 1, Deadline: 1000, Sig: "0x1"}
	if IsExpired(payload, time.Unix(999, 0)) {
		t.Fatal("chưa hết hạn mà báo hết hạn")
	}
	if !IsExpired(payload, time.Unix(1001, 0)) {
		t.Fatal("đã hết hạn mà báo chưa")
	}
}