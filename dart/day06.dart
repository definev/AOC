import 'dart:io';

String parseInput(String fileName) {
  return File(fileName).readAsStringSync();
}

void main() {
  final input = parseInput('../input/day06.txt');
  firstHalfProblem(input);
  lastHalfProblem(input);
}

int processFrame(String input, int frameWidth) {
  int index = 0;

  while (index < input.length) {
    String framePage = input.substring(index, index + frameWidth);
    if (framePage.runes.toSet().length == frameWidth) {
      return index + frameWidth;
    }
    index++;
  }

  return -1;
}

void firstHalfProblem(String input) {
  print(processFrame(input, 4));
}

void lastHalfProblem(String input) {
  print(processFrame(input, 14));
}
