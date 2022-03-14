package sorter

import (
	"sort"

	"github.com/go-resty/resty/v2"
)

var SORTER_MAP = map[string]Sorter{
	"size": new(SizeSorter),
}

type Sorter interface {
	sort.Interface
	Sort([]resty.Response)
}
