//
//  Day04.swift
//
//
//  Created by Bui Duong on 04/12/2023.
//

struct FarmMap {
    var from: String
    var to: String
    var ranges: [(Int, Int, Int)]

    static func parse(_ raw: String) -> FarmMap {
        var sequences = raw.split(separator: "\n")
        let title = sequences.first?.split(separator: " ")[0].split(separator: "-to-")

        sequences.removeFirst()

        var ranges: [(Int, Int, Int)] = []

        for sequence in sequences {
            let list = sequence.split(separator: " ")
            ranges.append(
                (Int(list[0])!, Int(list[1])!, Int(list[2])!))
        }

        return FarmMap(
            from: String(title?[0] ?? ""),
            to: String(title?[1] ?? ""),
            ranges: ranges
        )
    }
}

struct Farm {
    var destinations: [Int]
    var maps: [FarmMap]

}

struct Day04: AdventDay {
    var data: String

    var entity: Farm {
        var lines = data.split(separator: "\n\n")

        let destinations = (lines.first?
            .split(separator: ": ")[1]
            .split(separator: " ").map { Int($0)! })!
        lines.removeFirst()
        var maps: [FarmMap] = []

        for mapLine in lines {
            maps.append(FarmMap.parse(String(mapLine)))
        }

        return Farm(destinations: destinations, maps: maps)
    }

    func describleRange(_ range: (Int, Int, Int)) -> String {
        var result = ""
        result += "s_start: \(range.1) s_end: \(range.1 + range.2)"
        result += " | d_start: \(range.0) d_end: \(range.0 + range.2)"
        return result
    }

    func convertThrough(_ destination: Int, _ maps: [FarmMap]) -> Int {
        var location = destination
        var road = "seed: \(location)"

        for map in maps {
            for range in map.ranges {
                let lowerBound = range.1
                let upperBound = range.1 + range.2
                if location >= lowerBound && location <= upperBound {
                    let translationSteps = location - range.1
                    location = range.0 + translationSteps
                    break
                }
            }
            road += "\n -> \(map.to): \(location)"
        }

        return location
    }

    func part1() -> Any {
        let data = entity
        var minValue = Int.max

        for destination in data.destinations {
            minValue = min(
                minValue,
                convertThrough(destination, data.maps)
            )
        }

        return minValue
    }

    func convertThroughAndSaveMaxIndex(_ destination: Int, _ maps: [FarmMap]) -> (Int, Int) {
        var location = destination
        var road = "seed: \(location)"
        var maximumTranslationSteps = Int.max

        for map in maps {
            var rangeIncluded: (Int, Int, Int)? = nil
            var nearestLowerBound: Int? = nil

            for range in map.ranges {
                let lowerBound = range.1
                let upperBound = range.1 + range.2

                if lowerBound > location {
                    let distance = lowerBound - location
                    nearestLowerBound = min(
                        nearestLowerBound ?? distance,
                        distance
                    )
                }

                if location >= lowerBound && location <= upperBound {
                    nearestLowerBound = nil
                    rangeIncluded = range
                    let translationSteps = location - range.1
                    let remainStep = (range.1 + range.2) - location
                    location = range.0 + translationSteps
                    maximumTranslationSteps = min(maximumTranslationSteps, remainStep)
                    break
                }
            }

            if let rangeIncluded = rangeIncluded {
                road += "\n -> \(map.to): \(location) \(describleRange(rangeIncluded)) - \(rangeIncluded.2) | \(maximumTranslationSteps)"
                continue
            }

            if let nearestLowerBound = nearestLowerBound {
                maximumTranslationSteps = min(
                    maximumTranslationSteps,
                    nearestLowerBound
                )
                road += "\n -> nearest \(map.to): \(location) | \(nearestLowerBound)"
                continue
            }
        }
        
        if maximumTranslationSteps == Int.max {
            maximumTranslationSteps = 1
        }

        road += "\n--- Max steps: \(maximumTranslationSteps) | Final location: \(location)"

        return (max(maximumTranslationSteps, 1), location)
    }

    func part2() -> Any {
        let data = entity
        var minValue = Int.max

        for i in stride(from: 0, to: data.destinations.count, by: 2) {
            let chunk = Array(data.destinations[i..<min(i + 2, data.destinations.count)])

            var destinationIndex = 0

            while destinationIndex < chunk[1] {
                let (step, location) = convertThroughAndSaveMaxIndex(chunk[0] + destinationIndex, data.maps)
                if minValue > location {
                    minValue = location
                }
                destinationIndex += step
            }
        }
        
        return minValue
    }
}
