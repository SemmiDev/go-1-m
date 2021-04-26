package entity

type (
	Student struct {
		Id    		int64  `json:"id"`
		Name  		string `json:"name" form:"name" binding:"required"`
		Identifier  string `json:"identifier" form:"identifier" binding:"required"`
		Email   	string `json:"email" form:"email" binding:"required,email"`
		Age   		int    `json:"age"  form:"age" binding:"required,numeric,gte=17"`
		DosenPAID   int64  `json:"dosen_pa" form:"dosen_pa" binding:"required"`
		Error 		error
	}

	DosenPA struct {
		Id    		int64  `json:"id"`
		Name  		string `json:"name" form:"name" binding:"required"`
		Identifier  string `json:"identifier" form:"identifier" binding:"required"`
		Email   	string `json:"email"  form:"email" binding:"required,email"`
		Age   		int    `json:"age"  form:"age" binding:"required,numeric,gte=20"`
		Error 		error
	}

	Join struct {
		StudentId    		int64  `json:"student_id"`
		StudentName  		string `json:"student_name"`
		StudentIdentifier  	string `json:"student_identifier"`
		StudentEmail   		string `json:"student_email"`
		StudentAge   		int    `json:"student_age"`
		StudentDosenPAId   	int64  `json:"student_dosen_pa_id"`
		DosenPAId    		int64  `json:"dosen_pa_id"`
		DosenPAName  		string `json:"dosen_pa_name"`
		DosenPAIdentifier  	string `json:"dosen_pa_identifier"`
		DosenPAEmail   		string `json:"dosen_pa_email"`
		DosenPAAge   		int    `json:"dosen_pa_age"`
	}
)