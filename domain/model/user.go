package model

type UserCreate struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Bio            string `json:"bio"`
	ProfilePicture []byte `json:"profile_picture"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdateUsername struct {
	Username string `json:"username"`
}

type UserUpdatePassword struct {
	Password string `json:"password"`
}

type UserUpdateEmail struct {
	Email string `json:"email"`
}

type UserUpdateProfile struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type UserUpdateProfilePicture struct {
	ProfilePicture []byte `json:"profile_picture"`
}

type UserList struct {
	Username       string `json:"username"`
	Name           string `json:"name"`
	ProfilePicture []byte `json:"profile_picture"`
	Bio            string `json:"bio"`
}