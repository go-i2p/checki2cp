//go:build openbsd

package util

func IsHeadless() bool {
	return true
}
