import 'dart:io';

class Point {
  Point(this.dx, this.dy);

  int dx;
  int dy;

  bool operator ==(Object? o) => (o is Point) ? (o.dx == dx && o.dy == dy) : false;

  @override
  String toString() => '<dx = $dx | dy = $dy>';

  @override
  int get hashCode => dx ^ dy;

  Point get copy => Point(dx, dy);
}

enum Direction {
  down,
  left,
  up,
  right;

  static Direction parse(String raw) {
    switch (raw) {
      case 'D':
        return down;
      case 'U':
        return up;
      case 'L':
        return left;
      case 'R':
        return right;
      default:
        return right;
    }
  }
}

class Command {
  const Command(this.direction, this.step);

  factory Command.parse(String raw) => Command(Direction.parse(raw[0]), int.parse(raw.split(' ')[1]));

  final Direction direction;
  final int step;
}

List<Command> parseInput(String fileName) => File(fileName).readAsLinesSync().map((e) => Command.parse(e)).toList();

class State {
  State(this.length) {
    _rope = List.generate(length, (index) => Point(0, 0));
  }
  int length;
  late List<Point> _rope;
}

void main() {
  final input = parseInput('../input/day09.txt');
  firstHalfProblem(input);
  lastHalfProblem(input);
}

void _solve(int length, List<Command> commands) {
  final state = State(length);
  final path = Set<Point>();

  var max = Point(0, 0);
  var min = Point(0, 0);

  void moveHead(Command command) {
    switch (command.direction) {
      case Direction.up:
        state._rope[0].dy++;
        break;
      case Direction.down:
        state._rope[0].dy--;
        break;
      case Direction.left:
        state._rope[0].dx--;
        break;
      case Direction.right:
        state._rope[0].dx++;
        break;
    }
    if (max.dx < state._rope[0].dx) max = Point(state._rope[0].dx, max.dy);
    if (max.dy < state._rope[0].dy) max = Point(max.dx, state._rope[0].dy);
    if (min.dx > state._rope[0].dx) min = Point(state._rope[0].dx, min.dy);
    if (min.dy > state._rope[0].dy) min = Point(min.dx, state._rope[0].dy);
  }

  void moveTail() {
    for (var i = 1; i < state.length; i++) {
      final delta = Point(
        state._rope[i - 1].dx - state._rope[i].dx,
        state._rope[i - 1].dy - state._rope[i].dy,
      );
      if (delta.dx.abs() <= 1 && delta.dy.abs() <= 1) return;

      if (delta.dy > 0) {
        state._rope[i].dy++;
      } else if (delta.dy < 0) {
        state._rope[i].dy--;
      }
      if (delta.dx > 0) {
        state._rope[i].dx++;
      } else if (delta.dx < 0) {
        state._rope[i].dx--;
      }
    }
    path.add(state._rope.last.copy);
  }

  for (var command in commands) {
    for (var i = 0; i < command.step; i++) {
      moveHead(command);
      moveTail();
    }
  }

  print(path.length);
}

void firstHalfProblem(List<Command> commands) {
  _solve(2, commands);
}

void lastHalfProblem(List<Command> commands) {
  _solve(10, commands);
}
