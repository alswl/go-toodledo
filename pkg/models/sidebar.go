package models

type Item struct {
	id    int64
	title string
}

func NewItem(id int64, title string) Item {
	return Item{id: id, title: title}
}

func (i Item) ID() int64 { return i.id }

func (i Item) Title() string { return i.title }

func (i Item) Description() string { return "" }

func (i Item) FilterValue() string { return i.title }
