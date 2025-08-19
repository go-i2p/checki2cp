package util

import (
	"log"
	"os"
	"os/user"
)

// UserFind makes sure that we never mis-identify the user account because of
// sudo
func UserFind() string {
	if os.Geteuid() == 0 {
		str := os.Getenv("SUDO_USER")
		return str
	}
	if un, err := user.Current(); err == nil {
		return un.Name
	}
	return ""
}

// CheckI2PUserName looks in various locations for the
// presence of an I2P router and guesses the username it
// should run under. On Windows it returns the EXE name.
func CheckI2PUserName() (string, error) {
	if username := checkWindowsRouters(); username != "" {
		return username, nil
	}
	if username := checkLinuxSystemRouters(); username != "" {
		return username, nil
	}
	if username := checkUserDirectoryRouters(); username != "" {
		return username, nil
	}
	return "", nil
}

// checkWindowsRouters detects Windows-based I2P router installations.
func checkWindowsRouters() string {
	if CheckFileExists(I2PD_WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2pd router detected")
		return "i2pd"
	}
	if CheckFileExists(WINDOWS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return "i2psvc"
	}
	if CheckFileExists(WINDOWS_ZERO_NSIS_DEFAULT_LOCATION) {
		log.Println("Windows i2p router detected")
		return "i2p-zero"
	}
	return ""
}

// checkLinuxSystemRouters detects system-level Linux I2P router installations.
func checkLinuxSystemRouters() string {
	for _, location := range I2PD_LINUX_SYSTEM_LOCATION {
		if CheckFileExists(location) {
			log.Println("Linux i2pd router detected")
			return "i2pd"
		}
	}
	for _, location := range LINUX_SYSTEM_LOCATION {
		if CheckFileExists(location) {
			log.Println("Linux i2p router detected")
			return "i2psvc"
		}
	}
	return ""
}

// checkUserDirectoryRouters detects user-directory I2P router installations.
func checkUserDirectoryRouters() string {
	if CheckFileExists(HOME_DIRECTORY_LOCATION) {
		log.Println("Linux i2p router detected")
		return UserFind()
	}
	if CheckFileExists(OSX_DEFAULT_LOCATION) {
		log.Println("OSX i2p router detected")
		return UserFind()
	}
	return ""
}
