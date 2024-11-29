package queryable

import (
	"github.com/titosilva/pdpr-go/internal/collections/iterator"
	"github.com/titosilva/pdpr-go/internal/maybe"
)

type Predicate[T any] func(T) bool

type Queryable[T any] interface {
	iterator.Iterable[T]
	All(pred Predicate[T]) bool
	Any(pred Predicate[T]) bool
	Where(pred Predicate[T]) Queryable[T]
	First() maybe.Maybe[T]

	Skip(offset int) Queryable[T]
	Take(count int) Queryable[T]
}
