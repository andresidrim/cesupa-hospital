package enums

type Sex string

const (
	Male   Sex = "male"
	Female Sex = "famale"
)

type BloodType string

const (
	APositive  BloodType = "A+"
	ANegative  BloodType = "A-"
	BPositive  BloodType = "B+"
	BNegative  BloodType = "B-"
	ABPositive BloodType = "AB+"
	ABNegative BloodType = "AB-"
	OPositive  BloodType = "O+"
	ONegative  BloodType = "O-"
)
