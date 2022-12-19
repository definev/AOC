import 'dart:io';
import 'dart:math';

import 'helper/algo.dart';
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
  final input = parseInput('../sample/day16.txt');

  // firstHalfProblem(input);
  lastHalfProblem(input);
}

class Node {
  final String id;
  final int rate;
  final Map<String, Tuple<int, int>> weights;

  Node(this.id, this.rate, this.weights);

  bool operator ==(Object? other) => (other is Node) ? other.id == id : false;

  @override
  int get hashCode => id.hashCode ^ rate;

  @override
  String toString() {
    return '$id | $rate | $weights';
  }
}

class Pair {
  Pair(this.first, this.last);

  final String first;
  final String last;

  operator ==(Object? other) => (other is Pair) ? other.first == first && other.last == last : false;

  @override
  int get hashCode => first.hashCode ^ last.hashCode;
}

void lastHalfProblem(Map<String, Valve> input) {
  Set<Node> nodes = {};
  for (var valve in input.entries) {
    if (valve.value.rate == 0) continue;
    final weightsEntries = valve.value.costs
        .map((key, value) {
          if (input[key]!.rate == 0) return MapEntry(key, Tuple(0, 0));
          return MapEntry(key, Tuple(input[key]!.rate, value));
        })
        .entries
        .toList()
      ..removeWhere((e) => e.key == valve.key || e.value.first == 0);
    selectionSort(
      weightsEntries,
      compareTo: (first, last) => last.value.first.compareTo(first.value.first),
    );

    nodes.add(
      Node(
        valve.value.id,
        valve.value.rate,
        Map.fromEntries(weightsEntries),
      ),
    );
  }

  final weight = getWeightConnect(input, 'AA', null, 1);
  Node start = Node('AA', 0, weight);
  nodes.add(start);

  int value = _findBestPathWithElephant(
    nodes,
    needTravered: nodes.map((e) => e.id).toSet()..removeWhere((e) => e == 'AA'),
    current: start,
    elephant: start,
    exceptCurrent: null,
    exceptElephant: null,
    minutes: 0,
    total: 0,
    triggerCurrent: 0,
    triggerElephant: 0,
  );
  print(value);
}

int _findBestPathWithElephant(
  Set<Node> nodes, {
  required Set<String> needTravered,
  required Node current,
  required String? exceptCurrent,
  required Node elephant,
  required String? exceptElephant,
  required int minutes,
  required int total,
  required int triggerCurrent,
  required int triggerElephant,
}) {
  final nextMinutes = minutes + 1;

  int nextTotal = total;

  if (nextMinutes == 26) return total;
  if (nextMinutes == triggerCurrent) {
    nextTotal += (26 - nextMinutes) * current.rate;
  }
  if (nextMinutes == triggerElephant) {
    nextTotal += (26 - nextMinutes) * elephant.rate;
  }
  if (needTravered.isEmpty) return nextTotal;

  int bestTotal = nextTotal;

  Set<Pair> pairs = {};

  for (var currentRelate in current.weights.keys) {
    for (var elephantRelate in elephant.weights.keys) {
      bool update = true;
      for (var pair in pairs) {
        if (pair.first == currentRelate && pair.last == elephantRelate) {
          update = false;
          break;
        }
        if (pair.first == elephantRelate && pair.last == currentRelate) {
          update = false;
          break;
        }
      }
      if (!update) continue;
      if (currentRelate == elephantRelate) continue;
      if (currentRelate == exceptCurrent ||
          currentRelate == exceptElephant ||
          elephantRelate == exceptCurrent ||
          elephantRelate == exceptElephant) continue;
      if (!needTravered.contains(currentRelate)) continue;
      if (!needTravered.contains(elephantRelate)) continue;
      pairs.add(Pair(currentRelate, elephantRelate));
    }
  }

  Set<String> triedStep = {};

  for (var pair in pairs) {
    final currentRelate = pair.first;
    final elephantRelate = pair.last;
    if (currentRelate == elephantRelate) continue;
    final stillCurrent = triggerCurrent >= nextMinutes || !needTravered.contains(currentRelate);
    final currentRelateNode = stillCurrent
        ? current
        : () {
            return nodes.firstWhere((e) => e.id == currentRelate);
          }();
    final nextTriggerCurrent = stillCurrent
        ? triggerCurrent
        : () {
            return triggerCurrent + current.weights[currentRelate]!.second + 1;
          }();

    final stillElephant = triggerElephant >= nextMinutes || !needTravered.contains(elephantRelate);
    final elephantRelateNode = stillElephant
        ? elephant
        : () {
            return nodes.firstWhere((e) => e.id == elephantRelate);
          }();
    final nextTriggerElephant = stillElephant
        ? triggerElephant
        : () {
            return triggerElephant + elephant.weights[elephantRelate]!.second + 1;
          }();

    final step = 'C:${current.id}->${currentRelateNode.id}|E:${elephant.id}->${elephantRelateNode.id}';
    if (triedStep.contains(step)) continue;
    triedStep.add(step);
    print(triedStep);

    final newTotal = _findBestPathWithElephant(
      nodes,
      needTravered: {...needTravered}..removeAll([currentRelateNode.id, elephantRelateNode.id]),
      current: currentRelateNode,
      exceptCurrent: currentRelateNode.id,
      elephant: elephantRelateNode,
      exceptElephant: elephantRelateNode.id,
      minutes: nextMinutes,
      total: nextTotal,
      triggerCurrent: nextTriggerCurrent,
      triggerElephant: nextTriggerElephant,
    );

    if (bestTotal >= newTotal) continue;
    bestTotal = max(bestTotal, newTotal);
    print('CURRENT BEST : $bestTotal');
  }

  return bestTotal;
}

Map<String, Tuple<int, int>> getWeightConnect(
  Map<String, Valve> input,
  String currentNode,
  String? except,
  int depth,
) {
  Map<String, Tuple<int, int>> value = {};
  final valve = input[currentNode]!;
  for (var linked in valve.linkedValves) {
    if (linked == except) continue;
    final linkedNode = input[linked]!;
    if (linkedNode.rate == 0) {
      value = {...value, ...getWeightConnect(input, linked, currentNode, depth + 1)};
      continue;
    }
    value[linked] = Tuple(input[linked]!.rate, depth);
  }

  return value;
}

void firstHalfProblem(Map<String, Valve> input) {
  final usableValves = Map.fromEntries(input.entries.where((e) => e.value.rate > 0)).keys.toList();
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
int _bestPath(Map<String, Valve> input, List<String> usableValves, int remainMinutes, String currentValveKey) {
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

  final linkedUsableValvesEntries = List<String> //
      .from(usableValves)
    ..removeWhere((e) => e == currentValveKey);

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
