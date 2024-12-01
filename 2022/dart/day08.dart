import 'dart:io';

List<List<int>> parseInput(String fileName) {
  final file = File(fileName).readAsLinesSync();
  return file.map((e) => e.split('').map((e) => int.parse(e)).toList()).toList();
}

void main() {
  final input = parseInput('../input/day08.txt');
  firstHalfProblem(input);
  lastHalfProblem(input);
}

bool canSeek(List<List<int>> input, int i, int j) {
  bool seekableTop = true;
  final center = input[i][j];
  for (var top = i - 1; top >= 0; top--) {
    if (center <= input[top][j]) {
      seekableTop = false;
      break;
    }
  }

  bool seekableDown = true;
  for (var down = i + 1; down < input.length; down++) {
    if (center <= input[down][j]) {
      seekableDown = false;
      break;
    }
  }

  bool seekableLeft = true;
  for (var left = j - 1; left >= 0; left--) {
    if (center <= input[i][left]) {
      seekableLeft = false;
      break;
    }
  }

  bool seekableRight = true;
  for (var right = j + 1; right < input.first.length; right++) {
    if (center <= input[i][right]) {
      seekableRight = false;
      break;
    }
  }
  if (seekableRight || seekableLeft || seekableTop || seekableDown) return true;
  return false;
}

void firstHalfProblem(List<List<int>> input) {
  int total = input.length * 4 - 4;
  for (var i = 1; i < input.length - 1; i++) {
    for (var j = 1; j < input.length - 1; j++) {
      if (canSeek(input, i, j)) {
        total += 1;
      }
    }
  }
  print(total);
}

int calculateView(List<List<int>> input, int i, int j) {
  final center = input[i][j];

  int seekableTop = 1;
  for (var top = i - 1; top >= 0; top--) {
    if (top == 0) break;
    if (center <= input[top][j]) break;
    seekableTop += 1;
  }

  int seekableDown = 1;
  for (var down = i + 1; down < input.length; down++) {
    if (down == input.length - 1) break;
    if (center <= input[down][j]) break;
    seekableDown += 1;
  }

  int seekableLeft = 1;
  for (var left = j - 1; left >= 0; left--) {
    if (left == 0) break;
    if (center <= input[i][left]) break;
    seekableLeft += 1;
  }

  int seekableRight = 1;
  for (var right = j + 1; right < input.length; right++) {
    if (right == input.length - 1) break;
    if (center <= input[i][right]) break;
    seekableRight += 1;
  }

  return seekableLeft * seekableRight * seekableDown * seekableTop;
}

void lastHalfProblem(List<List<int>> input) {
  int maxView = 0;
  for (var i = 1; i < input.length - 1; i++) {
    for (var j = 1; j < input.length - 1; j++) {
      final view = calculateView(input, i, j);
      if (view > maxView) maxView = view;
    }
  }
  print(maxView);
}
