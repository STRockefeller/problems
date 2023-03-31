package fp

func Curry[T1, T2, R any](f func(T1, T2) R) func(T1) func(T2) R {
	return func(t1 T1) func(T2) R {
		return func(t2 T2) R {
			return f(t1, t2)
		}
	}
}
