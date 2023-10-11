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

func (o Option[T]) IsSome() bool {
	return !o.IsNone()
}

func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("unwrap a none option")
	}
	return o.val[0]
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.IsNone() {
		return def
	}
	return o.val[0]
}

func (o Option[T]) UnwrapOrElse(fn func() T) T {
	if o.IsNone() {
		return fn()
	}
	return o.val[0]
}

func (o Option[T]) Expect(msg string) T {
	if o.IsNone() {
		panic(msg)
	}
	return o.val[0]
}

type SomeResolve[T any] func(val T) interface{}

type NoneResolve[T any] func() interface{}

func (o Option[T]) Match(some SomeResolve[T], none NoneResolve[T]) interface{} {
	if o.IsNone() {
		return none()
	} else {
		return some(o.val[0])
	}
}
