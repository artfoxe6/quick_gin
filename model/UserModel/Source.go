package UserModel

func (users *Users) ToArray() []map[string]interface{} {
	temp := make([]map[string]interface{}, 0)
	for _, v := range *users {
		temp = append(temp, map[string]interface{}{
			"id":            v.ID,
			"user_name":     v.UserName,
			"age":           v.Age,
			"last_login_at": v.LastLoginAt,
			"articles":      v.Articles.ToArray(),
		})
	}
	return temp
}

func (user *User) Source() map[string]interface{} {
	return map[string]interface{}{
		"user_name":     user.UserName,
		"age":           user.Age,
		"last_login_at": user.LastLoginAt,
	}
}
