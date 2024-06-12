package isok

type FuncSignature[T any] func(a ...interface{}) (T, error)

type Zero struct {
}

type TwoResult[F any, S any] struct {
	f F
	s S
}

type ThreeResult[F, S, T any] struct {
	f F
	s S
	t T
}

type FourResult[F, S, T, FO any] struct {
	f  F
	s  S
	t  T
	fo FO
}

func (t *FourResult[F, S, T, FO]) GetSecond() S {
	return t.s
}

func (t *FourResult[F, S, T, FO]) GetFirst() F {
	return t.f
}

func (t *FourResult[F, S, T, FO]) GetThird() T {
	return t.t
}

func (t *FourResult[F, S, T, FO]) GetFourth() FO {
	return t.fo
}

func (t *TwoResult[F, S]) GetFirst() F {
	return t.f
}

func (t *TwoResult[F, S]) GetSecond() S {
	return t.s
}

func (t *ThreeResult[F, S, T]) GetFirst() S {
	return t.s
}

func (t *ThreeResult[F, S, T]) GetSecond() T {
	return t.t
}

func (t *ThreeResult[F, S, T]) GetThird() F {
	return t.f
}

func ResultFunc[T any](function FuncSignature[T], props ...interface{}) Result[T] {
	result, err := function(props...)

	if err != nil {
		return NewErr[T](err)
	}

	return NewOk(result)
}

func ResultFun1[T any](result T, err error) Result[T] {
	if err != nil {
		return NewErr[T](err)
	}

	return NewOk(result)
}

func ResultFun0(err error) Result[int] {
	if err != nil {
		return NewErr[int](err)
	}

	return NewOk[int](0)
}

func ResultFunc2[F any, S any](r1 F, r2 S, err error) Result[TwoResult[F, S]] {
	if err != nil {
		return NewErr[TwoResult[F, S]](err)
	}

	var a = TwoResult[F, S]{
		f: r1,
		s: r2,
	}

	return NewOk[TwoResult[F, S]](a)
}

func ResultFunc3[F, S, T any](r1 F, r2 S, r3 T, err error) Result[ThreeResult[F, S, T]] {
	if err != nil {
		return NewErr[ThreeResult[F, S, T]](err)
	}

	var a = ThreeResult[F, S, T]{
		f: r1,
		s: r2,
		t: r3,
	}

	return NewOk[ThreeResult[F, S, T]](a)
}

func ResultFunc4[F, S, T, FO any](r1 F, r2 S, r3 T, r4 FO, err error) Result[FourResult[F, S, T, FO]] {
	if err != nil {
		return NewErr[FourResult[F, S, T, FO]](err)
	}

	var a = FourResult[F, S, T, FO]{
		f:  r1,
		s:  r2,
		t:  r3,
		fo: r4,
	}

	return NewOk[FourResult[F, S, T, FO]](a)
}

func NewOk[T any](v T) Result[T] {
	return Result[T]{ok: &v}
}

func NewErr[T any](e error) Result[T] {
	return Result[T]{err: e}
}

type Result[T any] struct {
	ok  *T
	err error
}

func (r *Result[T]) IsErr() bool {
	return r.err != nil
}

func (r *Result[T]) IsOk() bool {
	return r.ok != nil && r.err == nil
}

func (r *Result[T]) Unwrap() T {
	if r.err != nil {
		panic(r.err)
	}

	return *r.ok
}

func (r *Result[T]) UnwrapErr() error {
	if r.ok != nil {
		panic("unable to unwrap nil result")
	}

	return r.err
}

// func melon() (int, int, error) {
// 	return 1, 1, nil
// }
