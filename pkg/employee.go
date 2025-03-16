package pkg

type Employee struct {
	ID               int    `db:"id" json:"id"`
	FirstName        string `db:"first_name" json:"first_name"`
	LastName         string `db:"last_name" json:"last_name"`
	Gender           string `db:"gender" json:"gender"`
	BirthYear        int    `db:"birth_year" json:"birth_year"`
	StartDate        string `db:"start_date" json:"start_date"`
	ContractType     string `db:"contract_type" json:"contract_type"`
	ContractDuration int    `db:"contract_duration" json:"contract_duration"`
	Department       string `db:"department" json:"department"`
	AnnualLeaveDays  int    `db:"annual_leave_days" json:"annual_leave_days"`
	FreeDays         int    `db:"free_days" json:"free_days"`
	PaidLeaveDays    int    `db:"paid_leave_days" json:"paid_leave_days"`
	ImagePath        string `db:"image_path" json:"image_path"`
}
