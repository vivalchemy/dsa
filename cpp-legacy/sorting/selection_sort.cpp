#include <iostream>

// swaps a and b
void swap(int &a, int &b) {
  // here a and b do not contain addresses but the references to the original
  // paased values c doesn't have references like cpp so we use pointers there
  int temp = a;
  a = b;
  b = temp;
}

// Selection sort to select the smallest element and place it at the beginning
void selectionSort(int arr[], int n) {
  for (int i = 0; i < n - 1; ++i) {
    int minIndex = i;
    for (int j = i + 1; j < n; ++j) {
      if (arr[j] < arr[minIndex]) {
        minIndex = j;
      }
    }
    swap(arr[i], arr[minIndex]);
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

  selectionSort(arr, n);

  std::cout << "\nSorted array: ";
  printArray(arr, n);

  return 0;
}
