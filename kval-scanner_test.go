package kvalscanner

import (
	"strings"
	"testing"
)

// Scanner test results struct
type scannerTest struct {
	scanvalue string
	expected  Token
}

// Scan basic Tokens
var scannerTests = []scannerTest{
	{"", EOF},
	{"*", LITERAL},
	{" ", WS},
	{"\n", WS},
	{"\r", WS},
	{"_", USCORE},
	{"\x1F", ILLEGAL},
}

// Test that basic token parsing works...
func TestScan(t *testing.T) {
	for _, expected := range scannerTests {
		s := NewScanner(strings.NewReader(expected.scanvalue))
		var tok Token
		var lit string
		for tok != EOF {
			tok, lit = s.Scan()
			if (tok != EOF && expected.expected != EOF) && tok != expected.expected {
				//EOF returned each scan, so ignore if it's not what we're testing...
				t.Errorf("FAIL: Got %d '%s' when %d '%s' was expected.", tok, lit, expected.expected, expected.scanvalue)
			}
		}
	}
}

// Test a range of literals that should be allowed by the library
func TestValidLiterals(t *testing.T) {
	var lits = []string{"āēīōūĀĒĪŌŪ", ">>>", "/)(*&^%$#@!>:!@#", "abc123", "子：？", "200", "\n\r", "string!\n\r"}
	for _, s := range lits {
		s := NewScanner(strings.NewReader(s))
		var tok Token
		var lit string
		for tok != EOF {
			tok, lit = s.Scan()
			if tok == ILLEGAL {
				t.Errorf("FAIL: Illegal token '%d' '%s' when expecting valid LITERAL value.", tok, lit)
			}
		}
	}
}
