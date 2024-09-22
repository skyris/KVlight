package delivery_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/skyris/KVlight/internal/delivery"
)

func TestGetRequest(t *testing.T) {
	// Test case 1: Valid input
	input := "test command\n"
	expected := "test command"
	reader := strings.NewReader(input)
	d := delivery.NewStdinDelivery(reader, nil)

	ctx := context.Background()
	result, err := d.GetRequest(ctx)
	if err != nil {
		t.Fatalf("GetRequest failed: %v", err)
	}
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}

	// Test case 2: No input (EOF)
	result, err = d.GetRequest(ctx)
	if !errors.Is(err, io.EOF) {
		t.Fatalf("expected io.EOF, got %v", err)
	}
	if result != "" {
		t.Errorf("expected empty string, got %q", result)
	}
}

func TestSendResponse(t *testing.T) {
	var bufIn bytes.Buffer
	var bufOut bytes.Buffer
	d := delivery.NewStdinDelivery(&bufIn, &bufOut)

	ctx := context.Background()

	// Test case 1: Sending a message without an error
	err := d.SendResponse(ctx, "test message", nil)
	if err != nil {
		t.Fatalf("SendResponse failed: %v", err)
	}
	expected := "msg: test message\n"
	if bufOut.String() != expected {
		t.Errorf("expected %q, got %q", expected, bufOut.String())
	}

	// Clear buffer for next test
	bufOut.Reset()

	// Test case 2: Sending an error
	sendErr := errors.New("something went wrong")
	err = d.SendResponse(ctx, "", sendErr)
	if err != nil {
		t.Fatalf("SendResponse failed: %v", err)
	}
	expected = "error: something went wrong\n"
	if bufOut.String() != expected {
		t.Errorf("expected %q, got %q", expected, bufOut.String())
	}

	// Clear buffer for next test
	bufOut.Reset()

	// Test case 3: Sending "done" when no message or error
	err = d.SendResponse(ctx, "", nil)
	if err != nil {
		t.Fatalf("SendResponse failed: %v", err)
	}
	expected = "done\n"
	if bufOut.String() != expected {
		t.Errorf("expected %q, got %q", expected, bufOut.String())
	}
}
