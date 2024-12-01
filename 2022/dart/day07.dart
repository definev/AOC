import 'dart:io';

enum ObjectType { directory, file }

class SystemObject {
  const SystemObject(
    this.root, {
    required this.name,
    required this.type,
    this.size,
  });

  final SystemObject? root;
  final String name;
  final ObjectType type;
  final int? size;

  String get path => '${root?.path ?? '/'}$name${type == ObjectType.directory ? '/' : ''}';

  @override
  String toString() => '$path';
}

abstract class Command {
  const Command();

  SystemObject? execute(SystemObject? root, {required List<SystemObject> objects});
}

class ChangeDirectoryCommand extends Command {
  const ChangeDirectoryCommand(this.name);

  final String name;

  @override
  SystemObject? execute(SystemObject? root, {required List<SystemObject> objects}) {
    if (name == '.') return root;
    if (name == '..') return root?.root;
    if (name == '/') return root;

    final currentObject = objects //
        .where((e) => e.root?.path == root?.path)
        .firstWhere((e) => e.name == name, orElse: () => throw Exception('Not found directory'));
    if (currentObject.type == ObjectType.file) {
      throw Exception('Can\'t go to file');
    }

    return currentObject;
  }
}

class CreateDirectoryCommand extends Command {
  const CreateDirectoryCommand(this.name);

  final String name;

  @override
  SystemObject? execute(SystemObject? root, {required List<SystemObject> objects}) {
    final dir = SystemObject(root, name: name, type: ObjectType.directory);
    objects.add(dir);
    return root;
  }
}

class CreateFileCommand extends Command {
  const CreateFileCommand(this.size, this.name);

  final int size;
  final String name;

  @override
  SystemObject? execute(SystemObject? root, {required List<SystemObject> objects}) {
    final dir = SystemObject(root, name: name, type: ObjectType.file, size: size);
    objects.add(dir);
    return root;
  }
}

class ListAllCommand extends Command {
  const ListAllCommand();

  @override
  SystemObject? execute(SystemObject? root, {required List<SystemObject> objects}) {
    final all = objects.where((e) => e.root?.path == root?.path);
    for (var line in all) {
      print(line.name);
    }
    return root;
  }
}

Command parseCommand(String input) {
  final raw = input.split(' ');
  if (raw[0] == '\$') {
    switch (raw[1]) {
      case 'cd':
        return ChangeDirectoryCommand(raw[2]);
      case 'ls':
        return ListAllCommand();
    }
  }
  switch (raw[0]) {
    case 'dir':
      return CreateDirectoryCommand(raw[1]);
    default:
      return CreateFileCommand(int.parse(raw[0]), raw[1]);
  }
}

List<Command> parseInput(String fileName) {
  final lines = File(fileName).readAsLinesSync();
  List<Command> results = [];

  for (var line in lines) {
    results.add(parseCommand(line));
  }

  return results;
}

void main() {
  final input = parseInput('../input/day07.txt');
  firstHalfProblem(input);
  lastHalfProblem(input);
}

int getObjectSize(List<SystemObject> objects, SystemObject? object) {
  int total = 0;
  if (object?.type == ObjectType.file) {
    total += object!.size!;
  } else {
    final directoryObjects = objects.where((ele) => ele.root?.path == object?.path);
    for (var object in directoryObjects) {
      total += getObjectSize(objects, object);
    }
  }

  return total;
}

void firstHalfProblem(List<Command> input) {
  SystemObject? currentContext = null;
  List<SystemObject> objects = [];

  for (var command in input) {
    currentContext = command.execute(currentContext, objects: objects);
  }

  final directories = objects.where((ele) => ele.type == ObjectType.directory);
  int total = 0;
  for (var directory in directories) {
    final directoryObjects = objects.where((ele) => ele.root?.path == directory.path).toList();
    for (var object in directoryObjects) {
      if (object.type == ObjectType.file) continue;
      final directoryTotal = getObjectSize(objects, object);
      if (directoryTotal <= 100000) {
        total += directoryTotal;
      }
    }
  }

  print(total);
}

void lastHalfProblem(List<Command> input) {
  SystemObject? currentContext = null;
  List<SystemObject> objects = [];

  for (var command in input) {
    currentContext = command.execute(currentContext, objects: objects);
  }

  final directories = objects.where((ele) => ele.type == ObjectType.directory);
  final needSpace = 30000000 - (70000000 - getObjectSize(objects, null));
  print('TOTAL : ${getObjectSize(objects, null)}');
  print(needSpace);
  int total = 30000000;
  for (var directory in directories) {
    final directoryObjects = objects.where((ele) => ele.root?.path == directory.path).toList();
    for (var object in directoryObjects) {
      if (object.type == ObjectType.file) continue;
      final directoryTotal = getObjectSize(objects, object);
      if (directoryTotal >= needSpace && directoryTotal < total) {
        total = directoryTotal;
      }
    }
  }

  print(total);
}
