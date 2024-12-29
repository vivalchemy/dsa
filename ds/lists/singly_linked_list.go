package lists

import (
	"errors"
	"fmt"
)

var SLL_EMPTY = errors.New("The linked list is empty")
var SLL_INDEX_OUT_OF_RANGE = errors.New("The index is out of range")
var SLL_LENGTH_NOT_GREATER_THAN_ZERO = errors.New("The length must be greater than 0")
var SLL_VALUE_NOT_FOUND = errors.New("The value not found in the linked list")
var SLL_INDEX_OUT_OF_RANGE_N = errors.New("The index + n is out of range")
var SLL_INDEX_OUT_OF_RANGE_SIZE = errors.New("The index is out of range")

// - [x] Empty() bool
// - [x] Length() int
// - [x] Clear()
// - [x] Print() error          // incase the ds is empty
// - [x] Values() []interface{} // incase the ds is empty
// - [x] String() string

// List[T comparable]
// - [x] Append(value ...T)
//- [x]  Prepend(value ...T)
// - [x] Insert(index int, value ...T) error
//
// - [x] GetFirst() T
// - [x] Get(index int) (T, error)
// - [x] GetSlice(index int, length int) ([]T, error)
// - [x] GetLast() T
//
// - [ ] RemoveFirst() error                      // incase the list is empty
// - [ ] Remove(index int) error                  // incase the list is empty
// - [ ] RemoveByValue(value T) error             // incase the list is empty
// - [ ] RemoveSlice(index int, length int) error // incase the list is empty
// - [ ] RemoveLast() error                       // incase the list is empty

// singlyLinkedListNode represents a single element in the linked list.
// It holds a value of type T and a pointer to the next singlyLinkedListNode in the list.
type singlyLinkedListNode[T comparable] struct {
	value T                        // The value stored in the node
	next  *singlyLinkedListNode[T] // Pointer to the next node in the linked list
}

// linkedList represents the linked list data structure.
// It holds a pointer to the head node of the list.
type singlyLinkedList[T comparable] struct {
	head *singlyLinkedListNode[T] // Pointer to the first node in the linked list
	size int
}

// Creates a new singly linked list.
// It returns a pointer to the newly created linked list instance.
func NewLinkedList[T comparable]() *singlyLinkedList[T] {
	return &singlyLinkedList[T]{} // Return a pointer to the newly created linked list with no nodes
}

// DataStructure interface

func (sll *singlyLinkedList[T]) Empty() bool {
	return sll.head == nil
}

func (sll *singlyLinkedList[T]) Length() int {
	return sll.size
}

func (sll *singlyLinkedList[T]) Clear() {
	sll.head = nil
	sll.size = 0
}

func (sll *singlyLinkedList[T]) Print() error {
	if sll.head == nil {
		return SLL_EMPTY
	}

	fmt.Println("Linked List:")
	fmt.Print("Size: ", sll.Length(), "\n Elements: ")

	current := sll.head
	fmt.Println(current.value, "->")
	for current.next != nil {
		current = current.next
		fmt.Print(current.value)
	}
	return nil
}

func (sll *singlyLinkedList[T]) Values() []T {
	if sll.head == nil {
		return []T{}
	}
	values := []T{}
	current := sll.head
	for current.next != nil {
		values = append(values, current.value)
		current = current.next
	}
	return values
}

func (sll *singlyLinkedList[T]) String() string {
	str := ""
	if sll.head == nil {
		return str
	}

	current := sll.head
	for current.next != nil {
		str += fmt.Sprintf("%v->", current.value)
		current = current.next
	}
	str += fmt.Sprintf("%v", current.value)
	return str
}

// List interface

