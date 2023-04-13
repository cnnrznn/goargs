package goargs

type Vals map[string][]string

func (v Vals) Contains(key string) bool {
	if _, ok := v[key]; ok {
		return true
	}
	return false
}
