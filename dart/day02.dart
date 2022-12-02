import 'dart:io';

const _compareMap = {
  'rock rock': 0,
  'paper paper': 0,
  'scissors scissors': 0,
  'rock paper': -1,
  'paper scissors': -1,
  'scissors rock': -1,
  'paper rock': 1,
  'scissors paper': 1,
  'rock scissors': 1,
};

enum Move implements Comparable<Move> {
  rock,
  paper,
  scissors;

  int compareTo(Move other) => _compareMap['${this.name} ${other.name}']!;

  int get score => this.index + 1;
}

class MoveParser {
  const MoveParser(this._rock, this._paper, this._scissors);

  final String _rock;
  final String _paper;
  final String _scissors;

  static const opponent = MoveParser('A', 'B', 'C');

  Move? parse(String raw) {
    if (raw == _rock) return Move.rock;
    if (raw == _paper) return Move.paper;
    if (raw == _scissors) return Move.scissors;
    return null;
  }
}

class ConditionParser {
  const ConditionParser(this._lose, this._draw, this._win);

  final String _lose;
  final String _draw;
  final String _win;

  Move? parse(Move opponent, String raw) {
    if (raw == _lose) {
      final index = opponent.index - 1 < 0 ? 2 : opponent.index - 1;
      return Move.values[index];
    }
    if (raw == _draw) {
      return opponent;
    }
    if (raw == _win) {
      final index = opponent.index + 1;
      return Move.values[index % 3];
    }
    return null;
  }
}

class Turn {
  Turn(this._opponentMove, this._move);

  final Move _opponentMove;
  final Move _move;

  int get score {
    var score = _move.score;
    final result = _move.compareTo(_opponentMove);
    if (result == 1) score += 6;
    if (result == 0) score += 3;
    return score;
  }

  static Turn parseByMove(
    String raw, {
    MoveParser opponentParser = MoveParser.opponent,
    required MoveParser parser,
  }) {
    final rawList = raw.split(' ');
    assert(rawList.length == 2);
    return Turn(opponentParser.parse(rawList[0])!, parser.parse(rawList[1])!);
  }

  static Turn parseByCondition(
    String raw, {
    MoveParser opponentParser = MoveParser.opponent,
    required ConditionParser parser,
  }) {
    final rawList = raw.split(' ');
    assert(rawList.length == 2);
    final opponentMove = opponentParser.parse(rawList[0])!;
    return Turn(opponentMove, parser.parse(opponentMove, rawList[1])!);
  }
}

List<String> parseInput() {
  final inputFile = File('../input/day02.txt');
  final rawInput = inputFile.readAsStringSync();
  return rawInput.split('\n').where((input) => input.isNotEmpty).map((e) => e.trim()).toList();
}

void main() {
  final input = parseInput();
  firstHalfProblem(input);
  lastHalfProblem(input);
}

void firstHalfProblem(List<String> input) {
  final parser = MoveParser('X', 'Y', 'Z');
  int score = 0;
  for (var line in input) {
    score += Turn.parseByMove(line, parser: parser).score;
  }
  print(score);
}

void lastHalfProblem(List<String> input) {
  final parser = ConditionParser('X', 'Y', 'Z');
  int score = 0;
  for (var line in input) {
    score += Turn.parseByCondition(line, parser: parser).score;
  }
  print(score);
}
