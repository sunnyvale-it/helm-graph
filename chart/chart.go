package chart

type Chart struct {
	Name    string
	Version string
	Repo    string
	Deps    []Chart
}

func getChartDependencies(chart Chart) []Chart {
	return []Chart{
		{
			Name:    "1",
			Repo:    "2",
			Deps:    nil,
			Version: "3",
		},
		{
			Name:    "x",
			Repo:    "y",
			Deps:    nil,
			Version: "z",
		},
	}
}

func (chart *Chart) Graph() {

	if chart == nil {
		return
	}

	for _, v := range getChartDependencies(*chart) {
		chart.Deps = append(chart.Deps, v)
	}

	for _, v := range chart.Deps {
		if len(v.Deps) > 0 {
			v.Graph()
		}
	}

}
