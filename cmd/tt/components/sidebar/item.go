package sidebar

type Item struct {
	id    int64
	title string
}

func (i Item) ID() int64 { return i.id }

func (i Item) Title() string { return i.title }

func (i Item) Description() string { return "" }

func (i Item) FilterValue() string { return i.title }

type ItemChangeSubscriber func(tab string, item Item) error
