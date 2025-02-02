#include <iostream>

// swaps a and b
void swap(int &a, int &b) {
  // here a and b do not contain addresses but the references to the original
  // paased values c doesn't have references like cpp so we use pointers there
  int temp = a;
  a = b;
  b = temp;
}

// Insertion sort to insert elements in their correct positions
void insertionSort(int arr[], int n) {
  for (int i = 1; i < n; ++i) {
    int key = arr[i];
    int j = i - 1;

    // Move elements greater than key to one position ahead
    while (j >= 0 && arr[j] > key) {
      arr[j + 1] = arr[j];
      --j;
    }

    // Insert key in its correct position
    arr[j + 1] = key;
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

  insertionSort(arr, n);

  std::cout << "\nSorted array: ";
  printArray(arr, n);

  return 0;
}
