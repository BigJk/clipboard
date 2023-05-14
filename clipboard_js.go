// +build js

package clipboard

import (
  "errors"
)

func readAll() (string, error) {
  return "", errors.New("unavailable in js") 
}

func writeAll(text string) error {
	return errors.New("unavailable in js") 
}
