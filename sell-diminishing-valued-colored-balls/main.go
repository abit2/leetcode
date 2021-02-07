package main

const mod = 1000000007

func valid(kv []int, m map[int]int, M, orders int) bool {
	for _, v := range kv {
		k, v := v, m[v]
		if k <= M {
			break
		}
		orders = (orders - (k-M)*v)
		if orders <= 0 {
			return true
		}
	}
	return false
}
func maxProfit(inv []int, orders int) int {
	m := make(map[int]int)
	mxV := math.MinInt64
	for _, v := range inv {
		m[v]++
		mxV = max(mxV, v)
	}

	var key []int
	for k, _ := range m {
		key = append(key, k)
	}

	sort.Slice(key, func(i, j int) bool {
		return key[i] > key[j]
	})

	L, R := 1, mxV
	for L <= R {
		M := (L + R) >> 1
		// fmt.Println("L ", L, " R ", R, " M ", M)
		if valid(key, m, M, orders) {
			L = M + 1
			continue
		}
		R = M - 1
	}
	// fmt.Println("L ", L)

	val := 0
	mod := int(1e9 + 7)
	var n, cnt int
	for _, v := range key {
		n, cnt = v, m[v]
		if n <= L {
			break
		}
		// fmt.Println("n ", n, " cnt ", cnt)
		orders = (orders - cnt*(n-L))
		val = (val%mod + ((n*n+n-L-L*L)/2*cnt)%mod) % mod
		// fmt.Println("order ", orders, " val ", val)
	}

	if orders > 0 {
		val = (val%(1e9+7) + (orders*L)%(1e9+7)) % (1e9 + 7)
	}

	return val
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
