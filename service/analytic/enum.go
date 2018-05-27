package analyticsvc

// Enum represents the analytic type
// type Enum string

const (
	EnumUserCount                 = "user_count"
	EnumRepoCount                 = "repo_count"
	EnumReposMostRecent           = "repos_most_recent"
	EnumRepoCountByUser           = "repo_count_by_user"
	EnumReposMostStars            = "repos_most_stars"
	EnumMostPopularLanguage       = "languages_most_popular"
	EnumLanguageCountByUser       = "language_count_by_user"
	EnumMostRecentReposByLanguage = "repos_most_recent_by_language"
	EnumReposByLanguage           = "repos_by_language"
)

// var enums = [...]string{
// 	"user_count",
// 	"repo_count",
// 	// "user_most_repos",
// 	// "repos_most_recent",
// 	// "repos_most_stars",
// 	// "languages",
// 	// "most_repos_by_language",
// }

// const (
// 	// EnumUserCount represents the type for the user_count
// 	EnumUserCount Enum = iota
// 	// EnumRepoCount represents the type for the repo_count
// 	EnumRepoCount
// )

// func (e Enum) String() string {
// 	return enums[e]
// }
