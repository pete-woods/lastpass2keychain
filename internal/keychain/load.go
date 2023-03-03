package keychain

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func Load(filename string, dryRun bool) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer closer(f, &err)

	r := bufio.NewReader(f)

	return readLastpassCSV(r, func(e entry) error {
		if e.URL == "http://sn" {
			fmt.Printf("SN %q %q %q\n", e.Grouping, e.Name, e.Extra)
			return nil
		}

		fmt.Printf("IP %q %q %q %q", e.Grouping, e.Name, e.URL, e.Username)
		err := ensureInternetPassword(e, dryRun)
		switch {
		case errors.Is(err, errDuplicate):
			fmt.Println(" | DUP")
			return nil
		case err != nil:
			fmt.Println(" | ERR")
			return err
		default:
			fmt.Println(" | ADD")
			return nil
		}
	})
}

func closer(c io.Closer, in *error) {
	cerr := c.Close()
	if *in == nil {
		*in = cerr
	}
}
