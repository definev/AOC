//
//  File.swift
//
//
//  Created by Bui Duong on 10/12/2023.
//

import Foundation

struct Race {
    var time: Int
    var distance: Int
}

struct Day05: AdventDay {
    var data: String

    var entites: [Race] {
        var races: [Race] = []
        let lines = data.split(separator: "\n")
            .map {
            String($0
                .split(separator: ":")[1]
                .trimmingCharacters(in: .whitespacesAndNewlines))
                .replacingOccurrences(of: "\\s+", with: " ", options: .regularExpression)
                .split(separator: " ")
                .map { Int($0)! }
        }

        for (index, time) in lines[0].enumerated() {
            races.append(Race(time: time, distance: lines[1][index]))
        }

        return races
    }

    var part2Race: Race {
        let lines = data.split(separator: "\n")
            .map {
            Int(String($0
                .split(separator: ":")[1]
                .trimmingCharacters(in: .whitespacesAndNewlines))
                .replacingOccurrences(of: "\\s", with: "", options: .regularExpression))!
        }
        return Race(time: lines[0], distance: lines[1])
    }

    func calculateWinTurn(_ race: Race) -> Int {
        var winTurn = 0
        for index in 1..<race.time {
            if index * (race.time - index) > race.distance {
                winTurn += 1
            }
        }
        return winTurn
    }

    func part1() -> Any {
        var result = 1
        for race in entites {
            let winTurn = calculateWinTurn(race)
            result *= winTurn
        }
        return result
    }

    func part2() -> Any {
        return calculateWinTurn(part2Race)
    }
}
