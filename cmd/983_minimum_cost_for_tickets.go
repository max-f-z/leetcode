package main

func mincostTickets(days []int, costs []int) int {
	max := days[len(days)-1]
	dict := make(map[int]bool)
	mem := make([]int, max+1)
	for i := 0; i <= max; i++ {
		mem[i] = -1
	}

	for _, v := range days {
		dict[v] = true
	}

	m := &mct{
		max:   max,
		costs: costs,
		dict:  dict,
		mem:   mem,
	}

	return m.mincostTicketsHelper(1)
}

type mct struct {
	max   int
	costs []int
	dict  map[int]bool
	mem   []int
}

func (m *mct) mincostTicketsHelper(i int) int {
	if i > m.max {
		return 0
	}

	if m.mem[i] != -1 {
		return m.mem[i]
	}
	var min int
	if m.dict[i] {
		min = m.mincostTicketsHelper(i+1) + m.costs[0]
		tmp7 := m.mincostTicketsHelper(i+7) + m.costs[1]
		tmp30 := m.mincostTicketsHelper(i+30) + m.costs[2]
		if tmp7 < min {
			min = tmp7
		}
		if tmp30 < min {
			min = tmp30
		}
	} else {
		min = m.mincostTicketsHelper(i + 1)
	}

	m.mem[i] = min

	return min
}
