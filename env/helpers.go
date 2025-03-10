//go:build linux || darwin

package env

func IsAdmin() bool {
	// if os.Geteuid() == 0 {
	// 	return true
	// }

	// // Check if the user belongs to the `sudo` or `wheel` group
	// cmd := exec.Command("groups")
	// cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true} // Avoid process group issues
	// output, err := cmd.Output()
	// if err != nil {
	// 	return false
	// }

	// groups := strings.Fields(string(output))
	// for _, group := range groups {
	// 	if group == "sudo" || group == "wheel" {
	// 		return true
	// 	}
	// }

	return false
}
