package coverage

import (
	"os"
	"testing"
	"time"
	"strings"
	"reflect"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

var ms = "1 2 \n 3 4"

// WRITE YOUR CODE BELOW
func TestPeopleLen(t *testing.T) {
	p := Person{
		firstName: "John",
		lastName: "Doe",
		birthDay: time.Now(),
	}

	ps := People{}

	if ps.Len() != 0 {
		t.Error("Incorrect length provided: expected 0, got ", ps.Len())
	}

	ps = append(ps, p)
	if ps.Len() != 1 {
		t.Error("Incorrect length provided: expected 1, got ", ps.Len())
	}
}

func TestPeopleLess(t *testing.T) {
	data := []struct{
		p1 Person
		p2 Person
		expected bool
	}{
		{
			Person{ 
				firstName: "John", 
				lastName: "Doe", 
				birthDay: time.Date(2009, time.October, 10, 23, 0, 0, 0, time.UTC),
			},
			Person{ 
				firstName: "Jim", 
				lastName: "Doe", 
				birthDay: time.Date(2009, time.October, 10, 23, 0, 0, 0, time.UTC),
			},
			false,
		},
		{
			Person{ 
				firstName: "Jim", 
				lastName: "Doe", 
				birthDay: time.Date(2009, time.October, 10, 23, 0, 0, 0, time.UTC),
			},
			Person{ 
				firstName: "Jim", 
				lastName: "Doe", 
				birthDay: time.Date(2009, time.October, 9, 23, 0, 0, 0, time.UTC),
			},
			true,
		},
		{
			Person{ 
				firstName: "Jim", 
				lastName: "Doe", 
				birthDay: time.Date(2009, time.October, 10, 23, 0, 0, 0, time.UTC), 
			},
			Person{ 
				firstName: "Jim", 
				lastName: "Doe", 
				birthDay: time.Date(2009, time.October, 10, 23, 0, 0, 0, time.UTC),
			},
			false,
		},
		{
			Person{ 
				firstName: "Jim", 
				lastName: "Doe", 
				birthDay: time.Date(2009, time.October, 10, 23, 0, 0, 0, time.UTC),
			},
			Person{
				firstName: "Jim", 
				lastName: "Johnson", 
				birthDay: time.Date(2009, time.October, 10, 23, 0, 0, 0, time.UTC),
			},
			true,
		},
	}

	for _, d := range data {
		ps := People{}
		ps = append(ps, d.p1, d.p2)

		result := ps.Less(0, 1)
		if result != d.expected {
			t.Errorf("Less() returns incorrect result: p1: %+v, p2: %+v expected %t, got %t", d.p1, d.p2, d.expected, result)
		}
	}
}

func TestPeopleSwap(t *testing.T) {
	ps := People{
		Person{
			firstName: "John",
			lastName: "Doe",
			birthDay: time.Date(2009, time.October, 10, 22, 0, 0, 0, time.UTC),
		},
		Person{
			firstName: "Jim",
			lastName: "Doe",
			birthDay: time.Date(2009, time.October, 10, 23, 0, 0, 0, time.UTC),
		},
	}

	ps.Swap(0, 1)
	if ps[0].firstName != "Jim" {
		t.Errorf("Elements not swapped correctly")
	}
}

func TestMatrixNew(t *testing.T) {
	m, err := New(ms)
	if err != nil {
		t.Fatal(err)
	}
	if m.rows != 2 || m.cols != 2 {
		t.Errorf("Incorrect number of rows or columns: expected 2*2, got %d*%d", m.rows, m.cols)
	}

	if len(m.data) != 4 {
		t.Error("Data has incorrect length: expected 4, got ", len(m.data))
	}

	ms = "1 2 \n 3"
	m, err = New(ms)

	if err.Error() != "Rows need to be the same length" {
		t.Fatal(err)
	}

	ms = "1 2 \n 3 four"
	m, err = New(ms)

	expErr := "strconv.Atoi"

	if !strings.Contains(err.Error(), expErr) {
		t.Errorf("Unexpected error, expected %s, got %s", expErr, err.Error())
	}
}

func TestMatrixRows(t *testing.T) {
	ms = "1 2 \n 3 4"
	m, err := New(ms)
	if err != nil {
		t.Fatal("Error when creating the matrix")
	}

	result := m.Rows()

	if !reflect.DeepEqual(result, [][]int{
		{1, 2},
		{3, 4},
	}) {
		t.Errorf("Incorrect data returned")
	}
}

func TestMatrixCols(t *testing.T) {
	m, err := New(ms)
	if err != nil {
		t.Fatal("Error when creating the matrix")
	}

	result := m.Cols()

	if !reflect.DeepEqual(result, [][]int{
		{1, 3},
		{2, 4},
	}) {
		t.Errorf("Incorrect data returned")
	}
}

func TestMatrixSet(t *testing.T) {
	m, err := New(ms)
	if err != nil {
		t.Fatal("Error when creating the matrix")
	}

	result := m.Set(1, 1, 10)
	if result && m.data[1 * m.cols + 1] != 10 {
		t.Error("Value was not set: expected 10, got ", m.data[1 * m.cols + 1])
	}

	result = m.Set(3, 1, 10)
	if result {
		t.Error("Set() returned true when out of boundary index was provided")
	}
}
