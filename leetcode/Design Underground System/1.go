type UndergroundSystem struct {
	distMap map[string]*distData
	userMap map[int]data
}

type data struct {
	stationName string
	t           int
}

type distData struct {
	total int
	count int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		distMap: make(map[string]*distData),
		userMap: make(map[int]data),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.userMap[id] = data{
		stationName: stationName,
		t:           t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	d := this.userMap[id]
	src := d.stationName
	dst := stationName
	delete(this.userMap, id)

	key := src + "-" + dst

	dist, ok := this.distMap[key]
	if !ok {
		this.distMap[key] = &distData{
			total: t - d.t,
			count: 1,
		}
	} else {
		dist.count++
		dist.total += (t - d.t)
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	key := startStation + "-" + endStation
	dist := this.distMap[key]
	return float64(dist.total) / float64(dist.count)
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
