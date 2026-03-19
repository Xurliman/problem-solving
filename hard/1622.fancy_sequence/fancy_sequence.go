package fancy_sequence

type Fancy struct {
	sequence      []int
	mod, add, mul int
}

func Constructor() Fancy {
	sequence := make([]int, 0)
	return Fancy{
		sequence: sequence,
		mod:      1_000_000_007,
		mul:      1,
		add:      0,
	}
}

func (this *Fancy) Append(val int) {
	this.sequence = append(this.sequence, val)
}

func (this *Fancy) AddAll(inc int) {
	for i := 0; i < len(this.sequence); i++ {
		this.sequence[i] = (this.sequence[i] + inc) % this.mod
	}
}

func (this *Fancy) MultAll(m int) {
	for i := 0; i < len(this.sequence); i++ {
		this.sequence[i] = (this.sequence[i] * m) % this.mod
	}
}

func (this *Fancy) GetIndex(idx int) int {
	if idx >= len(this.sequence) {
		return -1
	}
	return this.sequence[idx] % this.mod
}

/**
 * Your Fancy object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Append(val);
 * obj.AddAll(inc);
 * obj.MultAll(m);
 * param_4 := obj.GetIndex(idx);
 */
