package main

import "testing"

func TestAddDataToRanking(t *testing.T) {
	err := addDataToRanking(Data{
		Name:  "test tetsuji",
		Point: 25,
	})
	if err != nil {
		t.Error(err)
	}
}
