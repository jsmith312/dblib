package dblib

import "testing"

func TestDeletionAndAddition(t *testing.T) {
	const dbpath = "foo.db"

	db := InitDB(dbpath)
	// Drops the table if it exists
	DropTable(db)
	defer db.Close()
	CreateTable(db)

	items := []Group{
		Group{1, "A"},
		Group{2, "B"},
	}
	StoreItem(db, items)
	readItems := ReadItem(db)
	t.Log(readItems)

	DeleteTable(db)

	items2 := []Group{
		Group{1, "C"},
		Group{3, "D"},
	}

	StoreItem(db, items2)

	readItems2 := ReadItem(db)
	for _, item := range readItems2 {
		if item.ID == 2 || item.Name == "A" {
			t.Error("should not have kept previous entry")
		}
	}
	t.Log(readItems2)
}

func TestDeleteTable(t *testing.T) {
	const dbpath = "foo.db"

	db := InitDB(dbpath)
	// Drops the table if it exists
	DropTable(db)
	defer db.Close()
	CreateTable(db)

	items := []Group{
		Group{1, "A"},
		Group{2, "B"},
	}
	StoreItem(db, items)
	readItems := ReadItem(db)
	t.Log(readItems)
	itemLen := len(readItems)
	if itemLen <= 0 || itemLen > 2 {
		t.Errorf("DB should only contain amount of items inserted(2), received: %d", itemLen)
	}
	DeleteTable(db)
	afterDrop := ReadItem(db)
	t.Log(afterDrop)
	afterLen := len(afterDrop)
	if afterLen != 0 {
		t.Errorf("DB should contain 0 items after deleting, received: %d", afterLen)
	}

}

func TestStroreItem(t *testing.T) {
	const dbpath = "foo.db"

	db := InitDB(dbpath)

	defer db.Close()
	CreateTable(db)

	items := []Group{
		Group{1, "A"},
		Group{2, "B"},
	}
	StoreItem(db, items)

	readItems := ReadItem(db)
	t.Log(readItems)

	items2 := []Group{
		Group{1, "C"},
		Group{3, "D"},
	}
	StoreItem(db, items2)

	readItems2 := ReadItem(db)

	t.Log(readItems2)
}
