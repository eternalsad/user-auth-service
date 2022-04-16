package models

type User struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	EmployeeID int64  `json:"employee_id" binding:"required"`
}

type Employee struct {
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
