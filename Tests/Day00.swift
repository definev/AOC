import XCTest

@testable import AdventOfCode

// Make a copy of this file for every day to ensure the provided smoke tests
// pass.
final class Day00Tests: XCTestCase {
  // Smoke test data provided in the challenge question
  let testData = """
    two1nine
    eightwothree
    abcone2threexyz
    xtwone3four
    4nineeightseven2
    zoneight234
    7pqrstsixteen
    """

  func testPart1() throws {
    let challenge = Day00(data: testData)
    XCTAssertEqual(String(describing: challenge.part1()), "142")
  }

  func testPart2() throws {
    let challenge = Day00(data: testData)
    XCTAssertEqual(String(describing: challenge.part2()), "281")
  }
}
