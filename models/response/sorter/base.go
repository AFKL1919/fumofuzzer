package sorter

import (
	"sort"

	"github.com/go-resty/resty/v2"
)

type BaseSorter interface {
	sort.Interface
	Sort([]resty.Response) []resty.Response
}
