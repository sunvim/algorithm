package ds

type Holder interface {
	At(i int) interface{}
	Insert(i int, data interface{}) error
	InsertFront(data interface{}) error
	InsertEnd(data interface{}) error
	Remove(i int) error
	RemoveFront() error
	RemoveEnd() error
	Next() bool
	Data() interface{}
}
