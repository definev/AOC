class Tuple<F, S> {
  const Tuple(this.first, this.second);

  final F first;
  final S second;

  @override
  String toString() {
    return 'FIRST : $first\nSECOND: $second';
  }
}

class Triple<F, S, T> {
  const Triple(this.first, this.second, this.third);

  final F first;
  final S second;
  final T third;

  @override
  String toString() {
    return 'FIRST : $first\nSECOND: $second\nTRIPLE: $third';
  }
}
