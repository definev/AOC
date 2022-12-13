int binarySearch(List<int> arr, int value) {
  int left = 0;
  int right = arr.length;

  while (left <= right) {
    int mid = (right - left) ~/ 2 + left;
    if (arr[mid] == value) return mid;

    if (arr[mid] > value) right = mid - 1;
    if (arr[mid] < value) left = mid + 1;
  }

  return -1;
}

List<T> selectionSort<T>(List<T> arr, {required int Function(T first, T last) compareTo}) {
  for (var i = 0; i < arr.length - 1; i++) {
    int min = i;
    for (var j = i + 1; j < arr.length; j++) {
      if (compareTo(arr[j], arr[min]) == -1) min = j;
    }
    var temp = arr[i];
    arr[i] = arr[min];
    arr[min] = temp;
  }

  return arr;
}

List<int> insertionSort(List<int> arr) {
  int i, j, last;
  for (i = 1; i < arr.length; i++) {
    last = arr[i];
    j = i;
    while (j > 0 && arr[j - 1] > last) {
      arr[j] = arr[j - 1];
      j = j - 1;
      arr[j] = last;
    }
  }

  return arr;
}
