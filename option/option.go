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

func Some[T any](value T) Option[T] {
	return Option[T]{val: []T{value}}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) IsNone() bool {
	return o.val == nil || len(o.val) == 0
}

func (o Option[T]) IsSome() T {
	if o.IsNone() {
		panic("Can not get value of None")
	}
	return o.val[0]
}

func (o Option[T]) Match(some func(val T) any, none func() any) any {
	if o.IsNone() {
		return none()
	} else {
		return some(o.val[0])
	}
}
