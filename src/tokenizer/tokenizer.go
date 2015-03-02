package tokenizer

import (
	"go/token"
	"go/scanner"
)

type TokenInfo struct {
	Position int
	Type string
	Literal string
}

func Tokenize(input string) []TokenInfo {
	var result []TokenInfo
	
	src := []byte(input)
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(src))
	
	var s scanner.Scanner
	s.Init(file, src, nil, scanner.ScanComments)
	
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		result = append(result, TokenInfo{fset.Position(pos).Offset, tok.String(), lit})
	}
	
	return result
}

