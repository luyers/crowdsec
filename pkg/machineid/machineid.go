package machineid

import (
	"io/ioutil"
	"strings"
)

const (
	// dbusPath is the default path for dbus machine id.
	dbusPath = "/var/lib/dbus/machine-id"
	// dbusPathEtc is the default path for dbus machine id located in /etc.
	// Some systems (like Fedora 20) only know this path.
	// Sometimes it's the other way round.
	dbusPathEtc = "/etc/machine-id"
)

// idea of code is stolen from https://github.com/denisbrodbeck/machineid/
// but here we are on Debian GNU/Linux
func ID() (string, error) {
	id, err := ioutil.ReadFile(dbusPath)
	if err != nil {
		// try fallback path
		id, err = ioutil.ReadFile(dbusPathEtc)
	}
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(id)), nil
}
