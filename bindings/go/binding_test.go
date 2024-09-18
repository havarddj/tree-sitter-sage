package tree_sitter_sage_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_sage "github.com/tree-sitter/tree-sitter-sage/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_sage.Language())
	if language == nil {
		t.Errorf("Error loading Sage grammar")
	}
}
