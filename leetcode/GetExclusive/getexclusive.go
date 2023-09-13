package main

import "fmt"

var count int

func main() {
	n := 0
	fmt.Scan(&n)
	sections := make([]section, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&sections[i].begin, &sections[i].end)
	}
	fmt.Println(deal(sections))
}

func deal(sections []section) byte {
	freeSecs := make([]section, 0)
	freeSecs = append(freeSecs, section{
		begin: 0,
		end:   100000,
	})
	for _, sec := range sections {
		for i, freesec := range freeSecs {
			if freesec.begin < sec.begin && freesec.end > sec.end {
				count++
				freeSecs[i].end = sec.begin
				freeSecs = append(freeSecs, section{
					begin: sec.end,
					end:   100000,
				})
				break
			}
		}
	}
	return byte(count)
}

type section struct {
	begin int
	end   int
}
