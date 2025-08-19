//go:build freebsd

package util

func IsHeadless() (bool, error) {
	return true, nil
}
