package age

type Age int

func (a Age) IsAdult() bool {
	return a >= 18
}
