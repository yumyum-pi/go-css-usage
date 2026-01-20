package css

import (
	"os"

	css2 "github.com/aymerick/douceur/css"
	"github.com/aymerick/douceur/parser"
)

func ReadFile(filepath string) (*css2.Stylesheet, error) {
	// parse css
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return parser.Parse(string(b))
}
