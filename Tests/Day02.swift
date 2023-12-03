//
//  File.swift
//  
//
//  Created by Bui Duong on 03/12/2023.
//

import XCTest

@testable import AdventOfCode

final class Day02Tests: XCTestCase {
    let testData = """
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
"""
    
    func testPart1() throws {
        let challenge = Day02(data: testData)
        XCTAssertEqual(String(describing: challenge.part1()), "4361")
    }
    
    func testPart2() throws {
        let challenge = Day02(data: testData)
        XCTAssertEqual(String(describing: challenge.part2()), "467835")
    }
}
