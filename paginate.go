package paginate

// Paginate type
type Paginate struct {
	page    int64
	perPage int64
	items   int64
}

// New creates paginate from page, per page, and items
func New(page, perPage, items int64) Paginate {
	if items < 0 {
		items = 0
	}
	if perPage <= 0 {
		perPage = 1
	}
	if page <= 0 {
		page = 1
	} else if maxPage := items / perPage; page > maxPage {
		page = max(maxPage, 1)
	}
	return Paginate{
		page:    page,
		perPage: perPage,
		items:   items,
	}
}

// FromLimitOffset creates new paginate from limit, offset and count
func FromLimitOffset(limit, offset, count int64) Paginate {
	if count < 0 {
		count = 0
	}
	if limit <= 0 {
		limit = 1
	}
	if offset < 0 {
		offset = 0
	} else if offset > count {
		offset = count
	}
	return Paginate{
		page:    offset / limit,
		perPage: limit,
		items:   count,
	}
}

// Page returns page
func (p Paginate) Page() int64 {
	return p.page
}

// PerPage returns per page
func (p Paginate) PerPage() int64 {
	return p.perPage
}

// Items returns items
func (p Paginate) Items() int64 {
	return p.items
}

// Count is the alias for Items
func (p Paginate) Count() int64 {
	return p.items
}

// Limit returns per page
func (p Paginate) Limit() int64 {
	return p.perPage
}

// Offset returns offset for current page
func (p Paginate) Offset() int64 {
	return (p.page - 1) * p.perPage
}

// LimitOffset returns limit and offet
func (p Paginate) LimitOffset() (limit, offset int64) {
	return p.Limit(), p.Offset()
}

// MaxPage returns max page
func (p Paginate) MaxPage() int64 {
	return max(p.items/p.perPage, 1)
}

// CanPrev returns is current page can go prev
func (p Paginate) CanPrev() bool {
	return p.page > 1
}

// CanNext returns is current page can go next
func (p Paginate) CanNext() bool {
	return p.page < p.MaxPage()
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

// Pages returns page numbers for paginate
//
// around is the number of page around the current page
// ex. if current page is 10 and around is 3
// the result is 0 7 8 9 10 11 12 13 0
//
// edge is the number of page at the edge
// ex. if current page is 10, max page is 20 and edge is 2
// the result is 1 2 0 10 0 19 20
//
// then if current page is 10, max page is 20,
// around is 3, and edge is 2
// the result is
// 1 2 0 7 8 9 10 11 12 13 0 19 20
func (p Paginate) Pages(around, edge int64) []int64 {
	xs := make([]int64, 0)
	maxPage := p.MaxPage()

	var current int64 = 1

	for m := min(edge, p.page-around-1); current <= m; current++ {
		xs = append(xs, current)
	}

	if current < p.page-around {
		xs = append(xs, 0)
	}

	current = max(current, p.page-around)
	for m := min(p.page+around, maxPage); current <= m; current++ {
		xs = append(xs, current)
	}

	if current < maxPage-edge {
		xs = append(xs, 0)
	}

	current = max(current, maxPage-edge+1)
	for ; current <= maxPage; current++ {
		xs = append(xs, current)
	}

	return xs
}
