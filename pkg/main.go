package pkg

import (
	"os"

	"github.com/ericchiang/css"
)

func ReadFile(filename string) (*[]byte, error) {
	// Attempt to open the file
	f, err := os.ReadFile(filename)
	if err != nil {
		// Check specifically if the error indicates the file does not exist
		return nil, err
	}
	return &f, nil
}

func Process(cssData *string) *css.Selector {
	sel, err := css.Parse(*cssData)
	if err != nil {
		panic(err)
	}
	return sel
}
