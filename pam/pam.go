package pam

/*
#cgo LDFLAGS: -lpam

#include <stdbool.h>
#include <stdlib.h>
#include <unistd.h>

bool login(const char *username, const char *password, const char *exec, pid_t *child_pid);
bool logout(void);
char *homedir(const char *username);
*/
import "C"
import (
	"errors"
	"unsafe"
)

// homedir gets the home directory for a user given their username
// or an error if the lookup was not successful
func Homedir(username string) (string, error) {
	cName := C.CString(username)
	defer C.free(unsafe.Pointer(cName))
	cHome := C.homedir(cName)
	if cHome == nil {
		return "", errors.New("failed to look up home directory")
	}
	// cHome points to pw->pw_dir from getpwnam's static buffer; do not free.
	return C.GoString(cHome), nil
}

// login logs in the username with password and returns the pid of the
// login process or an error if login failed
func Login(username string, password string, exec string) (int, error) {
	cUser := C.CString(username)
	defer C.free(unsafe.Pointer(cUser))
	cPass := C.CString(password)
	defer C.free(unsafe.Pointer(cPass))
	cExec := C.CString("exec " + exec)
	defer C.free(unsafe.Pointer(cExec))

	var child C.pid_t
	ok := bool(C.login(cUser, cPass, cExec, &child))
	if !ok {
		return 0, errors.New("login failed")
	}
	return int(child), nil
}

// logout requests the user log out or returns an error
func Logout() error {
	if !bool(C.logout()) {
		return errors.New("failed to log out user")
	}
	return nil
}
