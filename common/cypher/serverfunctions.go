package cypher

import (
	"os"
	"os/user"
	"syscall"
)

func ChangeUser(username string) error {
	u, err := user.Lookup(username)
	if err != nil {
		return err
	}
	uid := u.Uid
	gid := u.Gid
	uidInt, _ := strconv.Atoi(uid)
	gidInt, _ := strconv.Atoi(gid)
	if err := syscall.Setgid(gidInt); err != nil {
		return err
	}
	if err := syscall.Setuid(uidInt); err != nil {
		return err
	}
	return nil
}