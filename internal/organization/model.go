package organization

type Organization struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Category   string `json:"category" gorm:"size:100;not null"` // ex: "Pimpinan"
	Jabatan    string `json:"jabatan" gorm:"size:100;not null"`
	Nama       string `json:"nama" gorm:"size:150;not null"`
	Quote      string `json:"quote" gorm:"type:text"`
	ProfilePic string `json:"profilePic" gorm:"size:255"` // URL image profile
}
