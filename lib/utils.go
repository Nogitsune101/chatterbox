package lib

import (
	"os"
	"strings"
)

// IsAppBinary determines if the application is being launched from go run or as a binary
// IMPORTANT - If golang ever changes its default go run temp app location, this will not be viable anymore
func IsAppBinary() bool {
	if strings.Contains(os.Args[0], "b001/exe/main") {
		return false
	}
	return true
}
