package main

import "testing"

func TestProcess(t *testing.T) {
	testCases := []struct {
		desc string
		in   []int
	}{
		{
			desc: "",
			in:   []int{1, 2, 3},
		},
		{
			desc: "",
			in:   []int{7, 4, 3},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ch := Process(tC.in)

			for _, i := range tC.in {
				val := <-ch
				if val != i {
					t.Fatalf("exptected reveive: %v, got: %v", i, val)
				}
			}
		})
	}
}
