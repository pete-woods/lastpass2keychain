package keychain

import (
	"testing"

	"github.com/keybase/go-keychain"
	"gotest.tools/v3/assert"
)

func Test_ensureKeychainItem(t *testing.T) {
	t.Skip()
	err := ensureInternetPassword(entry{
		Name:     "Example dot com",
		URL:      "https://example.com:123/the/path",
		Username: "the-username",
		Password: "the-password",
		Extra:    "ABCDEFG",

		TOTP:     "",
		Grouping: "",
		Fav:      "",
	}, false)
	assert.Assert(t, err)
}

func TestListAllInternalPasswords(t *testing.T) {
	t.Skip()
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassInternetPassword)
	query.SetMatchLimit(keychain.MatchLimitAll)
	query.SetReturnAttributes(true)
	results, err := keychain.QueryItem(query)
	assert.Assert(t, err)
	for _, r := range results {
		t.Logf("%#v", r)
	}
}
