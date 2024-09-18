import XCTest
import SwiftTreeSitter
import TreeSitterSage

final class TreeSitterSageTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_sage())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Sage grammar")
    }
}
