import 'dart:io';
import 'dart:math';

enum Axis { dx, dy, dz }

class Position {
  const Position(this.x, this.y, this.z);

  factory Position.parse(String raw) {
    final list = raw.split(',');
    return Position(int.parse(list[0]), int.parse(list[1]), int.parse(list[2]));
  }

  final int x;
  final int y;
  final int z;

  @override
  String toString() => '$x,$y,$z';

  Position operator -(Position position) => Position(x - position.x, y - position.y, z - position.z);

  Position operator +(Position position) => Position(x + position.x, y + position.y, z + position.z);

  Position operator *(Position position) => Position(x * position.x, y * position.y, z * position.z);

  Position operator /(Position position) => Position(x ~/ position.x, y ~/ position.y, z ~/ position.z);

  bool operator ==(Object? o) => o is Position ? o.x == x && o.y == y && o.z == z : false;

  @override
  int get hashCode => x ^ (y * 2) ^ (z * 3);

  Position copyWith({int? x, int? y, int? z}) => Position(x ?? this.x, y ?? this.y, z ?? this.z);
}

List<Position> parseInput(String fileName) => File(fileName).readAsLinesSync().map(Position.parse).toList();

void main() {
  final input = parseInput('../input/day18.txt');
  firstHalfProblem(input);
  lastHalfProblem(input);
}

int getSide(List<Position> input) {
  int side = 0;
  for (var position in input) {
    int initialSide = 6;
    for (var other in input) {
      if (position == other) continue;
      final delta = position - other;
      if (delta.x == 1 && delta.y == 0 && delta.z == 0) initialSide -= 1;
      if (delta.x == -1 && delta.y == 0 && delta.z == 0) initialSide -= 1;
      if (delta.x == 0 && delta.y == 1 && delta.z == 0) initialSide -= 1;
      if (delta.x == 0 && delta.y == -1 && delta.z == 0) initialSide -= 1;
      if (delta.x == 0 && delta.y == 0 && delta.z == 1) initialSide -= 1;
      if (delta.x == 0 && delta.y == 0 && delta.z == -1) initialSide -= 1;
    }
    side += initialSide;
  }
  return side;
}

void firstHalfProblem(List<Position> input) {
  final side = getSide(input);
  print(side);
}

void lastHalfProblem(List<Position> input) {
  final minPos =
      input.fold<Position>(Position(40, 40, 40), (m, p) => Position(min(m.x, p.x), min(m.y, p.y), min(m.z, p.z))) -
          Position(1, 1, 1);
  final maxPos =
      input.fold<Position>(Position(-1, -1, -1), (m, p) => Position(max(m.x, p.x), max(m.y, p.y), max(m.z, p.z))) +
          Position(1, 1, 1);

  final result = <Position>{};
  final neighbours = [
    Position(1, 0, 0),
    Position(-1, 0, 0),
    Position(0, 1, 0),
    Position(0, -1, 0),
    Position(0, 0, 1),
    Position(0, 0, -1),
  ];
  final queue = [minPos];

  while (queue.isNotEmpty) {
    final p = queue.removeLast();

    if ((minPos.x <= p.x && p.x <= maxPos.x) &&
        (minPos.y <= p.y && p.y <= maxPos.y) &&
        (minPos.z <= p.z && p.z <= maxPos.z) &&
        !input.contains(p) &&
        !result.contains(p)) {
      result.add(p);
      queue.addAll(neighbours.map((e) => Position(p.x + e.x, p.y + e.y, p.z + e.z)));
    }
  }
  final allPos = [
    for (int x = minPos.x; x <= maxPos.x; x++)
      for (int y = minPos.y; y <= maxPos.y; y++)
        for (int z = minPos.z; z <= maxPos.z; z++) Position(x, y, z)
  ];

  for (var p in result) {
    allPos.remove(p);
  }
  for (var p in input) {
    allPos.remove(p);
  }

  print(getSide(input) - getSide(allPos));
}
