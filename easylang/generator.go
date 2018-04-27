package easylang

type Generator interface {
	GenerateCode(name string) string
}