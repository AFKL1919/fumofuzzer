package sorter

import (
	"sort"

	"github.com/go-resty/resty/v2"
)

type SizeSorter []resty.Response

func (sortp SizeSorter) Len() int {
	return len(sortp)
}

func (sortp SizeSorter) Less(i, j int) bool {
	return sortp[i].Size() < sortp[j].Size()
}

func (sortp SizeSorter) Swap(i, j int) {
	sortp[i], sortp[j] = sortp[j], sortp[i]
}

func (sortp SizeSorter) Sort(resps []resty.Response) {
	sortp = resps
	sort.Sort(sortp)
}
