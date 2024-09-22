package compute

import (
	"context"
	"strings"

	"github.com/skyris/KVlight/pkg/interfaces"
	"github.com/skyris/KVlight/pkg/issues"
	"github.com/skyris/KVlight/pkg/types"
)

type Compute struct{}

var _ interfaces.Computer = &Compute{}

func (c *Compute) Parse(_ context.Context, in string) ([]string, error) {
	in = strings.Trim(in, " ")
	args := strings.Fields(in)
	if err := Validate(args); err != nil {
		return nil, err
	}
	args[0] = strings.ToUpper(args[0])

	return args, nil
}

func NewCompute() *Compute {
	return &Compute{}
}

func Validate(args []string) error {
	amount := len(args)
	if amount < 2 || amount > 3 {
		return issues.ErrInvalidArgumentCount
	}
	command := strings.ToUpper(args[0])
	if amount == 2 && (command == types.CommandDEL || command == types.CommandGET) {
		return nil
	} else if amount == 3 && command == types.CommandSET {
		return nil
	}
	return issues.ErrInvalidCommand
}
