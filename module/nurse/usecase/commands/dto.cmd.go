package nursecommands

type CreateNurseAccountCmdDTO struct {
	MajorId          string `json:"major-id"`
	NursePicture     string `json:"nurse-picture"`
	FullName         string `json:"full-name"`
	CitizenId        string `json:"citizen-id"`
	Gender           bool   `json:"gender"`
	PhoneNumber      string `json:"phone-number"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	Dob              string `json:"dob"`
	Address          string `json:"address"`
	Ward             string `json:"ward"`
	District         string `json:"district"`
	City             string `json:"city"`
	CurrentWorkPlace string `json:"current-work-place"`
	EducationLevel   string `json:"education-level"`
	Experience       string `json:"experience"`
	Certificate      string `json:"certificate"`
	GoogleDriveUrl   string `json:"google-drive-url"`
	Slogan           string `json:"slogan"`
}

type AccountInfoDTO struct {
	RoleName    string `json:"role-name"`
	FullName    string `json:"full-name"`
	PhoneNumber string `json:"phone-number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type ResponseCreateAccountDTO struct {
	Id string `json:"id"`
}
