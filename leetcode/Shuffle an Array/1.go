type Solution struct {
	ori  []int
	nums []int
}

func Constructor(nums []int) Solution {
	rand.Seed(time.Now().UnixNano())

	ori := make([]int, len(nums))
	for i, v := range nums {
		ori[i] = v
	}
	return Solution{
		nums: nums,
		ori:  ori,
	}
}

func (this *Solution) Reset() []int {
	for i, v := range this.ori {
		this.nums[i] = v
	}
	return this.nums
}

func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.nums); i++ {
		j := rand.Intn(len(this.nums))
		temp := this.nums[i]
		this.nums[i] = this.nums[j]
		this.nums[j] = temp
	}
	return this.nums
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
