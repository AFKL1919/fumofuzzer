package filter

import (
	"crypto/md5"
	"fmt"
)

type Md5PayloadFilter struct{}

func (fitler *Md5PayloadFilter) Encode(value string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(value)))
}
