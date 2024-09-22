package compute_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/skyris/KVlight/internal/compute"
	"github.com/skyris/KVlight/pkg/issues"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      []string
		expectErr bool
	}{
		{
			name:      "Empty input",
			input:     "",
			want:      nil,
			expectErr: true,
		},
		{
			name:      "Invalid command",
			input:     "invalidCommand",
			want:      nil,
			expectErr: true,
		},
		{
			name:      "Valid input with multiple words",
			input:     "set key value",
			want:      []string{"SET", "key", "value"},
			expectErr: false,
		},
		{
			name:      "Valid input with single word",
			input:     "get key",
			want:      []string{"GET", "key"},
			expectErr: false,
		},
		{
			name:      "Valid input with single word and too many spaces",
			input:     " get  key ",
			want:      []string{"GET", "key"},
			expectErr: false,
		},
		{
			name:      "Valid delete command",
			input:     "del key",
			want:      []string{"DEL", "key"},
			expectErr: false,
		},
	}

	computer := &compute.Compute{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := computer.Parse(context.Background(), tt.input)

			if (err != nil) != tt.expectErr {
				t.Errorf("got %q want %v", err, tt.expectErr)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %q want %q", got, tt.want)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr error
	}{
		{
			"Valid DELETE with 2 args",
			[]string{"del", "arg1"},
			nil,
		},
		{
			"Valid GET with 2 args",
			[]string{"get", "arg1"},
			nil,
		},
		{
			"Valid ADD with 3 args",
			[]string{"set", "arg1", "arg2"},
			nil,
		},
		{
			"Invalid command",
			[]string{"update", "arg1"},
			issues.ErrInvalidCommand,
		},
		{
			"Too few arguments",
			[]string{"get"},
			issues.ErrInvalidArgumentCount,
		},
		{
			"Too many arguments",
			[]string{"add", "arg1", "arg2", "arg3"},
			issues.ErrInvalidArgumentCount,
		},
		{
			"Valid ADD, wrong argument count",
			[]string{"add", "arg1"},
			issues.ErrInvalidCommand,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := compute.Validate(tt.args); !errors.Is(err, tt.wantErr) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
