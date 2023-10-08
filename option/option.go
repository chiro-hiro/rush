package option

type Option[T any] struct {
	val []T
}

func From[T any](value T, err error) Option[T] {
	if err == nil {
		return Option[T]{val: []T{value}}
	}
	return Option[T]{}
}

func (o Option[T]) None() bool {
	return o.val == nil || len(o.val) == 0
}

func (o Option[T]) Some() T {
	if o.None() {
		panic("Can not get value of None")
	}
	return o.val[0]
}

func (o Option[T]) Match(some func(val T) any, none func() any) any {
	if o.None() {
		return none()
	} else {
		return some(o.Some())
	}
}
