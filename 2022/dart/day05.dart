import 'dart:io';

import 'helper/tuple.dart';

class Command {
  const Command(this.quantity, {required this.from, required this.to});
  factory Command.parse(String input) {
    final raw = input.split(' ');
    return Command(
      int.parse(raw[1]),
      from: int.parse(raw[3]) - 1,
      to: int.parse(raw[5]) - 1,
    );
  }

  final int quantity;
  final int from;
  final int to;

  @override
  String toString() {
    return 'move $quantity from $from to $to';
  }
}

typedef CrateStack = List<String>;

List<CrateStack> _parseCrates(String input) {
  final raw = input.split('\n');
  final length = int.parse(raw.last.trim().split('   ').last);
  raw.removeLast();

  List<CrateStack> result = List.generate(length, (_) => []);

  for (int i = raw.length - 1; i >= 0; i -= 1) {
    final line = raw[i];
    for (int j = 0; j < length; j += 1) {
      final crate = line.substring(j * 4, j * 4 + 3);
      if (crate != '   ') {
        result[j].add(crate);
      }
    }
  }

  return result;
}

List<Command> _parseCommands(String input) {
  final raw = input.split('\n');
  return raw.map((e) => Command.parse(e)).toList();
}

Tuple<List<CrateStack>, List<Command>> parseInput(String fileName) {
  final file = File(fileName).readAsStringSync();
  final raw = file.split('\n\n');
  final crates = _parseCrates(raw[0]);
  final commands = _parseCommands(raw[1]);
  return Tuple(crates, commands);
}

void main() {
  firstHalfProblem('../input/day05.txt');
  lastHalfProblem('../input/day05.txt');
}

void _execute(List<CrateStack> stacks, Command command, {bool reversed = true}) {
  final from = stacks[command.from];
  final moveStack = () {
    if (reversed) return from.sublist(from.length - command.quantity).reversed;
    return from.sublist(from.length - command.quantity);
  }();
  from.removeRange(from.length - command.quantity, from.length);
  stacks[command.to] = [...stacks[command.to], ...moveStack];
}

void firstHalfProblem(String fileName) {
  final input = parseInput(fileName);
  List<CrateStack> result = input.first;

  for (final command in input.second) {
    _execute(result, command);
  }

  print(result.fold('', (prev, ele) => '$prev${ele.last[1]}'));
}

void lastHalfProblem(String fileName) {
  final input = parseInput(fileName);
  List<CrateStack> result = input.first;

  for (final command in input.second) {
    _execute(result, command, reversed: false);
  }

  print(result.fold('', (prev, ele) => '$prev${ele.last[1]}'));
}
