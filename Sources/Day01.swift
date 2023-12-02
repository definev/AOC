//
//  Day01.swift
//
//
//  Created by Bui Duong on 02/12/2023.
//

import Algorithms

struct Bag {
    var red: Int = 0
    var green: Int = 0
    var blue: Int = 0

    static func parseLine(_ raw: String) -> Bag {
        var bag = Bag()
        let colors = raw.split(separator: ", ").map { $0.split(separator: " ") }
        for index in 0..<colors.count {
            let color = colors[index]
            switch color[1] {
            case "red":
                bag.red = Int(color[0])!
            case "green":
                bag.green = Int(color[0])!
            case "blue":
                bag.blue = Int(color[0])!
            default:
                print("something went wrong!")
            }
        }

        return bag
    }
}

struct Game {
    let id: Int
    let sets: [Bag]

    static func parseLine(_ raw: String) -> Game {
        let splitRaw = raw.split(separator: ":")
        let gamePart = splitRaw[0].trimmingCharacters(in: .whitespacesAndNewlines)
        let setsPart = splitRaw[1].trimmingCharacters(in: .whitespacesAndNewlines)
        let id = Int(gamePart.split(separator: " ")[1].trimmingCharacters(in: .whitespacesAndNewlines))!

        let sets = setsPart.split(separator: ";").map { Bag.parseLine(String($0)) }

        return Game(id: id, sets: sets)
    }
}

struct Day01: AdventDay {
    // Save your data in a corresponding text file in the `Data` directory.
    var data: String

    var entities: [Game] {
        data.split(separator: "\n").map { Game.parseLine(String($0)) }
    }

    func isPossibleGame(_ game: Game) -> Bool {
        for set in game.sets {
            if set.red > 12 || set.green > 13 || set.blue > 14 {
                return false
            }
        }
        return true
    }

    // Replace this with your solution for the first part of the day's challenge.
    func part1() -> Any {
        let total = try! entities.reduce<Int>(0) { total, game in if isPossibleGame(game) { return total + game.id } else { return total } }
        return total
    }

    func minimumSetCube(_ game: Game) -> Int {
        var minimumBag = Bag()
        for set in game.sets {
            if minimumBag.red < set.red {
                minimumBag.red = set.red
            }
            if minimumBag.green < set.green {
                minimumBag.green = set.green
            }
            if minimumBag.blue < set.blue {
                minimumBag.blue = set.blue
            }
        }
        return minimumBag.red * minimumBag.green * minimumBag.blue
    }

    // Replace this with your solution for the second part of the day's challenge.
    func part2() -> Any {
        let total = try! entities.reduce<Int>(0) { total, game in total + minimumSetCube(game) }
        return total
    }
}
