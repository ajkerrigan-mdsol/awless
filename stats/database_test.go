package stats

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestOpenDbGeneratesIdForNewDb(t *testing.T) {
	f, e := ioutil.TempFile(".", "test.db")
	if e != nil {
		t.Fatal(e)
	}
	defer os.Remove(f.Name())

	db, err := OpenDB(f.Name())
	if err != nil {
		t.Fatal(err)
	}

	newId, err := db.GetStringValue(AWLESS_ID_KEY)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := len(newId), 64; got != want {
		t.Fatalf("got %d; want %d", got, want)
	}
	db.Close()

	db, err = OpenDB(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	id, _ := db.GetStringValue(AWLESS_ID_KEY)
	if got, want := id, newId; got != want {
		t.Fatalf("got %s; want %s", got, want)
	}
}

func TestGetSetDatabaseValues(t *testing.T) {
	f, e := ioutil.TempFile(".", "test.db")
	if e != nil {
		t.Fatal(e)
	}
	defer os.Remove(f.Name())

	db, err := OpenDB(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	value, e := db.GetStringValue("mykey")
	if e != nil {
		t.Fatal(e)
	}
	if got, want := value, ""; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}

	e = db.SetStringValue("mykey", "myvalue")
	if e != nil {
		t.Fatal(e)
	}

	value, e = db.GetStringValue("mykey")
	if e != nil {
		t.Fatal(e)
	}
	if got, want := value, "myvalue"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}

	i, e := db.GetIntValue("myintkey")
	if e != nil {
		t.Fatal(e)
	}
	if got, want := i, 0; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}

	e = db.SetIntValue("myintkey", 10)
	if e != nil {
		t.Fatal(e)
	}

	i, e = db.GetIntValue("myintkey")
	if e != nil {
		t.Fatal(e)
	}
	if got, want := i, 10; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}

	stamp, e := db.GetTimeValue("mytimekey")
	if e != nil {
		t.Fatal(e)
	}
	if got, want := stamp.IsZero(), true; got != want {
		t.Fatalf("got %t, want %t", got, want)
	}

	now := time.Now()
	e = db.SetTimeValue("mytimekey", now)
	if e != nil {
		t.Fatal(e)
	}

	stamp, e = db.GetTimeValue("mytimekey")
	if e != nil {
		t.Fatal(e)
	}
	if got, want := stamp, now; !want.Equal(want) {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func TestAwlessIdGenerator(t *testing.T) {
	id1, _ := generateAwlessId()
	id2, _ := generateAwlessId()

	if got, want := len(id1), 64; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if got, want := len(id2), 64; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
	if got, want := (id1 != id2), true; got != want {
		t.Fatalf("got %t, want %t", got, want)
	}
}