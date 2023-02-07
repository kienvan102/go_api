package studentmodel

type Student struct{
	Id int `json:"id" gorm:"column:id"`
	FullName string `json:"fullname" gorm:"column:fname"`
	Sex string `json:"sex" gorm:"column:sex"`
	Phone int `json:"phone" gorm:"column:phone"`
	Email string `json:"email" gorm:"column:email"`
}
func (Student) TableName() string { return "students" }

type StudentCreate struct{
	Id int `json:"id" gorm:"column:id"`
	FullName string `json:"fullname" gorm:"column:fname"`
	Sex string `json:"sex" gorm:"column:sex"`
	Phone int `json:"phone" gorm:"column:phone"`
	Email string `json:"email" gorm:"column:email"`
}
func (StudentCreate) TableName() string { return Student{}.TableName()}


type StudentUpdate struct{
	Phone *int `json:"phone" gorm:"column:phone"`
	Email *string `json:"email" gorm:"column:email"`
}
func (StudentUpdate) TableName() string { return Student{}.TableName()}


