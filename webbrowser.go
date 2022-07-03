// Package webbrowser provides functions to open a URL in a browser window.
package webbrowser

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

// OpenBrowser opens a browser window to the specified URL.
func Open(url string) error {
	var err error
	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default:
		err = errors.New("unsupported platform")
	}

	if err != nil {
		return fmt.Errorf("unable too open url %w", err)
	}
	return nil
}
