package main

import "testing"

func TestConfig_getHoldings(t *testing.T) {
	all, err := testApp.currentHoldings()
	if err != nil {
		t.Error("fail to get current holdings from database:", err)
	}

	if len(all) != 2 {
		t.Error("wrong number of rows returned: expected 2 - got", len(all))
	}
}

func TestConfig_getHoldingSlice(t *testing.T) {
	slice := testApp.getHoldingsSlice()

	if len(slice) != 3 {
		t.Error("wrong number of rows returned: expected 3 - got", len(slice))
	}
}
