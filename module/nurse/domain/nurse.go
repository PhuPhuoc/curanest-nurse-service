package nursedomain

import "github.com/google/uuid"

// Nurse struct
type Nurse struct {
	id                 uuid.UUID
	gender             bool
	nurse_picture      string
	nurse_name         string
	citizen_id         string
	dob                string
	address            string
	ward               string
	district           string
	city               string
	current_work_place string
	education_level    string
	experience         string
	certificate        string
	google_drive_url   string
	slogan             string
	rate               float64
}

// Các hàm getter
func (n *Nurse) GetID() uuid.UUID {
	return n.id
}

func (n *Nurse) GetNursePicture() string {
	return n.nurse_picture
}

func (n *Nurse) GetNurseName() string {
	return n.nurse_name
}

func (n *Nurse) GetGender() bool {
	return n.gender
}

func (n *Nurse) GetCitizenID() string {
	return n.citizen_id
}

func (n *Nurse) GetDOB() string {
	return n.dob
}

func (n *Nurse) GetAddress() string {
	return n.address
}

func (n *Nurse) GetWard() string {
	return n.ward
}

func (n *Nurse) GetDistrict() string {
	return n.district
}

func (n *Nurse) GetCity() string {
	return n.city
}

func (n *Nurse) GetCurrentWorkPlace() string {
	return n.current_work_place
}

func (n *Nurse) GetEducationLevel() string {
	return n.education_level
}

func (n *Nurse) GetExperience() string {
	return n.experience
}

func (n *Nurse) GetCertificate() string {
	return n.certificate
}

func (n *Nurse) GetGoogleDriveURL() string {
	return n.google_drive_url
}

func (n *Nurse) GetSlogan() string {
	return n.slogan
}

func (n *Nurse) GetRate() float64 {
	return n.rate
}

// Hàm khởi tạo Nurse
func NewNurse(
	id uuid.UUID,
	gender bool,
	nursePicture string,
	nurseName string,
	citizenID string,
	dob string,
	address string,
	ward string,
	district string,
	city string,
	currentWorkPlace string,
	educationLevel string,
	experience string,
	certificate string,
	googleDriveURL string,
	slogan string,
	rate float64,
) (*Nurse, error) {
	nurse := &Nurse{
		id:                 id,
		gender:             gender,
		nurse_picture:      nursePicture,
		nurse_name:         nurseName,
		citizen_id:         citizenID,
		dob:                dob,
		address:            address,
		ward:               ward,
		district:           district,
		city:               city,
		current_work_place: currentWorkPlace,
		education_level:    educationLevel,
		experience:         experience,
		certificate:        certificate,
		google_drive_url:   googleDriveURL,
		slogan:             slogan,
		rate:               rate,
	}

	return nurse, nil
}
