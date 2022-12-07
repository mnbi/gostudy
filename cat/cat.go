package cat

import (
	"errors"
)

func Cat(pathname string) error {
	if pathname == "" {
		return errors.New("empty pathname")
	}

	return nil
}