func (sll *singlyLinkedList[T]) Append(value ...T) {
	length := len(value)
	if length == 0 {
		return
	}

	// create a linked list with the newNode pointing to the head
	newNode := tempListWithoutHead(nil, value...)
	sll.size += length

	if sll.head == nil {
		sll.head = newNode
		return
	}

	current := sll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (sll *singlyLinkedList[T]) Prepend(values ...T) {
	length := len(values)
	if length == 0 {
		return
	}
	// create a linked list with the returnNode pointing to the first nodei in the values
	sll.head = tempListWithoutHead(sll.head, values...)
	sll.size += length
}

func (sll *singlyLinkedList[T]) Insert(index int, value ...T) error {
	length := len(value)
	if length == 0 {
		return SLL_EMPTY
	}

	if index < 0 || index > sll.size {
		return SLL_INDEX_OUT_OF_RANGE
	}

	// create a linked list with the newNode pointing to the head

	current := sll.head
	for current.next != nil && index > 1 {
		current = current.next
		index--
	}

	current.next = tempListWithoutHead(current.next, value...)
	sll.size += length

	return nil
}

// - [ ] GetFirst() T
// - [ ] Get(index int) (T, error)
// - [ ] GetLast() T
func (sll *singlyLinkedList[T]) GetFirst() (T, error) {
	var zeroValue T
	if sll.head == nil {
		return zeroValue, SLL_EMPTY
	}
	return sll.head.value, nil
}

func (sll *singlyLinkedList[T]) Get(index int) (T, error) {
	var zeroValue T
	if sll.head == nil {
		return zeroValue, SLL_EMPTY
	}

	if index < 0 || index >= sll.size {
		return zeroValue, SLL_INDEX_OUT_OF_RANGE
	}

	current := sll.head
	for current.next != nil && index > 0 {
		current = current.next
		index--
	}

	return current.value, nil
}

func (sll *singlyLinkedList[T]) GetSlice(index int, length int) ([]T, error) {
	if index < 0 || index >= sll.size {
		return nil, SLL_EMPTY
	}

	if length <= 0 {
		return nil, SLL_LENGTH_NOT_GREATER_THAN_ZERO
	}
	// Create a slice to hold the elements
	result := make([]T, 0, length)

	current := sll.head

	// Traverse to the specified index
	for i := 0; i < index && current != nil; i++ {
		current = current.next
	}

	// Collect up to n elements from the current position
	for current != nil && length > 0 {
		result = append(result, current.value)
		current = current.next
		length--
	}

	return result, nil
}

func (sll *singlyLinkedList[T]) GetLast() (T, error) {
	var zeroValue T
	if sll.head == nil {
		return zeroValue, SLL_EMPTY
	}

	current := sll.head
	for current.next != nil {
		current = current.next
	}
	return current.value, nil
}

// - [ ] RemoveFirst() error                      // incase the list is empty
func (sll *singlyLinkedList[T]) RemoveFirst() error {
	if sll.head == nil {
		return SLL_EMPTY
	}
	sll.head = sll.head.next
	sll.size--
	return nil
}

// - [ ] Remove(index int) error                  // incase the list is empty
// - [ ] RemoveFirstByValue(value T) error             // incase the list is empty
func (sll *singlyLinkedList[T]) RemoveFirstByValue(value T) error {
	if sll.head == nil {
		return SLL_EMPTY
	}

	// Check if the head node has the value to remove
	if sll.head.value == value {
		sll.head = sll.head.next
		sll.size--
		return nil
	}

	// Traverse the list to find the value
	current := sll.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			sll.size--
			return nil
		}
		current = current.next
	}

	// Value not found
	return SLL_VALUE_NOT_FOUND
}

// - [ ] RemoveByValue(value T) error             // incase the list is empty
func (sll *singlyLinkedList[T]) RemoveByValue(value T) error {
	if sll.head == nil {
		return SLL_EMPTY
	}

	// Handle the case where the head needs to be removed
	for sll.head != nil && sll.head.value == value {
		sll.head = sll.head.next
		sll.size--
	}

	current := sll.head

	// Traverse the list to find and remove occurrences of the value
	for current.next != nil {
		// Check if the next node's value matches the value to remove. Hence it will remove the last node it is matches too
		if current.next.value == value {
			current.next = current.next.next
			sll.size--
		} else {
			current = current.next
		}
	}

	return nil

}

// - [ ] RemoveSlice(index int, length int) error // incase the list is empty
func (sll *singlyLinkedList[T]) RemoveSlice(index int, length int) error {
	if sll.head == nil {
		return SLL_EMPTY
	}

	if index < 0 || index >= sll.size {
		return SLL_INDEX_OUT_OF_RANGE
	}

	if length <= 0 {
		return SLL_LENGTH_NOT_GREATER_THAN_ZERO
	}

	if index+length > sll.size {
		return SLL_INDEX_OUT_OF_RANGE_SIZE
	}

	if index == 0 {
		for i := 0; i < length && sll.head != nil; i++ {
			sll.head = sll.head.next
			sll.size--
		}
		return nil
	}

	current := sll.head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.next
	}

	toRemove := current.next
	for i := 0; i < length && toRemove != nil; i++ {
		toRemove = toRemove.next
		sll.size--
	}

	current.next = toRemove

	return nil
}

// - [ ] RemoveLast() error                       // incase the list is empty
func (sll *singlyLinkedList[T]) RemoveLast() error {
	if sll.head == nil {
		return SLL_EMPTY
	}

	if sll.head.next == nil {
		sll.head = nil
		return nil
	}

	current := sll.head
	for current.next.next != nil {
		current = current.next
	}

	current.next = nil
	return nil
}

// -----------------------------------------------------------------------------
// Helper functions
// -----------------------------------------------------------------------------

// create a linked list with the newNode pointing to the head
func tempListWithoutHead[T comparable](finalPointer *singlyLinkedListNode[T], values ...T) *singlyLinkedListNode[T] {
	lengthOfValues := len(values)
	newNode := &singlyLinkedListNode[T]{value: values[lengthOfValues-1], next: finalPointer}
	for i := lengthOfValues - 2; i >= 0; i-- {
		newNode = &singlyLinkedListNode[T]{value: values[i], next: newNode}
	}
	return newNode
}
