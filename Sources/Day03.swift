//
//  Day03.swift
//
//
//  Created by Bui Duong on 04/12/2023.
//

import Foundation

extension RandomAccessCollection where Element: Comparable {
    func binarySearch(_ element: Element) -> (index: Index, found: Bool) {
        let index = partitioningIndex(where: { $0 >= element })
        let found = index != endIndex && self[index] == element
        return (index, found)
    }
}

struct LotteDay {
    var winningNumbers: [Int]
    var numbers: [Int]
}

struct Day03: AdventDay {
    var data: String

    var entities: [LotteDay] {
        let rawLines = data.split(separator: "\n")
        var result: [LotteDay] = []
        for line in rawLines {
            var day: LotteDay = LotteDay(winningNumbers: [], numbers: [])
            let segments = line.split(separator: ": ")[1].split(separator: " | ")
            let winningNumbersRaw = segments[0].trimmingCharacters(in: .whitespacesAndNewlines)
            let numbersRaw = segments[1].trimmingCharacters(in: .whitespacesAndNewlines)
            for part in winningNumbersRaw.split(separator: " ") {
                if part.isEmpty {
                    continue
                }
                day.winningNumbers.append(Int(part.trimmingCharacters(in: .whitespacesAndNewlines))!)
            }

            for part in numbersRaw.split(separator: " ") {
                if part.isEmpty {
                    continue
                }
                day.numbers.append(Int(part.trimmingCharacters(in: .whitespacesAndNewlines))!)
            }
            day.winningNumbers.sort { $0 < $1 }
            day.numbers.sort { $0 < $1 }
            result.append(day)
        }
        return result
    }

    func calculatePoint(_ day: LotteDay) -> Int {
        var result = 0
        for number in day.numbers {
            let (_, exists) = day.winningNumbers.binarySearch(number)
            if exists {
                if result == 0 { result = 1 }
                else { result *= 2 }
            }
        }
        return result
    }

    func findNumberOfWinCard(_ day: LotteDay) -> Int {
        var result = 0
        for number in day.numbers {
            let (_, exists) = day.winningNumbers.binarySearch(number)
            if exists {
                result += 1
            }
        }
        return result
    }

    func part1() -> Any {
        let days = entities
        var total = 0
        for day in days {
            total += calculatePoint(day)
        }
        return total
    }

    func part2() -> Any {
        let days = entities
        var totalWinDays: [Int] = Array(repeating: 1, count: days.count)
        for (index, day) in days.enumerated() {
            let winDay = findNumberOfWinCard(day)
            let repeated = totalWinDays[index]
            if winDay <= 0 { continue }
            for dayNum in 1...winDay {
                let dayIndex = dayNum + index
                if dayIndex >= totalWinDays.count { break }
                totalWinDays[dayIndex] = totalWinDays[dayIndex] + 1 * repeated
            }
        }

        return totalWinDays.reduce(0) { $0 + $1 }
    }
}
