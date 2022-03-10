package filter

type PayloadFilter interface {
	Encode(string) string
}

type NonePayloadFilter struct{}

func (fitler *NonePayloadFilter) Encode(value string) string {
	return value
}
