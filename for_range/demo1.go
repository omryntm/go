package for_range

func Demo1() {
	sehirler := []string{"Ankara", "İstanbul", "İzmir"}
	for i := 0; i < len(sehirler); i++ {
		println(sehirler[i])
	}
	for i, sehir := range sehirler {
		println(i, sehir)
	}
}
