package data

type ActiveProject struct {
	RootPath string
}

func GetProject() ActiveProject {
	return ActiveProject{
		RootPath: "C:\\Users\\User\\Project\\KickIt\\blog\\",
	}
}
