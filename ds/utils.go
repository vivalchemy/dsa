package ds

import "log"

type DataStructure interface {
	Empty() bool
	Length() int
	Clear()
	Print() error          // incase the ds is empty
	Values() []interface{} // incase the ds is empty
	String() string
}

func IsEmpty(DS DataStructure) bool {
	return DS.Empty()
}

func Len(DS DataStructure) int {
	return DS.Length()
}

func Clear(DS DataStructure) {
	DS.Clear()
}

func Print(DS DataStructure) {
	err := DS.Print()
	if err != nil {
		log.Println(err)
	}
}

func Values(DS DataStructure) []interface{} {
	return DS.Values()
}

func String(DS DataStructure) string {
	return DS.String()
}
