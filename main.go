package main

// Список
type List struct {
	head, tail *Item
	len        int
}

// Добавление значения в конец списка
func (l *List) pushBack(value interface{}) {
	item := &Item{value: value, list: l}
	if l.head == nil { // если список пуст
		l.head = item // новый элемент определяем как head
	} else {
		l.tail.next = item // последнему элементу добавляем адрес на новый элемент
		item.prev = l.tail // новому элементу добавляем адрес предыдущего элемента
	}
	l.tail = item // обновляем последний элемент списка
	l.len++
}

// Добавление значения в начало списка
func (l *List) pushFront(value interface{}) {
	item := &Item{value: value, list: l}
	if l.head == nil { // если список пуст
		l.head = item // новый элемент определяем как head
	} else {
		l.head.prev = item // добавляем текущему первому элементу адрес на новый элемент
		item.next = l.head // добавляем текущий адрес первого элемента списка новому элементу
	}
	l.head = item // обновляем текущий первый элемент списка
	l.len++
}

// Длинна списка
func (l *List) Len() int {
	return l.len
}

// Возвращает первый Item
func (l *List) First() *Item {
	return l.head
}

// Возвращает последний Item
func (l *List) Last() *Item {
	return l.tail
}

// Элемент списка
type Item struct {
	value      interface{} // данные узла любого типа
	prev, next *Item       // адрес следующего и предыдущего элемента списка
	list       *List
}

// Возвращает следующий Item
func (i *Item) Next() *Item {
	return i.next
}

// Возвращает предыдущий Item
func (i *Item) Prev() *Item {
	return i.prev
}

// Возвращает значение
func (i *Item) Value() interface{} {
	return i.value
}

//Удаление элемента
func (i *Item) Remove() {
	if i.prev == nil && i.next != nil { // первый элемент
		i.list.head = i.next
		i.list.head.prev = nil
	} else if i.prev != nil && i.next == nil { // последний элемент
		i.list.tail = i.prev
		i.list.tail.next = nil
	} else if i.prev == nil && i.next == nil { // единственный элемент
		i.list.head = nil
		i.list.tail = nil
	} else { // элемент в середине списка
		i.prev.next = i.next
		i.next.prev = i.prev
	}
	i.list.len--
}

func main() {

}
