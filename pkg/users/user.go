package users

type User struct {
	UserName     string
	CreatedTasks []int // Которые он создал
	UserTasks    []int // Которые ему задали
	ChatId       int64
}
