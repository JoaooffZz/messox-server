package db

type UpdateDB interface{
	UpUserProfile(userID int, profile string) error
	UpUserBio(userID int, bio string) error
}