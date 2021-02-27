package serializer

import "Goco/internal/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`
	Token     string `json:"token"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Token: user.AuthKey,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
func BuildUsers(users []model.User) (list []User) {
	for _, user := range users {
		list = append(list, BuildUser(user))
	}
	return list
}
