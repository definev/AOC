import 'dart:io';

class Op {
  Op(this.cycle, this.value);
  factory Op.parse(String raw) {
    final list = raw.split(' ');
    if (list[0] == 'noop') return noop;
    return Op(2, int.parse(list[1]));
  }

  static Op noop = Op(1, 0);

  final int cycle;
  final int value;
}

List<Op> parseInput(String fileName) => File(fileName).readAsLinesSync().map(Op.parse).toList();

void main() {
  final input = parseInput('../input/day10.txt');
  firstHalfProblem(input);
  lastHalfProblem(input);
}

void firstHalfProblem(List<Op> input) {
  int cycle = 0;
  int value = 1;
  int opPtr = 0;

  int result = 0;
  final onTimeOp = <int>[];
  for (var i = 0; i < input.length; i++) {
    int time = input[i].cycle;
    for (int j = 0; j < i; j++) {
      time += input[j].cycle;
    }
    onTimeOp.add(time);
  }

  while (opPtr < input.length) {
    cycle++;
    final op = input[opPtr];

    if ((cycle - 20) % 40 == 0 && cycle <= 220) {
      result += value * cycle;
    }

    if (onTimeOp[opPtr] == cycle) {
      value += op.value;
      opPtr++;
    }
  }

  print(result);
}

void lastHalfProblem(List<Op> input) {
  int cycle = 0;
  int value = 0;
  int opPtr = 0;

  String result = '';
  final onTimeOp = <int>[];
  for (var i = 0; i < input.length; i++) {
    int time = input[i].cycle;
    for (int j = 0; j < i; j++) {
      time += input[j].cycle;
    }
    onTimeOp.add(time);
  }

  while (opPtr < input.length) {
    final op = input[opPtr];

    cycle++;

    final crtPos = cycle % 40;
    final spritePos = value % 40;

    if (crtPos >= spritePos + 1 && crtPos < spritePos + 4) {
      result += '#';
    } else {
      result += '.';
    }

    if (cycle % 40 == 0) {
      print(result);
      result = '';
    }

    if (onTimeOp[opPtr] == cycle) {
      value += op.value;
      opPtr++;
    }
  }
}
