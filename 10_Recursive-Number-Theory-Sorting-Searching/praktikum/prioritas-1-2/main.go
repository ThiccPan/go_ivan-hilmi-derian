package main

import "fmt"

type Pair struct {
	name  string
	count int
	next  *Pair
}

func (p Pair) String() string {
	return fmt.Sprint(p.name, "->", p.count)
}

type ItemList struct {
	head *Pair
}

func (l *ItemList) insert(item Pair) {
	list := &item
	p := l.head
	// list pair kosong
	if l.head == nil {
		l.head = list
		return
	}
	// item pair baru lebih kecil dari head
	if list.count <= l.head.count {
		list.next = l.head
		l.head = list
		return
	}
	// jalan sampai ketemu item pair di list yang lebih kecil
	for p.next != nil {
		if item.count <= p.next.count {
			item.next = p.next
			p.next = &item
			return
		}
		p = p.next
	}
	p.next = &item
}

func mostAppearItem(items []string) []Pair {
	itemMap := make(map[string]int)
	max := 0

	// masukkan ke map
	for _, v := range items {
		itemMap[v]++
	}
	itemList := ItemList{}

	// map ke linkedlist lalu di sort
	for i, v := range itemMap {
		max++
		item := Pair{name: i, count: v}
		itemList.insert(item)
	}

	itemPair := make([]Pair, 0)
	p := itemList.head

	// masuk ke slice
	for {
		itemPair = append(itemPair, *p)
		if p.next == nil {
			break
		}
		p = p.next
	}

	return itemPair
}

func main() {
	fmt.Println(mostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(mostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(mostAppearItem([]string{"football", "basketball", "tennis"}))
}
