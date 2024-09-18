from unittest import TestCase

import tree_sitter, tree_sitter_sage


class TestLanguage(TestCase):
    def test_can_load_grammar(self):
        try:
            tree_sitter.Language(tree_sitter_sage.language())
        except Exception:
            self.fail("Error loading Sage grammar")
