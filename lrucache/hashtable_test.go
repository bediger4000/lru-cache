package lru

import "testing"

func Test_Insert(t *testing.T) {
	tbl := NewTable(11)

	keyA := NewStringData("A")

	if node := tbl.Lookup(keyA); node != nil {
		t.Fatalf("Found non-inserted key \"A\"")
	}

	item := NewCacheStringItem("bubbles")

	if founditem, insertedit := tbl.Insert(item); insertedit {
	} else {
		if founditem != nil {
			t.Log("returned false, but also returned a cache item")
		}
		t.Fatalf("Failed to insert cache item")
	}

	key := NewStringData("bubbles")
	if node := tbl.Lookup(key); node == nil {
		t.Fatalf("did not find newly inserted key/value")
	}
}
func Test_BigInsert(t *testing.T) {
	var testInputs = []string{
		"1", "2", "3", "4", "5",
		"6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15",
		"16", "17", "18", "19", "20",
	}
	tbl := NewTable(11)

	for count, str := range testInputs {
		item := NewCacheStringItem(str)

		if founditem, insertedit := tbl.Insert(item); !insertedit {
			if founditem != nil {
				t.Logf("returned false, count %d, but also returned a cache item", count)
			}
			t.Fatalf("Failed to insert cache item %d for %s", count, str)
		}
	}
	for count, str := range testInputs {
		key := NewStringData(str)
		if node := tbl.Lookup(key); node == nil {
			t.Logf("did not find key %d, %s, in table", count, str)
		}
	}
}
