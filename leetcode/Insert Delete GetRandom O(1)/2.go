type RandomizedSet struct {
	m     map[int]int
	list  []int
	count int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m:     make(map[int]int),
		list:  []int{},
		count: 0,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}

	if this.count >= len(this.list) {
		this.list = append(this.list, val)
	} else {
		this.list[this.count] = val
	}

	this.m[val] = this.count
	this.count++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.m[val]
	if !ok {
		return false
	}

	// remove element from the list by swapping it with last index, and update the index from the map
	this.list[idx] = this.list[this.count-1]
	this.m[this.list[idx]] = idx
	this.count--
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.list[rand.Intn(this.count)]
}
