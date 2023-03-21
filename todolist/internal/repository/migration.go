package repository

func migration() {
	err := DB.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}