//
//  Day02.swift
//
//
//  Created by Bui Duong on 03/12/2023.
//

import Foundation

struct Day02: AdventDay {
    var data: String

    var maps: [[Character]] {
        let lines = data.split(separator: "\n")
        var result: [[Character]] = []

        for line in lines {
            result.append(Array(line))
        }

        return result
    }

    func extractNumber(_ line: String, pos: Int) -> (Int, Int) {
        let arrLine = Array(line)
        var leftBound = pos
        var rightBound = pos
        while leftBound > 0 {
            if arrLine[leftBound - 1].isNumber {
                leftBound -= 1
            } else {
                break
            }
        }
        while rightBound < line.count - 1 {
            if arrLine[rightBound + 1].isNumber {
                rightBound += 1
            } else {
                break
            }
        }
        return (leftBound, rightBound)
    }

    func searchNumberAround(_ xAxis: Int, _ yAxis: Int, _ maps: inout [[Character]], line: String) -> [Int] {
        var result: [Int] = []
        let xBoundMin =
            if xAxis > 0 { xAxis - 1 } else { 0 }
        let xBoundMax =
            if xAxis < line.count - 1 { xAxis + 1 } else { line.count }
        let yBoundMin =
            if yAxis > 0 { yAxis - 1 } else { 0 }
        let yBoundMax =
            if yAxis < maps.count - 1 { yAxis + 1 } else { maps.count }

        for x in xBoundMin...xBoundMax {
            for y in yBoundMin...yBoundMax {
                if maps[y][x].isNumber {
                    let line = String(maps[y])
                    let (left, right) = extractNumber(line, pos: x)
                    let num = Int(line[line.index(line.startIndex, offsetBy: left)...line.index(line.startIndex, offsetBy: right)])!
                    for index in left...right {
                        maps[y][index] = "."
                    }
                    result.append(num)
                }
            }
        }

        return result
    }

    func processNumberAroundSymbol(
        _ xAxis: Int,
        _ yAxis: Int,
        _ mutatedMap: inout [[Character]],
        _ line: String,
        _ converter: (Character, [Int]) -> Int) -> Int {
        converter(
            maps[yAxis][xAxis],
            searchNumberAround(
                xAxis, yAxis, &mutatedMap,
                line: String(line)
            )
        )

    }

    func part1() -> Any {
        var mutatedMap = maps
        var nums: [Int] = []
        for (yAxis, line) in mutatedMap.enumerated() {
            for (xAxis, char) in line.enumerated() {
                if !(char.isNumber || char == ".") {
                    nums.append(
                        processNumberAroundSymbol(
                            xAxis, yAxis,
                                &mutatedMap, String(line),
                            {
                                $1.reduce(0) { $0 + $1 }
                            }
                        )
                    )
                }
            }
        }
        return nums.reduce(0) { $0 + $1 }
    }

    func part2() -> Any {
        var mutatedMap = maps
        var nums: [Int] = []
        for (yAxis, line) in mutatedMap.enumerated() {
            for (xAxis, char) in line.enumerated() {
                if !(char.isNumber || char == ".") {
                    nums.append(
                        processNumberAroundSymbol(
                            xAxis, yAxis,
                                &mutatedMap, String(line),
                            {
                                switch $0 {
                                case let sym where sym == "*" && $1.count == 2:
                                    $1.reduce(1) { $0 * $1 }
                                default:
                                    0
                                }
                            }
                        )
                    )
                }
            }
        }
        return nums.reduce(0) { $0 + $1 }
    }
}
