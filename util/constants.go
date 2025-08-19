// constants.go
package util

// File paths and locations for various I2P router installations
// All constants consolidated from non-platform-specific code

var (
	// I2CP_HOST is the i2cp host
	// Moved from: constant.go (non-platform-specific)
	I2CP_HOST string

	// I2CP_PORT is the i2cp port
	// Moved from: constant.go (non-platform-specific)
	I2CP_PORT string

	// WINDOWS_DEFAULT_LOCATION is the default location of the I2P router on Windows
	// Moved from: constant.go (non-platform-specific)
	WINDOWS_DEFAULT_LOCATION string = `C:\\Program Files\i2p\i2psvc.exe`

	// WINDOWS_ZERO_NSIS_DEFAULT_LOCATION is the default location of the I2P Zero router on Windows
	// Moved from: constant.go (non-platform-specific)
	WINDOWS_ZERO_NSIS_DEFAULT_LOCATION string = `C:\\Program Files\I2P-Zero\router\i2p-zero.exe`

	// I2PD_WINDOWS_DEFAULT_LOCATION is the default location of the I2Pd router on Windows
	// Moved from: constant.go (non-platform-specific)
	I2PD_WINDOWS_DEFAULT_LOCATION string = `C:\\Program Files\I2Pd\i2pd.exe`

	// LINUX_SYSTEM_LOCATION is the default location of the I2P router on Linux
	// Moved from: constant.go (non-platform-specific)
	LINUX_SYSTEM_LOCATION []string = []string{"/usr/bin/i2prouter", "/usr/sbin/i2prouter"}

	// I2PD_LINUX_SYSTEM_LOCATION is the default location of the I2Pd router on Linux
	// Moved from: constant.go (non-platform-specific)
	I2PD_LINUX_SYSTEM_LOCATION []string = []string{"/usr/sbin/i2pd", "/usr/bin/i2pd"}

	// I2P_ASUSER_HOME_LOCATION This is the path to the default I2P config directory when running as a user
	// Moved from: constant.go (non-platform-specific)
	I2P_ASUSER_HOME_LOCATION string = Inithome(Home())

	// HOME_DIRECTORY_LOCATION can use config/run from a user's $HOME directory, this is the path to that router
	// Moved from: constant.go (non-platform-specific)
	HOME_DIRECTORY_LOCATION string = Inithome("/i2p/i2prouter")

	// OSX_DEFAULT_LOCATION is the default location of the I2P router on OSX
	// Moved from: constant.go (non-platform-specific)
	OSX_DEFAULT_LOCATION string = Inithome("/Library/Application Support/i2p/clients.config")
)
