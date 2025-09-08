package db

type UpdateDB interface{
	UserProfile(profile string) error
	UserBio(bio string) error
}