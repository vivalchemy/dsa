#include <iostream>

// swaps a and b
void swap(int &a, int &b) {
  // here a and b do not contain addresses but the references to the original
  // paased values c doesn't have references like cpp so we use pointers there
  int temp = a;
  a = b;
  b = temp;
}

// uses modified bubble sort to bubble up the highest element to end
void modBubbleSort(int arr[], int n) {
  for (int i = 0; i < n - 1; i++) {
    int count = 0;
    for (int j = 0; j < n - i - 1; j++) {
      if (arr[j + 1] < arr[j]) {
        swap(arr[j + 1], arr[j]);
        count = 1;
      }
    }
    if (count == 0)
      break;
  }
}

// Takes user input for the given array
void getArray(int arr[], int n) {
  std::cout << "Enter the elements: ";
  for (int i = 0; i < n; ++i) {
    std::cin >> arr[i];
  }
}

// Prints the elements of the array to the standard out
void printArray(int arr[], int n) {
  for (int i = 0; i < n; ++i) {
    std::cout << arr[i] << " ";
  }
}

int main() {
  int n;
  std::cout << "Enter the number of elements: ";
  std::cin >> n;
  int arr[n];

  getArray(arr, n);

  std::cout << "\nUnsorted array: ";
  printArray(arr, n);

  modBubbleSort(arr, n);

  std::cout << "\nSorted array: ";
  printArray(arr, n);

  return 0;
}
