package data

// SnippetMarker is a specific type to declare snippet marks
type SnippetMarker string

const (
	// SnippetMarkerBash is the default snippet mark for bash files
	SnippetMarkerBash SnippetMarker = "###(?P<snippetname>[^#]+)###$"
	// SnippetMarkerGolang is the default snippet mark for Golang
	SnippetMarkerGolang SnippetMarker = "// --> (?P<snippetname>[^#]+)$"
)
