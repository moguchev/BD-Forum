package messages

const (
	UserNotFound          = "Can't find user with nickname: "
	UserAlreadyExists     = "User already exists"
	UserWasCreated        = "User was created"
	ConflictsInUserUpdate = "New user data has conflicts with existed data"
	ForumAlreadyExists    = "Forum already exists"
	ForumNotFound         = "Can't find forum: "
	ThreadAlreadyExists   = "Thread already exists"
	ThreadNotFound        = "Can't find thread with slug or id: "
	ParentNotFound        = "Can't find parrent post: "
)
