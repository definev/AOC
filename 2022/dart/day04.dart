import 'dart:io';
import 'dart:math';

import 'helper/tuple.dart';

List<Tuple<String, String>> parseInput(String fileName) {
  return File(fileName).readAsLinesSync().map((e) {
    final list = e.split(',');
    return Tuple(list[0], list[1]);
  }).toList();
}

void main() {
  final input = parseInput('../input/day04.txt');
  firstHalfProblem(input);
  lastHalfProblem(input);
}

void firstHalfProblem(List<Tuple<String, String>> input) {
  final processedInput = input.map((e) {
    final first = e.first.split('-');
    final second = e.second.split('-');
    return Tuple(
      Tuple(int.parse(first[0]), int.parse(first[1])),
      Tuple(int.parse(second[0]), int.parse(second[1])),
    );
  }).toList();

  int total = 0;

  for (var ele in processedInput) {
    final isSecondInFirst = ele.first.first <= ele.second.first && ele.first.second >= ele.second.second;
    final isFirstInSecond = ele.second.first <= ele.first.first && ele.second.second >= ele.first.second;
    if (isFirstInSecond || isSecondInFirst) {
      total += 1;
    }
  }

  print(total);
}

void lastHalfProblem(List<Tuple<String, String>> input) {
  final processedInput = input.map((e) {
    final first = e.first.split('-');
    final second = e.second.split('-');
    return Tuple(
      Tuple(int.parse(first[0]), int.parse(first[1])),
      Tuple(int.parse(second[0]), int.parse(second[1])),
    );
  }).toList();

  int total = 0;

  for (var ele in processedInput) {
    final range = min(ele.first.second, ele.second.second) - max(ele.first.first, ele.second.first);
    if (range >= 0) {
      total += 1;
    }
  }

  print(total);
}
