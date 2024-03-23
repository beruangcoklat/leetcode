type RandomizedSet struct {
	m map[int]struct{}
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]struct{}),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = struct{}{}
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; ok {
		delete(this.m, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	arr := []int{}
	for val := range this.m {
		arr = append(arr, val)
	}
	return arr[rand.Intn(len(arr))]
}