package result

import "errors"

type Result[T any] struct {
	Ok  T
	Err error
}

type ResultTuple[T any] struct {
}

func Ok[T any](ok T) Result[T] {
	return Result[T]{Ok: ok}
}

func Err[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

func ErrStr[T any](err string) Result[T] {
	return Result[T]{Err: errors.New(err)}
}

func From[T any](result T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok[T](result)
}

func (r Result[T]) Unwrap() T {
	if r.Err != nil {
		panic(r.Err)
	}
	return r.Ok
}

func (r Result[T]) UnwrapOr(def T) T {
	if r.Err != nil {
		return def
	}
	return r.Ok
}

func (r Result[T]) UnwrapOrElse(f func() T) T {
	if r.Err != nil {
		return f()
	}
	return r.Ok
}

func (r Result[T]) IsOk() bool {
	return r.Err == nil
}

func (r Result[T]) IsErr() bool {
	return r.Err != nil
}

func (r Result[T]) Expect(message string) T {
	if r.Err != nil {
		panic("Expect: " + message + " " + r.Err.Error())
	}
	return r.Ok
}
