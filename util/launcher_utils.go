// launcher_utils.go
package util

import (
	"log"
)

// Launcher Detection Functions
// All functions moved from: launch.go (non-platform-specific)

// checkI2PdRouterInstallation checks for i2pd router installations on Windows and Linux.
func checkI2PdRouterInstallation() bool {
	if CheckFileExists(I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return true
	}
	if CheckFileExists(I2PD_LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2pd router detected")
		return true
	}
	if CheckFileExists(I2PD_LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2pd router detected")
		return true
	}
	return false
}

// checkJavaI2PWindowsInstallation checks for Java I2P router installations on Windows.
func checkJavaI2PWindowsInstallation() bool {
	if CheckFileExists(WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return true
	}
	if CheckFileExists(WINDOWS_ZERO_NSIS_DEFAULT_LOCATION) {
		log.Println("Windows i2p zero router detected")
		return true
	}
	return false
}

// checkJavaI2PLinuxInstallation checks for Java I2P router installations on Linux.
func checkJavaI2PLinuxInstallation() bool {
	if CheckFileExists(LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2p router detected")
		return true
	}
	if CheckFileExists(LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2p router detected")
		return true
	}
	return false
}

// checkUserDirectoryInstallation checks for I2P router installations in user directories.
func checkUserDirectoryInstallation() bool {
	if CheckFileExists(HOME_DIRECTORY_LOCATION) {
		log.Println("Linux i2p router detected")
		return true
	}
	if CheckFileExists(OSX_DEFAULT_LOCATION) {
		log.Println("OSX i2p router detected")
		return true
	}
	return false
}

// CheckI2PIsInstalledDefaultLocation looks in various locations for the
// presence of an I2P router.
func CheckI2PIsInstalledDefaultLocation() (bool, error) {
	if checkI2PdRouterInstallation() {
		return true, nil
	}
	if checkJavaI2PWindowsInstallation() {
		return true, nil
	}
	if checkJavaI2PLinuxInstallation() {
		return true, nil
	}
	if checkUserDirectoryInstallation() {
		return true, nil
	}
	return false, nil
}
