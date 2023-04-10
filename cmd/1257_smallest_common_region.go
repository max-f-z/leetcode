package main

//lint:ignore U1000 unused
func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
	region := map[string]string{}

	for _, r := range regions {
		for i := 1; i < len(r); i++ {
			region[r[i]] = r[0]
		}
	}

	p := region[region1]
	q := region[region2]

	for p != q {
		if p != "" {
			p = region[p]
		} else {
			p = region2
		}

		if q != "" {
			q = region[q]
		} else {
			q = region1
		}
	}

	return p

}
