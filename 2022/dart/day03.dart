import 'dart:io';

import 'helper/tuple.dart';

Map<String, int> scoreMap = {};

int fetchScore() {
  for (int i = 97; i < 97 + 26; i++) scoreMap[String.fromCharCode(i)] = i - 96;
  for (int i = 65; i < 65 + 26; i++) scoreMap[String.fromCharCode(i)] = i - 64 + 26;
  return 0;
}

List<String> parseInput(String fileName) => File(fileName).readAsLinesSync();

List<Tuple<String, String>> parseFirstProblemInput(List<String> input) => input //
    .map((e) => Tuple(e.substring(0, e.length ~/ 2), e.substring(e.length ~/ 2, e.length)))
    .toList();

List<Triple<String, String, String>> parseLastProblemInput(List<String> input) {
  final result = <Triple<String, String, String>>[];
  for (int i = 0; i < input.length ~/ 3; i++) {
    result.add(
      Triple(
        input[i * 3],
        input[i * 3 + 1],
        input[i * 3 + 2],
      ),
    );
  }

  return result;
}

void main() {
  final input = parseInput('../input/day03.txt');
  fetchScore();
  // firstHalfProblem(parseFirstProblemInput(input));
  lastHalfProblem(parseLastProblemInput(input));
}

void firstHalfProblem(List<Tuple<String, String>> input) {
  final chars = scoreMap.keys.toList();
  String sameStr = '';
  for (var line in input) {
    var left = line.first;
    var right = line.second;
    for (var char in chars) {
      var newLeft = line.first.replaceAll(char, '');
      var newRight = line.second.replaceAll(char, '');
      if (newLeft.length < left.length && newRight.length < right.length) {
        sameStr = '$sameStr$char';
        left = newLeft;
        right = newRight;
      }
    }
  }

  print(sameStr.split('').fold<int>(0, (total, char) => total + scoreMap[char]!));
}

bool _isSameChar(List<String> list, List<String> newList) {
  bool result = true;
  for (var i = 0; i < list.length; i++) {
    result = result && list[i].length != newList[i].length;
  }
  return result;
}

void lastHalfProblem(List<Triple<String, String, String>> input) {
  final chars = scoreMap.keys.toList();
  String sameStr = '';
  for (var line in input) {
    var list = [line.first, line.second, line.third];
    for (var char in chars) {
      var newList = list.map((e) => e.replaceAll(char, '')).toList();
      if (_isSameChar(list, newList)) {
        sameStr = '$sameStr$char';
        list = newList;
      }
    }
  }

  print(sameStr.split('').fold<int>(0, (total, char) => total + scoreMap[char]!));
}
