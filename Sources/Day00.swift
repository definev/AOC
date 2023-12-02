import Algorithms

struct Day00: AdventDay {
    // Save your data in a corresponding text file in the `Data` directory.
    var data: String

    // Splits input data into its component parts and convert from string.
    var entities: [String] {
        data.split(separator: "\n").map { String($0) }
    }

    func getDigit(raw: String) -> Int {
        var num: String = ""
        for digit in raw {
            if digit.isNumber {
                num.append(digit)
            }
        }
        let digit1 = num[num.index(num.startIndex, offsetBy: 0)]
        let digit2 = num[num.index(num.startIndex, offsetBy: num.count - 1)]
        return Int("\(digit1)\(digit2)")!
    }

    // Replace this with your solution for the first part of the day's challenge.
    func part1() -> Any {
        var nums: [Int] = []
        for entity in entities {
            nums.append(getDigit(raw: entity))
        }

        return nums.reduce(0, { prev, num in prev + num })
    }

    let digits = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

    func processingDigitWithSring(raw: String) -> Int {
        var num = ""


        var index = 0
    processing:
        while index < raw.count {
            let idx = raw.index(raw.startIndex, offsetBy: index)
            if raw[idx].isNumber {
                num.append(raw[idx])
                index += 1
                continue
            }

            var lastIndex = 0
            for i in 1...5 {
                lastIndex = index + i
                if lastIndex > raw.count {
                    break
                }
                let firstIndexIdx = raw.index(raw.startIndex, offsetBy: index)
                let lastIndexIdx = raw.index(raw.startIndex, offsetBy: lastIndex)
                let digit = String(raw[firstIndexIdx..<lastIndexIdx])
                if let digitIndex = digits.firstIndex(of: digit) {
                    num.append("\(digitIndex + 1)".first!)
                    index = index + 1
                    continue processing
                }
            }
            index += 1
        }

        let digit1 = num[num.index(num.startIndex, offsetBy: 0)]
        let digit2 = num[num.index(num.startIndex, offsetBy: num.count - 1)]
        let result = if num.count == 1 {
            Int("\(digit1)\(digit1)")!
        } else {
            Int("\(digit1)\(digit2)")!
        }
        return result
    }

    // Replace this with your solution for the second part of the day's challenge.
    func part2() -> Any {
        var nums: [Int] = []
        for entity in entities {
            nums.append(processingDigitWithSring(raw: entity))
        }

        return nums.reduce(0, { prev, num in prev + num })
    }
}
