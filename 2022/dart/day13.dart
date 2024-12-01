import 'dart:convert';
import 'dart:io';

import 'helper/algo.dart';
import 'helper/tuple.dart';

List<Tuple<dynamic, dynamic>> parseInput(String fileName) => File(fileName) //
        .readAsStringSync()
        .split('\n\n')
        .map((e) {
      final raw = e.split('\n');
      return Tuple(jsonDecode(raw[0]), jsonDecode(raw[1]));
    }).toList();

void main() {
  final input = parseInput('../input/day13.txt');
  // firstHalfProblem(input);
  lastHalfProblem(input);
}

bool? compare(dynamic left, dynamic right) {
  if (left is int && right is int) {
    if (left == right) return null;
    return left < right;
  }
  if (left is List && right is List) {
    int index = 0;
    if (left.length == 0 && right.length == 0) return null;
    if (left.length == 0) return true;
    if (right.length == 0) return false;

    while (true) {
      if (index >= left.length && index >= right.length) return null;
      if (index >= left.length && index < right.length) return true;
      if (index < left.length && index >= right.length) return false;

      final indexCompare = compare(left[index], right[index]);
      if (indexCompare == null) {
        index++;
        continue;
      }
      return indexCompare;
    }
  }
  if (left is int && right is List) {
    return compare([left], right);
  }

  if (left is List && right is int) {
    return compare(left, [right]);
  }

  throw Exception('Cannot happen');
}

void firstHalfProblem(List<Tuple<dynamic, dynamic>> input) {
  int total = 0;

  for (var pair in input) {
    if (compare(pair.first, pair.second) == true) {
      print('${input.indexOf(pair)}');
      total += input.indexOf(pair) + 1;
    }
  }

  print(total);
}

void lastHalfProblem(List<Tuple<dynamic, dynamic>> input) {
  List<dynamic> list = [];

  for (var tuple in input) {
    list.add(tuple.first);
    list.add(tuple.second);
  }

  dynamic pivotOne = [
    [2]
  ];
  dynamic pivotTwo = [
    [6]
  ];
  list.addAll([pivotOne, pivotTwo]);

  list = selectionSort(
    list,
    compareTo: (first, last) {
      final compared = compare(first, last);
      if (compared == null) return 0;
      if (compared == false) return 1;
      return -1;
    },
  );

  print((list.indexOf(pivotOne) + 1) * (list.indexOf(pivotTwo) + 1));
}
