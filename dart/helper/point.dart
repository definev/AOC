class Point {
  Point(this.dx, this.dy);

  int dx;
  int dy;

  @override
  String toString() => '<dx = $dx | dy = $dy>';

  @override
  int get hashCode => dx ^ dy;

  bool operator ==(Object? o) => (o is Point) ? (o.dx == dx && o.dy == dy) : false;

  Point get copy => Point(dx, dy);

  Point operator -(Point other) => Point(dx - other.dx, dy - other.dy);
  Point operator +(Point other) => Point(dx + other.dx, dy + other.dy);
  Point operator *(Point other) => Point(dx * other.dx, dy * other.dy);

  Point move(Direction direction) {
    switch (direction) {
      case Direction.down:
        return Point(dx, dy + 1);
      case Direction.up:
        return Point(dx, dy - 1);
      case Direction.left:
        return Point(dx - 1, dy);
      case Direction.right:
        return Point(dx + 1, dy);
    }
  }
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

  Direction get opposite {
    switch (this) {
      case down:
        return up;
      case up:
        return down;
      case left:
        return right;
      case right:
        return left;
    }
  }
}
