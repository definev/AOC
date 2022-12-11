import 'dart:io';

import 'helper/tuple.dart';

class MonkeyData {
  MonkeyData(this.items, this.operation, this.testcase);

  factory MonkeyData.parse(String raw) {
    final lines = raw.split('\n');
    List<int> items = lines[1] //
        .substring('  Starting items: '.length)
        .split(', ')
        .map((e) => int.parse(e))
        .toList();
    String operation = lines[2].substring('  Operation: '.length);

    int divisibleBy = int.parse(lines[3].substring('  Test: divisible by '.length));
    int trueIndex = int.parse(lines[4].substring('    If true: throw to monkey '.length));
    int falseIndex = int.parse(lines[5].substring('    If false: throw to monkey '.length));

    return MonkeyData(
      items,
      operation,
      Testcase(divisibleBy, trueIndex, falseIndex),
    );
  }

  List<int> items;
  final String operation;
  final Testcase testcase;

  @override
  String toString() {
    return '\nSTARTING ITEM : $items\nOPERATION : $operation\nTESTCASE : $testcase\n';
  }
}

class Testcase {
  Testcase(this.divisibleBy, this.trueIndex, this.falseIndex);

  final int divisibleBy;
  final int trueIndex;
  final int falseIndex;

  @override
  String toString() {
    return '\n--> DIVISBLE_BY : $divisibleBy\n--> TRUE : $trueIndex\n--> FALSE : $falseIndex';
  }
}

List<MonkeyData> parseInput(String fileName) => File(fileName) //
    .readAsStringSync()
    .split('\n\n')
    .map(MonkeyData.parse)
    .toList();

int evalWorryLevel(String operation, int worryLevel) {
  final raw = operation.trim().split(' = ')[1].split(' ');
  final first = int.tryParse(raw[0]) ?? worryLevel;
  final second = int.tryParse(raw[2]) ?? worryLevel;

  switch (raw[1]) {
    case '+':
      return first + second;
    case '*':
      return first * second;
    default:
      throw Exception('Parse wrong $operation');
  }
}

void main() {
  final input = parseInput('../sample/day11.txt');
  final watch = Stopwatch();
  watch.start();
  firstHalfProblem(input);
  lastHalfProblem(input);
  print('TIME : ${watch.elapsedMilliseconds} ms');
  watch.reset();
}

void firstHalfProblem(List<MonkeyData> input) {
  List<int> monkeyTimes = List.generate(input.length, (index) => 0);
  for (var i = 0; i < 20; i++) {
    for (var index = 0; index < input.length; index++) {
      final monkey = input[index];

      monkeyTimes[index] += monkey.items.length;

      List<Tuple<int, int>> moveIndex = [];
      for (var i = 0; i < monkey.items.length; i++) {
        int item = int.parse(monkey.items[i].toString());
        item = evalWorryLevel(monkey.operation, item);
        item = item ~/ 3;

        if (item % monkey.testcase.divisibleBy == 0) {
          moveIndex.add(Tuple(i, monkey.testcase.trueIndex));
        } else {
          moveIndex.add(Tuple(i, monkey.testcase.falseIndex));
        }

        monkey.items[i] = item;
      }

      for (var move in moveIndex) {
        input[move.second].items.add(monkey.items[move.first]);
      }
      List<int> newItems = [];
      for (var i = 0; i < monkey.items.length; i++) {
        if (moveIndex.map((e) => e.first).contains(i)) continue;
        newItems.add(monkey.items[i]);
      }
      monkey.items = newItems;
    }
  }

  monkeyTimes.sort((a, b) => a.compareTo(b));
  print(monkeyTimes[monkeyTimes.length - 1] * monkeyTimes[monkeyTimes.length - 2]);
}

void printTimes(int round, List<int> times) {
  print('== After round $round ==');
  for (var i = 0; i < times.length; i++) {
    print('Monkey $i inspected items ${times[i]} times.');
  }
  print('\n');
}

void lastHalfProblem(List<MonkeyData> input) {
  List<int> monkeyTimes = List.generate(input.length, (index) => 0);
  final modulo = input.fold(1, (val, ele) => val * ele.testcase.divisibleBy);

  for (var i = 0; i < 10000; i++) {
    for (var index = 0; index < input.length; index++) {
      final monkey = input[index];

      monkeyTimes[index] += monkey.items.length;

      List<Tuple<int, int>> moveIndex = [];
      for (var i = 0; i < monkey.items.length; i++) {
        int item = monkey.items[i];
        item = evalWorryLevel(monkey.operation, item) % modulo;
        if (item % monkey.testcase.divisibleBy == 0) {
          moveIndex.add(Tuple(i, monkey.testcase.trueIndex));
        } else {
          moveIndex.add(Tuple(i, monkey.testcase.falseIndex));
        }

        monkey.items[i] = item;
      }
      for (var move in moveIndex) {
        input[move.second].items.add(monkey.items[move.first]);
      }
      List<int> newItems = [];
      for (var i = 0; i < monkey.items.length; i++) {
        if (moveIndex.map((e) => e.first).contains(i)) continue;
        newItems.add(monkey.items[i]);
      }
      monkey.items = newItems;
    }
  }
  monkeyTimes.sort((a, b) => a.compareTo(b));
  print(monkeyTimes[monkeyTimes.length - 1] * monkeyTimes[monkeyTimes.length - 2]);
}
