package goargs

type Vals map[string][]string

func (v Vals) Contains(key string) bool {
	_, ok := v[key]
	return ok
}
