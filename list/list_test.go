package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type SomeType struct {
	Id int
	Name string
}

func TestNewList(t *testing.T) {
	t.Run("Should create list of int", func(t *testing.T) {
		list := NewList[int]()
		assert.NotNil(t, list)
	})

	t.Run("Should create list of strings", func(t *testing.T) {
		list := NewList[string]()
		assert.NotNil(t, list)
	})

	t.Run("Should return first string", func(t *testing.T) {
		list := NewList[string]()
		list.Add("Sumit")
		assert.Equal(t, "Sumit", *list.First())
	})

	t.Run("Should create list of SomeType", func(t *testing.T) {
		list := NewList[SomeType]()
		assert.NotNil(t, list)
	})

	t.Run("Should return first object", func(t *testing.T) {
		list := NewList[SomeType]()
		list.Add(SomeType{
			Id:   1,
			Name: "Sumit",
		})
		assert.NotNil(t, list.First())
	})

	t.Run("Should return nil", func(t *testing.T) {
		list := NewList[SomeType]()
		assert.Nil(t, list.First())
	})

	t.Run("Should find object by predicate", func(t *testing.T) {
		list := NewList[SomeType]()

		list.Add(SomeType{
			Id:   2,
			Name: "Chandu",
		})

		list.Add(SomeType{
			Id:   3,
			Name: "Viraj",
		})

		list.Add(SomeType{
			Id:   1,
			Name: "Sumit",
		})

		item := list.Find(func(item SomeType) bool {
			return item.Id == 2
		})

		assert.Equal(t, 2, item.Id)
	})

}