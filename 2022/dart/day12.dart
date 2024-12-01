import 'dart:collection';
import 'dart:io';

import 'helper/point.dart';
import 'helper/tuple.dart';

class Journey {
  Journey(this.map, this.start, this.end);

  final List<List<String>> map;
  final Point start;
  final Point end;

  String getHeight(Point p) => map[p.dy][p.dx];

  bool isOutOfBound(Point p) => p.dx < 0 || p.dy < 0 || p.dx > map.first.length - 1 || p.dy > map.length - 1;

  void log(List<Point> path, Point current) {
    StringBuffer buffer = StringBuffer();
    for (var i = 0; i < map.length; i++) {
      for (var j = 0; j < map.first.length; j++) {
        buffer.write(Point(j, i) == current
            ? 'X'
            : path.contains(Point(j, i))
                ? '@'
                : '.');
      }
      buffer.writeln();
    }

    print(buffer.toString());
  }
}

Journey parseInput(String fileName) {
  final lines = File(fileName).readAsLinesSync();

  List<List<String>> map = List.generate(lines.length, (index) => []);
  Point start = Point(0, 0);
  Point end = Point(0, 0);

  for (var i = 0; i < lines.length; i++) {
    final line = lines[i];

    for (var j = 0; j < line.length; j++) {
      map[i].add(line[j]);
      if (line[j] == 'S') start = Point(j, i);
      if (line[j] == 'E') end = Point(j, i);
    }
  }

  return Journey(map, start, end);
}

void main() {
  final input = parseInput('../input/day12.txt');
  firstHalfProblem(input);
  lastHalfProblem(input);
}

/**
 * Sabqponm
 * abcryxxl
 * accszExk
 * acctuvwj
 * abdefghi
 */

void firstHalfProblem(Journey input) {
  Queue<Tuple<Point, int>> path = Queue.from([Tuple(input.start, 0)]);

  List<List<int>> heightMap = [
    for (final line in input.map)
      line.map(
        (e) {
          if (e == 'S') return 1;
          if (e == 'E') return 26;
          return e.codeUnitAt(0) - 'a'.codeUnitAt(0) + 1;
        },
      ).toList(),
  ];

  Set<Point> visited = Set();

  while (path.isNotEmpty) {
    final current = path.removeFirst();

    if (visited.contains(current.first)) continue;
    visited.add(current.first);

    if (input.end == current.first) {
      print(current.second);
      visited.remove(current.first);
      continue;
    }

    bool canMove = false;
    for (final direction in Direction.values) {
      final next = current.first.move(direction);
      if (input.isOutOfBound(next)) continue;
      if (heightMap[next.dy][next.dx] - heightMap[current.first.dy][current.first.dx] > 1) continue;

      canMove = true;
      path.add(Tuple(next, current.second + 1));
    }

    if (!canMove) visited.remove(current.first);
  }
}

void lastHalfProblem(Journey input) {
  Queue<Tuple<Point, int>> path = Queue.from([
    Tuple(input.start, 0),
  ]);

  for (var i = 0; i < input.map.length; i++) {
    for (var j = 0; j < input.map.first.length; j++) {
      if (input.map[i][j] == 'a') path.add(Tuple(Point(j, i), 0));
    }
  }

  List<List<int>> heightMap = [
    for (final line in input.map)
      line.map(
        (e) {
          if (e == 'S') return 1;
          if (e == 'E') return 26;
          return e.codeUnitAt(0) - 'a'.codeUnitAt(0) + 1;
        },
      ).toList(),
  ];

  Set<Point> visited = Set();

  while (path.isNotEmpty) {
    final current = path.removeFirst();

    if (visited.contains(current.first)) continue;
    visited.add(current.first);

    if (input.end == current.first) {
      print(current.second);
      visited.remove(current.first);
      continue;
    }

    bool canMove = false;
    for (final direction in Direction.values) {
      final next = current.first.move(direction);
      if (input.isOutOfBound(next)) continue;
      if (heightMap[next.dy][next.dx] - heightMap[current.first.dy][current.first.dx] > 1) continue;

      canMove = true;
      path.add(Tuple(next, current.second + 1));
    }

    if (!canMove) visited.remove(current.first);
  }
}
