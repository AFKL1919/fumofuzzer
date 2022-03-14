package sorter

import (
	"sort"

	"github.com/go-resty/resty/v2"
)

type SizeSorter struct {
	FuzzedResps []resty.Response
}

func (sortp *SizeSorter) setFuzzedResps(resps []resty.Response) {
	sortp.FuzzedResps = append(sortp.FuzzedResps, resps...)
}

func (sortp SizeSorter) Len() int {
	return len(sortp.FuzzedResps)
}

func (sortp SizeSorter) Less(i, j int) bool {
	return sortp.FuzzedResps[i].Size() < sortp.FuzzedResps[j].Size()
}

func (sortp SizeSorter) Swap(i, j int) {
	sortp.FuzzedResps[i], sortp.FuzzedResps[j] = sortp.FuzzedResps[j], sortp.FuzzedResps[i]
}

func (sortp SizeSorter) Sort(resps []resty.Response) {
	sortp.setFuzzedResps(resps)
	sort.Sort(sortp)
}
