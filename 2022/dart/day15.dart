import 'dart:io';
import 'dart:math';

import 'helper/algo.dart';
import 'helper/point.dart';
import 'helper/tuple.dart';

List<Tuple<Point, Point>> parseInput(String fileName) {
  List<Tuple<Point, Point>> points = [];
  final lines = File(fileName).readAsLinesSync();

  for (var line in lines) {
    line = line.substring('Sensor at '.length);
    final p = line.split(': closest beacon is at ');
    final sensor = p[0];
    final rawSensor = sensor.split(', ');
    final beacon = p[1];
    final rawBeacon = beacon.split(', ');

    points.add(
      Tuple(
        Point(
          int.parse(rawSensor[0].substring(2)),
          int.parse(rawSensor[1].substring(2)),
        ),
        Point(
          int.parse(rawBeacon[0].substring(2)),
          int.parse(rawBeacon[1].substring(2)),
        ),
      ),
    );
  }

  return points;
}

void main() {
  final input = parseInput('../input/day15.txt');
  firstHalfProblem(input);
  lastHalfProblem(input);
}

class Border {
  const Border({
    required this.radius,
    required this.center,
  });

  factory Border.fromSAndB({
    required Point sensor,
    required Point beacon,
  }) {
    var delta = sensor - beacon;
    delta = Point(delta.dx.abs(), delta.dy.abs());
    final radius = delta.dx + delta.dy;
    return Border(radius: radius, center: sensor);
  }

  final int radius;
  final Point center;

  Set<Point> lineAt(int dy) {
    int min = center.dy - radius;
    int max = center.dy + radius;
    if (dy > max || dy < min) return {};

    int delta = (dy - center.dy).abs();
    int range = radius - delta;
    return {for (int dx = center.dx - range; dx <= center.dx + range; dx++) Point(dx, dy)};
  }

  Tuple<Point, Point>? lineAtRange(int dy) {
    int min = center.dy - radius;
    int max = center.dy + radius;
    if (dy > max || dy < min) return null;

    int delta = (dy - center.dy).abs();
    int range = radius - delta;
    return Tuple(Point(center.dx, dy) - Point(range, 0), Point(center.dx, dy) + Point(range, 0));
  }
}

extension on List<Border> {
  Set<Point> lineAt(int dy) {
    Set<Point> points = {};
    for (var border in this) {
      points.addAll(border.lineAt(dy));
    }
    return points;
  }

  Set<Tuple<Point, Point>> lineAtRange(int dy) {
    Set<Tuple<Point, Point>> points = {};
    for (var border in this) {
      final range = border.lineAtRange(dy);
      if (range != null) points.add(range);
    }
    return points;
  }
}

void firstHalfProblem(List<Tuple<Point, Point>> input) {
  final line = 2000000;
  final borders = input //
      .map((e) => Border.fromSAndB(sensor: e.first, beacon: e.second))
      .toList();

  Set<Point> points = borders.lineAt(line);
  final beacons = input.map((e) => e.second).toList();
  for (var beacon in beacons) {
    points.remove(beacon);
  }

  print(points.length);
}

void lastHalfProblem(List<Tuple<Point, Point>> input) {
  final top = input.fold(9999999, (m, e) => m > min(e.first.dy, e.second.dy) ? min(e.first.dy, e.second.dy) : m);
  final bottom = input.fold(-9999999, (m, e) => m < max(e.first.dy, e.second.dy) ? max(e.first.dy, e.second.dy) : m);
  final borders = input //
      .map((e) => Border.fromSAndB(sensor: e.first, beacon: e.second))
      .toList();

  for (var dy = max(0, top); dy <= min(4000000, bottom); dy++) {
    final points = selectionSort(
      borders.lineAtRange(dy).toList(),
      compareTo: (first, last) => first.second.dx.compareTo(last.second.dx),
    );
    if (points.isEmpty) continue;

    bool contain(int dx) {
      for (var point in points) {
        if (dx >= point.first.dx && dx <= point.second.dx) return true;
      }
      return false;
    }

    for (var ptr = 0; ptr < points.length - 1; ptr++) {
      final lastPoint = points[ptr].second.dx;
      final firstPoint = points[ptr + 1].first.dx;
      if (firstPoint - lastPoint == 2 && !contain(lastPoint + 1)) {
        print(4000000 * (lastPoint + 1) + dy);
      }
    }
  }
}
