package chore

import (
	"github.com/hashicorp/go-multierror"
	"github.com/magefile/mage/sh"
)

// BatchRunV run commands in batch
func BatchRunV(commandList [][]string) error {
	var mErr error
	for _, args := range commandList {
		if err := sh.RunV(args[0], args[1:]...); err != nil {
			mErr = multierror.Append(mErr, err)
			continue
		}
	}
	return mErr
}
