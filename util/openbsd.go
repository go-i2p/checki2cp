//go:build openbsd

package util

func IsHeadless() (bool, error) {
	return true, nil
}
