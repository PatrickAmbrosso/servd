//go:build windows

package env

import (
	"golang.org/x/sys/windows"
)

// Check if the process is running as an administrator
func IsAdmin() bool {
	var sid *windows.SID
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid,
	)
	if err != nil {
		return false
	}

	token := windows.Token(0)
	isAdmin, err := token.IsMember(sid)
	return err == nil && isAdmin
}
