package util

import (
	"fmt"
	"log"
)

// checkI2PDRouterLocations checks for I2PD (i2pd) router installations
func checkI2PDRouterLocations() (string, bool) {
	if CheckFileExists(I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return I2PD_WINDOWS_DEFAULT_LOCATION, true
	}
	if CheckFileExists(I2PD_LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2pd router detected")
		return I2PD_LINUX_SYSTEM_LOCATION[0], true
	}
	if CheckFileExists(I2PD_LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2pd router detected")
		return I2PD_LINUX_SYSTEM_LOCATION[1], true
	}
	return "", false
}

// checkJavaI2PWindowsLocations checks for Java I2P router installations on Windows
func checkJavaI2PWindowsLocations() (string, bool) {
	if CheckFileExists(WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return WINDOWS_DEFAULT_LOCATION, true
	}
	if CheckFileExists(WINDOWS_ZERO_NSIS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return WINDOWS_ZERO_NSIS_DEFAULT_LOCATION, true
	}
	return "", false
}

// checkJavaI2PLinuxLocations checks for Java I2P router installations on Linux
func checkJavaI2PLinuxLocations() (string, bool) {
	if CheckFileExists(LINUX_SYSTEM_LOCATION[0]) {
		log.Println("Linux i2p router detected")
		return LINUX_SYSTEM_LOCATION[0], true
	}
	if CheckFileExists(LINUX_SYSTEM_LOCATION[1]) {
		log.Println("Linux i2p router detected")
		return LINUX_SYSTEM_LOCATION[1], true
	}
	return "", false
}

// checkUserDirectoryLocations checks for I2P router installations in user directories
func checkUserDirectoryLocations() (string, bool) {
	if CheckFileExists(HOME_DIRECTORY_LOCATION) {
		log.Println("Linux i2p router detected")
		return HOME_DIRECTORY_LOCATION, true
	}
	if CheckFileExists(OSX_DEFAULT_LOCATION) {
		log.Println("OSX i2p router detected")
		return OSX_DEFAULT_LOCATION, true
	}
	return "", false
}

// FindI2PIsInstalledDefaultLocation looks in various locations for the
// presence of an I2P router, reporting back the location
func FindI2PIsInstalledDefaultLocation() (string, error) {
	if location, found := checkI2PDRouterLocations(); found {
		return location, nil
	}
	if location, found := checkJavaI2PWindowsLocations(); found {
		return location, nil
	}
	if location, found := checkJavaI2PLinuxLocations(); found {
		return location, nil
	}
	if location, found := checkUserDirectoryLocations(); found {
		return location, nil
	}
	return "", fmt.Errorf("i2p router not found")
}
