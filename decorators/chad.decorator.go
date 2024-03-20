package decorators

type Object[T comparable] func(T) (T, error)

func ChadDecorator[T comparable](fn Object[T]) Object[T] {
	return func(args T) (T, error) {
		data, err := fn(args)

		if err != nil {
			return data, err
		}

		return data, nil
	}
}
