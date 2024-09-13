package timeBasedKeyValueStore

type ValStamp struct {
	Val  string
	Time int
}

type TimeMap struct {
	store map[string][]ValStamp
}

func Constructor() TimeMap {
	return TimeMap{store: make(map[string][]ValStamp)}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := t.store[key]; !ok {
		t.store[key] = make([]ValStamp, 0)
	}
	t.store[key] = append(t.store[key], ValStamp{value, timestamp})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	var res string
	var values []ValStamp
	if _, ok := t.store[key]; ok {
		values = t.store[key]
	}

	left, right := 0, len(values)-1
	for left <= right {
		mid := (left + right) / 2
		if values[mid].Time <= timestamp {
			res = values[mid].Val
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
