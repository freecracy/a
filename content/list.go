package content

import (
	"os"
	"strings"
)

func GetAllFile(name string) []string {
	s := []string{}
	file, _ := os.Open(name)
	files, _ := file.Readdir(0)
	for _, f := range files {
		s = append(s, strings.TrimSuffix(strings.ToLower(f.Name()), ".md"))
	}
	return s
}
