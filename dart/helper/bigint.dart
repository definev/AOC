const BASE = 10;

class BigInt implements Comparable<BigInt> {
  BigInt(this.number);

  List<int> number;

  static BigInt? tryParse(String raw) {
    if (int.tryParse(raw[0]) == null) return null;
    if (raw.isEmpty) throw Exception('BigInt cannot be a blank');

    List<int> number = [];

    for (var char in raw.codeUnits) {
      if (char < '0'.codeUnitAt(0) || char > '9'.codeUnitAt(0)) {
        throw Exception('Not a number');
      } else {
        number.add(char - '0'.codeUnitAt(0));
      }
    }

    return BigInt(number.reversed.toList());
  }

  static BigInt intParse(int data) => tryParse(data.toString())!;

  static List<int> _expand(List<int> number, int length) => [...number, for (int i = 0; i < length; i++) 0];

  static void equalLength(BigInt a, BigInt b) {
    if (a.number.length < b.number.length) a.number = _expand(a.number, b.number.length - a.number.length);
    if (a.number.length > b.number.length) b.number = _expand(b.number, a.number.length - b.number.length);
  }

  @override
  bool operator ==(Object? other) => other is BigInt ? number.join() == other.number.join() : false;

  void removePrecision() {
    while (number.length > 1 && number.last == 0) number.removeLast();
  }

  @override
  String toString() => number.reversed.join();

  BigInt operator +(BigInt other) {
    equalLength(this, other);

    int carry = 0;
    BigInt res = BigInt([]);
    for (int i = 0; i < number.length; ++i) {
      int d = number[i] + other.number[i] + carry;
      carry = d ~/ 10;
      res.number.add(d % 10);
    }

    if (carry != 0) res.number.add(1);

    other.removePrecision();
    return res;
  }

  BigInt operator -(BigInt other) {
    equalLength(this, other);

    int d = 0, carry = 0;
    BigInt res = BigInt([]);
    for (int i = 0; i < number.length; i++) {
      d = number[i] - other.number[i] - carry;

      if (d < 0) {
        d += 10;
        carry = 1;
      } else
        carry = 0;

      res.number.add(d);
    }

    while (res.number.length > 1 && res.number.last == 0) res.number.removeLast();

    other.removePrecision();
    return res;
  }

  BigInt operator *(BigInt other) {
    List<int> res = List.generate(number.length + other.number.length + 1, (index) => 0);
    for (int i = 0; i < number.length; ++i) {
      for (int j = 0; j < other.number.length; ++j) {
        res[i + j] += number[i] * other.number[j];
        res[i + j + 1] += res[i + j] ~/ BASE;
        res[i + j] %= BASE;
      }
    }
    return BigInt(res)..removePrecision();
  }

  int operator %(int x) {
    int r = 0;
    for (int i = number.length - 1; i >= 0; --i) {
      r = (r * BASE + number[i]) % x;
    }
    return r;
  }

  @override
  int compareTo(BigInt other) {
    if (number.length > other.number.length) return 1;
    if (number.length < other.number.length) return -1;
    for (var i = number.length - 1; i >= 0; i--) {
      if (number[i] > other.number[i]) return 1;
      if (number[i] < other.number[i]) return -1;
    }

    return 0;
  }
}
