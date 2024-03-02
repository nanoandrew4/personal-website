package components

var (
	projects = []Project{
		{
			Name:                 "nsteg",
			DescriptionLocaleKey: "projects.nsteg-description",
			Link:                 "/nsteg",
		},
		{
			Name:                 "calculator",
			DescriptionLocaleKey: "projects.calculator-description",
			Link:                 "/calculator",
		},
		{
			Name:                 "Adventure Tracks",
			DescriptionLocaleKey: "projects.adventure-tracks-description",
			Link:                 "/adventure-tracks",
		},
		{
			Name:                 "Lineage OS for umi/cmi (Mi 10)",
			DescriptionLocaleKey: "projects.los",
			Link:                 "/los",
		},
		{
			Name:                 "iArt",
			DescriptionLocaleKey: "projects.iart",
			Link:                 "/iArt",
		},
	}
)

type Project struct {
	Name                 string
	DescriptionLocaleKey string
	Link                 string
}
