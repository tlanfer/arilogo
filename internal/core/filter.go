package core

type Predicate[A any] func(A) bool

func Filter[A any](in []A, match Predicate[A]) (out []A) {
	for _, a := range in {
		if match(a) {
			out = append(out, a)
		}
	}
	return out
}
