#include <iostream>

// swaps a and b
void swap(int &a, int &b) {
  // here a and b do not contain addresses but the references to the original paased values c doesn't have references like cpp so we use pointers there
  int temp = a;
  a = b;
  b = temp;
}

// uses bubble sort to bubble up the lowest element to start
void bubbleSort(int arr[], int n) {
  for (int i = 0; i < n - 1; ++i) { //n
    for (int j = i + 1; j < n; ++j) {
      if (arr[j] < arr[i]) {
        swap(arr[j], arr[i]);
      }
    }
  }
}

// Takes user input for the given array
void getArray(int arr[], int n){
  std::cout << "Enter the elements: ";
  for (int i = 0; i < n; ++i) {
    std::cin >> arr[i];
  }
}

// Prints the elements of the array to the standard out
void printArray(int arr[], int n){
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

  bubbleSort(arr, n);

  std::cout << "\nSorted array: ";
  printArray(arr, n);

  return 0;
}
