package analytic

// Enum represents the analytic type
type Enum int

var enums = [...]string{
	"user_count",
	"repo_count",
	// "user_most_repos",
	// "repos_most_recent",
	// "repos_most_stars",
	// "languages",
	// "most_repos_by_language",
}

const (
	// EnumUserCount represents the type for the user_count
	EnumUserCount Enum = iota
	// EnumRepoCount represents the type for the repo_count
	EnumRepoCount
)

func (e Enum) String() string {
	return enums[e]
}
