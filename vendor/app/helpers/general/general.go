package general

import (
	"bufio"
	"io"
	"os"
)

func Scanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	var str string
	for scanner.Scan() {
		str = scanner.Text()
		break
	}
	return str
}

func Copy(dst, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}
