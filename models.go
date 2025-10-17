type Industry string

const (
	Information_Technology   Industry = "Information Technology"
	Sales                    Industry = "Sales"
	Automotive               Industry = "Automotive"
	Retail_Ecommerce         Industry = "Retail & E-commerce"
	Hospitality_Tourism      Industry = "Hospitality & Tourism"
	Legal_Services           Industry = "Legal"
	Education                Industry = "Education"
	Transportation_Logistics Industry = "Transportation & Logistics"
	Entertainment_Media      Industry = "Entertainment & Media"
	Finance_Banking          Industry = "Finance & Banking"
	Real_Estate              Industry = "Real Estate"
	Healthcare               Industry = "Healthcare"
	Telecommunications       Industry = "Telecommunications"
)

type Job_Title struct {
	Title string
	Department
}

type Employee struct {
	Firstname string
	Lastname  string
}

type Company struct {
	Name        string
	Divisions   []Division
	Industry    Industry
	C_Headcount int64
}

type Division struct {
	Name        string
	Director    Employee
	D_Headcount int64
}

type Department struct {
	Name    string
	Manager Employee
}



