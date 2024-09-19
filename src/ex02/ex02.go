package ex02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func compare(s1, s2 *bufio.Scanner) ([]string, []string) {
	paths1, paths2 := make(map[string]bool, 100), make(map[string]bool, 100)
	removed, added := make([]string, 0, 100), make([]string, 0, 100)

	for s1.Scan() {
		line := s1.Text()
		paths1[line] = true
	}

	for s2.Scan() {
		line := s2.Text()
		paths2[line] = true
	}

	for path2 := range paths2 {
		_, ok := paths1[path2]
		if !ok {
			added = append(added, path2)
		}
	}

	for path1 := range paths1 {
		_, ok := paths2[path1]
		if !ok {
			removed = append(removed, path1)
		}
	}
	return removed, added
}

func CompareFS(f1Path *string, f2Path *string) {
	f1, err := os.Open(*f1Path)
	if err != nil {
		log.Fatal("Incorrect old file path, try again\n")
	}
	f2, err := os.Open(*f2Path)
	if err != nil {
		log.Fatal("Incorrect new file path, try again")
	}
	defer f1.Close()
	defer f2.Close()

	if !strings.HasSuffix(*f1Path, ".txt") ||
		!strings.HasSuffix(*f2Path, ".txt") {
		log.Fatal("Unsupported file type, only can work with .txt dumps\n")
	}

	r1, r2 := bufio.NewScanner(f1), bufio.NewScanner(f2)
	removed, added := compare(r1, r2)

	for i := 0; i < len(added); i++ {
		fmt.Printf("ADDED %s\n", added[i])
	}
	for i := 0; i < len(removed); i++ {
		fmt.Printf("REMOVED %s\n", removed[i])
	}
}
