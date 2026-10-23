package pterm

import "os/exec"

func LookPath(cmds []string, fallback string) string {
	for _, cmd := range cmds {
		if str, err := exec.LookPath(cmd); err == nil && str != "" {
			return cmd
		}
	}

	return fallback
}
