package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "packagepaths <importpath>\n")
		os.Exit(2)
	}
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
	}
	importPath := flag.Arg(0)

	cfg := packages.Config{
		Mode: packages.LoadAllSyntax,
		ParseFile: func(fset *token.FileSet, filename string, src []byte) (*ast.File, error) {
			return parser.ParseFile(fset, filename, src, parser.ParseComments)
		},
	}
	pkgs, err := packages.Load(&cfg, importPath)
	if err != nil {
		log.Fatalf("cannot load package: %v", err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("packages.Load returned %d packages, not 1", len(pkgs))
	}
	pkg := pkgs[0]
	files := make(map[string]bool)
	for _, f := range pkg.Syntax {
		if tokFile := pkg.Fset.File(f.Pos()); tokFile != nil {
			files[tokFile.Name()] = true
		}
	}
	var fileStrs []string
	for f := range files {
		fileStrs = append(fileStrs, f)
	}
	sort.Strings(fileStrs)
	fmt.Println(strings.Join(fileStrs, "\n"))
}
