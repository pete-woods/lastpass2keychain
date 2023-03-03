package keychain

import (
	"strings"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

func Test_readLastpassCSV(t *testing.T) {
	tests := []struct {
		name        string
		in          string
		wantEntries []entry
	}{
		{
			name: "Basic",
			in: `url,username,password,totp,extra,name,grouping,fav
http://www.amazon.co.uk,amazon.co.uk@email.com,the-password,,,amazon.co.uk,Shopping,0
https://google.com/,email@email.com,the-password,,,google.com,Email,0
https://www.facebook.com,facebook@email.com,the-password,,,facebook.com,Social,0
https://twitter.com,twitter.com@email.com,the-password,,,twitter.com,Social,0
https://auth.api.sonyentertainmentnetwork.com,playstation.com@email.com,the-password,,"ABCD
EFGH
IJKL
MNOP
QRST
UVWZ",playstation.com,Games,0`,
			wantEntries: []entry{
				{
					URL:      "http://www.amazon.co.uk",
					Username: "amazon.co.uk@email.com",
					Password: "the-password",
					Name:     "amazon.co.uk",
					Grouping: "Shopping",
					Fav:      "0",
				},
				{
					URL:      "https://google.com/",
					Username: "email@email.com",
					Password: "the-password",
					Name:     "google.com",
					Grouping: "Email",
					Fav:      "0",
				},
				{
					URL:      "https://www.facebook.com",
					Username: "facebook@email.com",
					Password: "the-password",
					Name:     "facebook.com",
					Grouping: "Social",
					Fav:      "0",
				},
				{
					URL:      "https://twitter.com",
					Username: "twitter.com@email.com",
					Password: "the-password",
					Name:     "twitter.com",
					Grouping: "Social",
					Fav:      "0",
				},
				{
					URL:      "https://auth.api.sonyentertainmentnetwork.com",
					Username: "playstation.com@email.com",
					Password: "the-password",
					Extra: `ABCD
EFGH
IJKL
MNOP
QRST
UVWZ`,
					Name:     "playstation.com",
					Grouping: "Games",
					Fav:      "0",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var entries []entry
			err := readLastpassCSV(strings.NewReader(tt.in), func(e entry) error {
				entries = append(entries, e)
				return nil
			})
			assert.Assert(t, err)
			assert.Check(t, cmp.DeepEqual(entries, tt.wantEntries))
		})
	}
}
