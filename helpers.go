package goLogz

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func whatIoChecker(param string) (io.Writer, error) {

	if param == "" {
		return os.Stdout, nil
	}

	switch param {
	case "STDOUT":
		return os.Stdout, nil
	case "STDERR":
		return os.Stderr, nil
	case "DISCARD":
		return ioutil.Discard, nil
	default:
		tmp := strings.Split(param, "/")
		folder := strings.Join(tmp[:len(tmp)-1], "/")
		if _, err := os.Stat(folder); os.IsNotExist(err) {
			err = os.Mkdir(folder, 0666)
			if err != nil {
				return nil, err
			}
		}
		f, err := os.OpenFile(param, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		return f, nil
	}
}
