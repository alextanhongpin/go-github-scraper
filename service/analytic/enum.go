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
