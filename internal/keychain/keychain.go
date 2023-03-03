package keychain

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/keybase/go-keychain"
)

var errDuplicate = errors.New("duplicate item")

func ensureInternetPassword(e entry, dryRun bool) error {
	u, err := url.Parse(e.URL)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %q", e.URL)
	}

	portStr := u.Port()
	port := 0
	if portStr != "" {
		port, err = strconv.Atoi(portStr)
		if err != nil {
			return fmt.Errorf("failed to parse port: %q", portStr)
		}
	}

	protocol := ""
	switch u.Scheme {
	case "https":
		protocol = "htps"
	case "http":
		protocol = "http"
	default:
		return fmt.Errorf("unknown scheme: %q", u.Scheme)
	}

	if dryRun {
		return nil
	}

	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassInternetPassword)

	// Internet password-specific
	item.SetProtocol(protocol)
	item.SetServer(u.Hostname())
	item.SetPort(int32(port))
	item.SetPath(u.Path)

	item.SetAccount(e.Username)
	item.SetLabel(e.Name)
	item.SetData([]byte(e.Password))
	item.SetComment(e.Extra)

	//item.SetAccessGroup("A123456789.group.com.mycorp")
	//item.SetSynchronizable(keychain.SynchronizableNo)
	//item.SetAccessible(keychain.AccessibleWhenUnlocked)

	err = keychain.AddItem(item)
	if err == keychain.ErrorDuplicateItem {
		return fmt.Errorf("%w: %s", errDuplicate, err)
	}
	return err
}
