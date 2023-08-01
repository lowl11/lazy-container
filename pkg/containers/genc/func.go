package genc

type EachFunc[T any] func(item T)
type WhereFunc[T any] func(item T) bool
type CountFunc[T any] func(item T) bool
type ConditionFunc[T any] func(item T) bool
type SortFunc[T any] func(i, j int) bool
type GetFunc[T any] func(item T) bool
