//
//  File.swift
//  
//
//  Created by Bui Duong on 10/12/2023.
//

import XCTest

@testable import AdventOfCode

final class Day05Tests : XCTestCase {
    var testData: String = """
Time:      7  15   30
Distance:  9  40  200
"""
    
    func testPart1() throws {
        let challenge = Day05(data: testData)
        XCTAssertEqual(String(describing: challenge.part1()), "288")
    }
    
    func testPart2() throws {
        let challenge = Day05(data: testData)
        XCTAssertEqual(String(describing: challenge.part2()), "71503")
    }
}
