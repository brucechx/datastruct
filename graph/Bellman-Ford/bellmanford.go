package Bellman_Ford

type BellmanGraph interface {
	NodeCount() int
	GetInbound(v int) map[int]int
}

func BellmanFord(g BellmanGraph, s int) (dists map[int]int) {
	L := map[int]map[int]int{}
	setDist := func(i, v, c int) {
		if _, ok := L[i]; !ok {
			L[i] = map[int]int{}
		}
		L[i][v] = c
	}
	getDist := func(i, v int) (c int, ok bool) {
		if _, ok := L[i]; !ok {
			return c, false
		}

		c, ok = L[i][v]
		return c, ok
	}

	// base case: hops = 0
	setDist(0, s, 0)

	for i := 1; i < g.NodeCount(); i++ {
		for v := 1; v <= g.NodeCount(); v++ {
			case1, ok1 := getDist(i-1, v)

			case2 := 0
			ok2 := false
			for u, w := range g.GetInbound(v) {
				if wCost, ok := getDist(i-1, u); ok {
					vCost := wCost + w
					if !ok2 {
						ok2 = true
						case2 = vCost
					} else if vCost < case2 {
						case2 = vCost
					}
				}
			}

			if !ok1 && !ok2 {
				continue
			} else if !ok1 {
				setDist(i, v, case2)
			} else if !ok2 {
				setDist(i, v, case1)
			} else if case1 < case2 {
				setDist(i, v, case1)
			} else {
				setDist(i, v, case2)
			}
		}
	}

	dists, ok := L[g.NodeCount()-1]
	if !ok {
		return map[int]int{}
	}

	return dists
}
