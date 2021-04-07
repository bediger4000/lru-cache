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

func Test_BigDelete(t *testing.T) {
	var testInputs = []string{
		"1", "2", "3", "4", "5",
		"6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15",
		"16", "17", "18", "19", "20",
	}
	tbl := NewTable(3)

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
		if !tbl.Delete(key) {
			t.Logf("did not delete key %d, %s, in table", count, str)
		}
	}
}
func Test_BigDelete2(t *testing.T) {
	var testInputs = []string{
		"1", "2", "3", "4", "5",
		"6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15",
		"16", "17", "18", "19", "20",
	}
	tbl := NewTable(3)

	for count, str := range testInputs {
		item := NewCacheStringItem(str)

		if founditem, insertedit := tbl.Insert(item); !insertedit {
			if founditem != nil {
				t.Logf("returned false, count %d, but also returned a cache item", count)
			}
			t.Fatalf("Failed to insert cache item %d for %s", count, str)
		}
	}
	var notInInputs = []string{
		"A1", "B2", "C3", "D4", "E5",
	}
	for count, str := range notInInputs {
		key := NewStringData(str)
		if tbl.Delete(key) {
			t.Logf("deleted key %d, %s, not in table", count, str)
		}
	}
}
func Test_BigDelete3(t *testing.T) {
	var testInputs = []string{
		"1", "2", "3", "4", "5",
		"6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15",
		"16", "17", "18", "19", "20",
	}
	tbl := NewTable(3)

	for count, str := range testInputs {
		item := NewCacheStringItem(str)

		if founditem, insertedit := tbl.Insert(item); !insertedit {
			if founditem != nil {
				t.Logf("returned false, count %d, but also returned a cache item", count)
			}
			t.Fatalf("Failed to insert cache item %d for %s", count, str)
		}
	}
	// out of (input-)order deletes
	var deleteInputs = []string{
		"7", "8", "6", "10", "15",
		"19", "2", "13", "9", "4", "3", "5",
		"12", "11", "14",
		"16", "1", "18", "17", "20",
	}
	for count, str := range deleteInputs {
		key := NewStringData(str)
		if !tbl.Delete(key) {
			t.Logf("did not delete key %d, %s, in table", count, str)
		}
	}
}
