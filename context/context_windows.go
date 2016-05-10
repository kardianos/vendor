// +build windows

package context

const gorootLookFor = `set GOROOT=`

func gorootFromGoEnv(s string) (string, error) {
	return s, nil
}
