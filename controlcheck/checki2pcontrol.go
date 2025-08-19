package checki2pcontrol

import (
	"github.com/go-i2p/go-i2pcontrol"
)

// CheckI2PControlEcho attempts a connection and an echo command on it.
// it returns true if the command is successful and false, with an error,
// if not.
func CheckI2PControlEcho(host, port, password, path string) (bool, error) {
	host = normalizeHost(host)
	port = normalizePort(port)

	// Try primary connection configuration
	if err := attemptI2PControlConnection(host, port, password, path); err == nil {
		return true, nil
	}

	// Try fallback configuration
	if err := attemptI2PControlConnection(host, "7657", password, "jsonrpc"); err == nil {
		return true, nil
	}

	// Both attempts failed
	return false, attemptI2PControlConnection(host, "7657", password, "jsonrpc")
}

// normalizeHost returns the default host if the provided host is empty.
func normalizeHost(host string) string {
	if host == "" {
		return "127.0.0.1"
	}
	return host
}

// normalizePort returns the default port if the provided port is empty.
func normalizePort(port string) string {
	if port == "" {
		return "7650"
	}
	return port
}

// attemptI2PControlConnection tries to establish an I2PControl connection and perform an echo test.
func attemptI2PControlConnection(host, port, password, path string) error {
	i2pcontrol.Initialize(host, port, path)

	if _, err := i2pcontrol.Authenticate(password); err != nil {
		return err
	}

	if _, err := i2pcontrol.Echo("Hello I2PControl"); err != nil {
		return err
	}

	return nil
}

// GetDefaultI2PControlPath probes default locations for the I2PControl API, returning
// either a working I2PControl API and no error, or the defaults of the embedded router
// and an error
func GetDefaultI2PControlPath(password ...string) (string, string, string, error) {
	host := "127.0.0.1"
	pass := extractPassword(password)

	// Try primary configuration
	port, path := "7650", ""
	if err := attemptI2PControlConnection(host, port, pass, path); err == nil {
		return host, port, path, nil
	}

	// Try fallback configuration
	port, path = "7657", "jsonrpc"
	err := attemptI2PControlConnection(host, port, pass, path)
	return host, port, path, err
}

// extractPassword returns the first password from the variadic parameter or default password.
func extractPassword(password []string) string {
	if len(password) > 0 {
		return password[0]
	}
	return "itoopie"
}
