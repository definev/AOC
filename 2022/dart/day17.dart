import 'dart:io';

import 'helper/point.dart';

String parseInput(String fileName) => File(fileName).readAsStringSync();

void main() {
  final input = parseInput('../input/day17.txt');
  firstHalfProblem(input);
  lastHalfProblem(input);
}

List<Point> createBrick(int type, int height) {
  switch (type % 5) {
    case 0:
      // ..####.
      return [for (int dx = 3; dx <= 6; dx++) Point(dx, height + 4)];
    case 1:
      // ...#...
      // ..###..
      // ...#...
      return [
        Point(4, height + 6),
        for (int dx = 3; dx <= 5; dx++) Point(dx, height + 5),
        Point(4, height + 4),
      ];
    case 2:
      // ....#..
      // ....#..
      // ..###..
      return [
        Point(5, height + 6),
        Point(5, height + 5),
        for (int dx = 3; dx <= 5; dx++) Point(dx, height + 4),
      ];
    case 3:
      // ..#....
      // ..#....
      // ..#....
      // ..#....
      return [for (int dy = height + 4; dy <= height + 7; dy++) Point(3, dy)];
    default:
      // ..##...
      // ..##...
      return [
        Point(3, height + 4),
        Point(4, height + 4),
        Point(3, height + 5),
        Point(4, height + 5),
      ];
  }
}

class JetRock {
  JetRock(this.jet, this.rock);

  final int jet;
  final int rock;

  @override
  bool operator ==(Object? o) => o is JetRock ? o.jet == jet && o.rock == rock : false;

  @override
  int get hashCode => jet ^ rock;
}

class Checkpoint {
  Checkpoint(this.index, this.height);

  final int index;
  final int height;

  @override
  bool operator ==(Object? o) => o is Checkpoint ? o.index == index && o.height == height : false;

  @override
  int get hashCode => index ^ height;
}

void firstHalfProblem(String jets) {
  _solve(jets, 2022);
}

void lastHalfProblem(String jets) {
  _solve(jets, 1000000000000);
}

void _solve(String jets, int target) {
  Map<JetRock, Checkpoint> checkpoints = {};
  int jet = 0;
  int height = 0;

  Map<Point, bool> marked = {};
  List<Point> current = createBrick(0, 0);

  rockFallin:
  for (var rockIndex = 0; rockIndex < target; rockIndex++) {
    current = createBrick(rockIndex, height);

    final jetRock = JetRock(jet, rockIndex % 5);
    final checkpoint = checkpoints[jetRock];

    if (checkpoint != null) {
      int n = target - rockIndex;
      int del = rockIndex - checkpoint.index;

      if (n % del == 0) {
        print(height + n ~/ del * (height - checkpoint.height));
        return;
      }
    }

    checkpoints[JetRock(jet, rockIndex % 5)] = Checkpoint(rockIndex, height);

    while (true) {
      final jetAction = jets[jet];
      jet = (jet + 1) % jets.length;
      switch (jetAction) {
        case '<':
          bool collidedWall = current.fold(false, (val, p) => val || p.dx == 1);
          bool collidedRock = current.fold(false, (val, p) => val || marked[Point(p.dx - 1, p.dy)] == true);
          if (collidedWall || collidedRock) break;
          current = current.map((e) => e.move(Direction.left)).toList();
          break;
        case '>':
          bool collidedWall = current.fold(false, (val, p) => val || p.dx == 7);
          bool collidedRock = current.fold(false, (val, p) => val || marked[Point(p.dx + 1, p.dy)] == true);
          if (collidedWall || collidedRock) break;
          current = current.map((e) => e.move(Direction.right)).toList();
          break;
      }

      final down = current.map((p) => p.move(Direction.down, true)).toList();
      for (var p in down) {
        if (marked[p] == true || p.dy == 0) {
          for (var p in current) {
            height = p.dy > height ? p.dy : height;
            marked[p] = true;
          }
          continue rockFallin;
        }
      }

      current = down;
    }
  }

  print(height);
}

void display(Map<Point, bool> marked, List<Point> current) {
  int height = [...marked.keys, ...current].fold<int>(0, (height, p) => height > p.dy ? height : p.dy);
  List<List<String>> map = [
    '+-------+'.split(''),
    for (int dy = 1; dy < height + 1; dy++) ['|', ...List.generate(7, (index) => '.'), '|']
  ];

  for (var p in marked.keys) {
    map[p.dy][p.dx] = '#';
  }

  for (var p in current) {
    map[p.dy][p.dx] = '@';
  }

  for (var line in map.reversed) {
    print(line.join());
  }
  print('');
}
