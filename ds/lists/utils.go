package lists

type List[T comparable] interface {
	Append(value ...T)
	Prepend(value ...T)
	Insert(index int, value ...T) error

	GetFirst() (T, error)
	Get(index int) (T, error)
	GetSlice(index int, length int) ([]T, error)
	GetLast() (T, error)

	RemoveFirst() error                      // incase the list is empty
	Remove(index int) error                  // incase the list is empty
	RemoveByValue(value T) error             // incase the list is empty
	RemoveSlice(index int, length int) error // incase the list is empty
	RemoveLast() error                       // incase the list is empty
}
