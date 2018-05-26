package database

type enum int

var enums = [...]string{"analytics", "profiles", "repos", "users"}

const (
	Analytics enum = iota
	Profiles
	Repos
	Users
)

func (e enum) String() string {
	return enums[e]
}
