package chart

type Chart struct {
	Name    string
	Version string
	Repo    string
	Deps    []Chart
}

func (c *Chart) Graph() {
}