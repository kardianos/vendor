// +build !windows

package context

import (
	"fmt"
	"strconv"
)

const gorootLookFor = `GOROOT=`

func gorootFromGoEnv(s string) (string, error) {
	goroot, err := strconv.Unquote(s)
	if err != nil {
		return "", fmt.Errorf("Failed to unquote GOROOT: %v", err)
	}
	return goroot, nil
}
