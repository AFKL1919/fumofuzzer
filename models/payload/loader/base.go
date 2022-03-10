package loader

type PayloadLoader interface {
	Load(string) []string
}
