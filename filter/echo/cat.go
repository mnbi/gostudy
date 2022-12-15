package cat

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Cat(pathname string) error {
	if pathname == "" {
		return errors.New("empty pathname")
	}

	f, err := os.Open(pathname)
	defer f.Close()
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}
