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
)

// homedir gets the home directory for a user given their username
// or an error if the lookup was not successful
func Homedir(username string) (string, error) {
	cName := C.CString(username)
	cHome := C.homedir(cName)
	if cHome == nil {
		return "", errors.New("Failed to look up home directory")
	}
	return C.GoString(cHome), nil
}

// login logs in the username with password and returns the pid of the
// login process or an error if login failed
func Login(username string, password string, exec string) (int, error) {
	cUser := C.CString(username)
	cPass := C.CString(password)
	cExec := C.CString("exec " + exec)

	var child C.pid_t
	ok := bool(C.login(cUser, cPass, cExec, &child))
	if !ok {
		return 0, errors.New("Login failed")
	}
	return int(child), nil
}

// logout requests the user log out or returns an error
func Logout() error {
	if !bool(C.logout()) {
		return errors.New("Failed to log out user")
	}

	return nil
}
