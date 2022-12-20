import 'dart:io';
import 'dart:math';

import 'helper/trotter/trotter.dart';
import 'helper/tuple.dart';

class Valve {
  Valve(this.id, this.rate, this.linkedValves);

  factory Valve.parse(String raw) {
    raw = raw.substring('Valve '.length);
    final list = raw.split(' has flow rate=');
    final tunnelAndValves = list[1].split('; tunnels lead to valves ');
    if (tunnelAndValves.length == 2) {
      return Valve(list[0], int.parse(tunnelAndValves[0]), tunnelAndValves[1].split(', '));
    }
    final tunnelAndValve = list[1].split('; tunnel leads to valve ');
    if (tunnelAndValve.length == 2) {
      return Valve(list[0], int.parse(tunnelAndValve[0]), [tunnelAndValve[1]]);
    }
    throw Exception('Parsing fail');
  }

  final String id;
  final int rate;
  final List<String> linkedValves;
  Map<String, int> costs = {};

  @override
  String toString() => '"ID: $id | RATE: $rate | LINKED_VALVE: $linkedValves"';
}

Map<String, Valve> parseInput(String fileName) {
  final valves = File(fileName).readAsLinesSync().map((e) => Valve.parse(e));
  final input = Map<String, Valve>.fromIterables(valves.map((e) => e.id), valves);

  for (var valve in input.entries) {
    Map<String, int> costs = {valve.key: 0};

    for (var cost = 0; costs.length < input.length; cost++) {
      final currentValves = costs //
          .entries
          .where((entry) => entry.value == cost)
          .map((e) => e.key)
          .toList();

      for (var valve in currentValves) {
        final data = input[valve]!;
        for (var linkedValve in data.linkedValves) {
          if (!costs.containsKey(linkedValve)) {
            costs[linkedValve] = cost + 1;
          }
        }
      }
    }

    valve.value.costs = costs;
  }

  return input;
}

void main() {
  final input = parseInput('../input/day16.txt');

  // firstHalfProblem(input);
  lastHalfProblem(input);
}

List<Tuple<Set<String>, Set<String>>> _shuffle(Set<String> usableValves, int length) {
  final comb = Combinations<String>(length, usableValves.toList());

  return comb() //
      .map((e) => Tuple(e.toSet(), {...usableValves}..removeAll(e.toSet())))
      .toList();
}

void lastHalfProblem(Map<String, Valve> input) {
  final usableValves = Map.fromEntries(input.entries.where((e) => e.value.rate > 0)).keys.toSet();

  int best = 0;
  for (var index = 1; index <= usableValves.length ~/ 2; index++) {
    final usablePairs = _shuffle(usableValves, index);

    for (var pair in usablePairs) {
      final firstBest = _bestPath(input, pair.first, 26, 'AA');
      final secondBest = _bestPath(input, pair.second, 26, 'AA');
      if (best < firstBest + secondBest) {
        print('NEW BEST : $best');
        best = max(best, firstBest + secondBest);
      }
    }
  }

  print(best);
}

void firstHalfProblem(Map<String, Valve> input) {
  final usableValves = Map.fromEntries(input.entries.where((e) => e.value.rate > 0)).keys.toSet();
  int remainMinutes = 30;
  int best = _bestPath(input, usableValves, remainMinutes, 'AA');
  print(best);
}

/// INPUT - The day input
/// usableValves - List of valves can use (rate > 0)
/// int - minute remain
/// currentValueKey - Current valve we stand
///
/// Step 1:
/// - Check is `currentValueKey` valve have rate
///   - Set value = rate * remainMinutes
///
/// Step 2:
/// - Remove `currentValveKey` if have to `usableValves`
///
/// Step 3:
/// - DFS to all `usableValves` remain,
int _bestPath(Map<String, Valve> input, Set<String> usableValves, int remainMinutes, String currentValveKey) {
  if (remainMinutes == 0) return 0;

  int value = 0;

  final useableCurrentValve = usableValves.contains(currentValveKey);
  if (useableCurrentValve) {
    remainMinutes -= 1;
    value = input[currentValveKey]!.rate * remainMinutes;
  }

  int innerValue = 0;
  final currentValve = input[currentValveKey]!;
  final travelCosts = currentValve.costs;

  final linkedUsableValvesEntries = {...usableValves}..removeWhere((e) => e == currentValveKey);

  for (var linkedValveKey in linkedUsableValvesEntries) {
    final linkedValve = input[linkedValveKey]!;
    final travelCost = travelCosts[linkedValve.id]!;
    if (travelCost < remainMinutes) {
      int linkedValue = _bestPath(
        input,
        linkedUsableValvesEntries,
        remainMinutes - travelCost,
        linkedValve.id,
      );
      innerValue = max(innerValue, linkedValue);
    }
  }

  return value + innerValue;
}
