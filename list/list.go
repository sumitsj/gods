package list

//type List[T any] interface {
//	First() *T
//	Add(item T)
//	Find(predicate func(T) bool) *T
//}

type list[T any] struct {
	Value []T
}

func (l *list[T]) First() *T {
	if len(l.Value) > 0 {
		return &l.Value[0]
	}

	return nil
}

func (l *list[T]) Add(item T) {
	l.Value = append(l.Value, item)
}

func (l *list[T]) Find(predicate func(T) bool) *T {
	for _, item := range l.Value {
		if predicate(item) {
			return &item
		}
	}
	return nil
}

func NewList[T any]() list[T] {
	return list[T]{
		Value: []T{},
	}
}