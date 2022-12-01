import 'dart:io';
import 'dart:math';

class Tuple<F, S> {
  const Tuple(this.first, this.second);

  final F first;
  final S second;
}

void main() {
  final input = File('../input/day01.txt').readAsStringSync();

  firstHalfProblem(input);
  lastHalfProblem(input);
}

void firstHalfProblem(String input) {
  final result = input //
      .split('\n\n')
      .map((raw) => raw.split('\n').map((cal) => int.parse(cal.trim())))
      .fold(-1, _filterFirstElf);

  print(result);
}

int _filterFirstElf(int prev, Iterable<int> elf) {
  final elfCal = elf.fold(0, (prev, curr) => prev + curr);
  return max(prev, elfCal);
}

/// _SECOND PART_

void lastHalfProblem(String input) {
  final result = input //
      .split('\n\n')
      .map((raw) => raw.split('\n').map((cal) => int.parse(cal.trim())))
      .fold(<int>[], _filterMostThree) //
      .fold(0, (prev, curr) => prev + curr);

  print(result);
}

List<int> _filterMostThree(List<int> prev, Iterable<int> elf) {
  final elfCal = elf.fold(0, (prev, curr) => prev + curr);
  if (prev.length < 3) {
    return [...prev, elfCal];
  }

  prev.sort((a, b) => a.compareTo(b));
  var canReplace = Tuple(false, -1);

  for (var i = 0; i < 3; i++) {
    canReplace = Tuple(canReplace.first || prev[i] < elfCal, i);
    break;
  }

  if (canReplace.first) {
    return [...prev]..[canReplace.second] = elfCal;
  }

  return prev;
}
