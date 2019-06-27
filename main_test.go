package main

import (
	"testing"
)

// Проверка добавления в конец списка
func TestList_PushBack_Last(t *testing.T) {
	table := []struct {
		item interface{}
	}{
		{"a"},
		{"b"},
	}
	list := &List{}
	for _, row := range table {
		list.pushBack(row.item)
	}

	item := list.Last()
	if item.value != "b" {
		t.Errorf("Item %d not added to end of list ", item.value)
	}
}

// Проверка добавления в начало списка
func TestList_PushFront_First(t *testing.T) {
	table := []struct {
		item interface{}
	}{
		{"a"},
		{"b"},
	}
	list := &List{}
	for _, row := range table {
		list.pushFront(row.item)
	}

	item := list.First()
	if item.value != "b" {
		t.Errorf("Item %d not added to top of list ", item.value)
	}
}

// Проверка правильности подсчета кол-ва элементов списка
func TestList_Len(t *testing.T) {
	table := []struct {
		item interface{}
	}{
		{10},
		{20},
		{30},
		{40},
	}
	list := &List{}
	for _, row := range table {
		list.pushFront(row.item)
	}

	lenList := list.Len()
	if lenList != 4 {
		t.Errorf("Length %d of list is not correct", lenList)
	}

}

// Проверка удаления элемента из списка
func TestItem_Remove_FirstItem(t *testing.T) {
	table := []struct {
		item interface{}
	}{
		{1},
		{2},
		{3},
	}
	list := &List{}
	for _, row := range table {
		list.pushBack(row.item)
	}

	list.First().Remove() // первый элемент
	if list.First().Value() == 1 {
		t.Errorf("First %d item was not deleted", list.First().Value())
	}
}

func TestItem_Remove_MiddleItem(t *testing.T) {
	table := []struct {
		item interface{}
	}{
		{1},
		{2},
		{3},
	}
	list := &List{}
	for _, row := range table {
		list.pushBack(row.item)
	}

	list.First().Next().Remove() // элемент в середине списка
	if list.First().Next().Value() == 2 {
		t.Errorf("Middle %d item was not deleted", list.First().Next().Value())
	}
}

func TestItem_Remove_LastItem(t *testing.T) {
	table := []struct {
		item interface{}
	}{
		{1},
		{2},
		{3},
	}
	list := &List{}
	for _, row := range table {
		list.pushBack(row.item)
	}

	list.Last().Remove() // последний элемент
	if list.Last().Value() == 3 {
		t.Errorf("Last %d item was not deleted", list.Last().Value())
	}
}

func TestItem_Remove_SingleItem(t *testing.T) {
	list := &List{}
	list.pushBack(1)

	list.First().Remove() // единственный элемент
	if list.len != 0 {
		t.Errorf("Single item was not deleted")
	}
}

// Проверка предыдущего элемента в списке
func TestItem_Prev_Value(t *testing.T) {
	table := []struct {
		item interface{}
	}{
		{10},
		{20},
		{30},
		{40},
	}
	list := &List{}
	for _, row := range table {
		list.pushBack(row.item)
	}

	item := list.Last().Prev() // Предыдущий элемент
	if item.Value() != 30 {
		t.Errorf("Item %d not previous", item.Value())
	}
}

// Проверка следующего элемента в списке
func TestItem_Next_Value(t *testing.T) {
	table := []struct {
		item interface{}
	}{
		{10},
		{20},
		{30},
		{40},
	}
	list := &List{}
	for _, row := range table {
		list.pushFront(row.item)
	}

	item := list.First().Next() // Следующий элемент
	if item.Value() != 30 {
		t.Errorf("Item %d not next", item.Value())
	}
}
