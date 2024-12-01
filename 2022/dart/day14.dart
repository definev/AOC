import 'dart:io';
import 'dart:math';

import 'helper/point.dart';

Map<Point, String> parseInput(String fileName) {
  Map<Point, String> map = {};
  final lines = File(fileName).readAsLinesSync();
  for (var line in lines) {
    final coors = line.split(' -> ').map((e) {
      final raw = e.split(',');
      return Point(int.parse(raw[0]), int.parse(raw[1]));
    }).toList();
    if (coors.length == 1) map[coors.first] = '#';
    for (var i = 0; i < coors.length - 1; i++) {
      final first = Point(min(coors[i].dx, coors[i + 1].dx), min(coors[i].dy, coors[i + 1].dy));
      final last = Point(max(coors[i].dx, coors[i + 1].dx), max(coors[i].dy, coors[i + 1].dy));
      for (var dx = first.dx; dx <= last.dx; dx++) {
        for (var dy = first.dy; dy <= last.dy; dy++) {
          map[Point(dx, dy)] = '#';
        }
      }
    }
  }

  return map;
}

void main() {
  // final input = parseInput('../input/day14.txt');
  // firstHalfProblem(input);
  final inputTwo = parseInput('../input/day14.txt');
  lastHalfProblem(inputTwo);
}

void log(Map<Point, String> input) {
  final left = input.keys.fold(9999999, (min, element) => min > element.dx ? element.dx : min);
  final right = input.keys.fold(-1, (max, element) => max < element.dx ? element.dx : max);
  final bottom = input.keys.fold(-1, (max, element) => max < element.dy ? element.dy : max);

  StringBuffer buffer = StringBuffer();

  for (var dy = 0; dy <= bottom; dy++) {
    for (var dx = left; dx <= right; dx++) {
      switch (input[Point(dx, dy)]) {
        case null:
          buffer.write('.');
          break;
        default:
          buffer.write(input[Point(dx, dy)]);
          break;
      }
    }
    buffer.writeln();
  }

  print(buffer);
}

void lastHalfProblem(Map<Point, String> input) {
  int total = 0;
  final maxHeight = input.keys.fold(0, (max, element) => max < element.dy ? element.dy : max) + 2;

  bool loop = true;
  while (loop) {
    var sand = Point(500, 0);
    while (true) {
      if (input[sand] != null) {
        loop = false;
        break;
      }

      final under = Point(sand.dx, sand.dy + 1);
      if (under.dy == maxHeight) {
        input[sand] = 'o';
        total += 1;
        break;
      }

      if (input[under] == null) {
        sand = under;
        continue;
      }

      final leftUnder = Point(under.dx - 1, under.dy);
      if (input[leftUnder] == null) {
        sand = leftUnder;
        continue;
      }

      final rightUnder = Point(under.dx + 1, under.dy);
      if (input[rightUnder] == null) {
        sand = rightUnder;
        continue;
      }

      input[sand.copy] = 'o';
      total += 1;
      break;
    }
  }
  log(input);
  print(total);
}

void firstHalfProblem(Map<Point, String> input) {
  int total = 0;
  final maxHeight = input.keys.fold(0, (max, element) => max < element.dy ? element.dy : max);
  bool loop = true;
  while (loop) {
    var sand = Point(500, 0);
    while (true) {
      final under = Point(sand.dx, sand.dy + 1);
      if (under.dy == maxHeight + 1) {
        loop = false;
        break;
      }

      if (input[under] == null) {
        sand = under;
        continue;
      }

      final leftUnder = Point(under.dx - 1, under.dy);
      if (input[leftUnder] == null) {
        sand = leftUnder;
        continue;
      }

      final rightUnder = Point(under.dx + 1, under.dy);
      if (input[rightUnder] == null) {
        sand = rightUnder;
        continue;
      }

      input[sand.copy] = 'o';
      total += 1;
      break;
    }
  }

  log(input);

  print(total);
}
