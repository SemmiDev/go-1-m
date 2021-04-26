package repositories

import (
	"grocery/models/entity"
	"grocery/utils"
)

func IndexWithPageStudent(limit int, offset int) (student []entity.Student) {
	db := utils.DBConn()
	defer db.Close()

	query := "SELECT student_id,name,identifier,email,age, dosen_pa_id FROM student LIMIT ? OFFSET ?"
	rows, err := db.Query(query, limit, offset)
	defer rows.Close()

	utils.PanicError(err)

	for rows.Next() {
		var std entity.Student
		err = rows.Scan(&std.Id, &std.Name, &std.Identifier, &std.Email, &std.Age, std.DosenPAID)
		student = append(student, std)
	}

	return
}

func CountStudent() int {
	db := utils.DBConn()
	defer db.Close()

	var count int
	query := "SELECT COUNT(*) FROM student"
	row := db.QueryRow(query)
	_ = row.Scan(&count)
	return count
}

func CreateStudent(student entity.Student) (id int64, queryErr error) {
	db := utils.DBConn()
	defer db.Close()

	query := "INSERT INTO student (name, identifier, email, age, dosen_pa_id) VALUES (?, ?, ?, ?, ?)"
	stmt, stmtErr := db.Prepare(query)
	utils.PanicError(stmtErr)

	res, queryErr := stmt.Exec(student.Name, student.Identifier, student.Email, student.Age, &student.DosenPAID)
	utils.PanicError(queryErr)

	id, getLastInsertIdErr := res.LastInsertId()
	utils.PanicError(getLastInsertIdErr)

	return
}

func FindStudentById(id int64) (student entity.Student, err error) {
	db := utils.DBConn()
	defer db.Close()

	query := "SELECT student_id, name, identifier, email, age, dosen_pa_id FROM student WHERE student_id = ?"

	row := db.QueryRow(query, id)
	err = row.Scan(&student.Id, &student.Name, &student.Identifier, &student.Email, &student.Age, &student.DosenPAID)

	return
}

func FindStudentByIdentifier(identifier string) (student entity.Student, err error) {
	db := utils.DBConn()
	defer db.Close()

	query := "SELECT student_id, name, identifier, email, age, dosen_pa_id FROM student WHERE student_identifier = ?"

	row := db.QueryRow(query, identifier)
	err = row.Scan(&student.Id, &student.Name, &student.Identifier, &student.Email, &student.Age, &student.DosenPAID)

	return
}

func PutStudentByID(id int64, student entity.Student) (entity.Student, error) {
	db := utils.DBConn()
	defer db.Close()

	query := "UPDATE student SET name = ?, identifier = ? , email = ?, age = ?, dosen_pa_id = ? WHERE student_id = ?"
	stmt, stmtErr := db.Prepare(query)
	utils.PanicError(stmtErr)

	_, queryErr := stmt.Exec(student.Name, student.Identifier, student.Email, &student.Age, &student.DosenPAID, id)
	utils.PanicError(queryErr)

	student.Id = id
	return student, queryErr
}

func DeleteStudentByID(student entity.Student) (queryErr error) {
	db := utils.DBConn()
	defer db.Close()

	query := "DELETE FROM student WHERE student_id = ?"
	stmt, stmtErr := db.Prepare(query)
	utils.PanicError(stmtErr)

	_, queryErr = stmt.Exec(student.Id)
	utils.PanicError(queryErr)

	return
}

func StudentJoinDosenPA() (studentJoinDosen []entity.Join, err error) {
	db := utils.DBConn()
	defer db.Close()

	query := "SELECT p.student_id, p.name, p.identifier, p.email, p.age, p.dosen_pa_id, dp.dosen_id, dp.name, dp.identifier, dp.email, dp.age FROM student p JOIN dosen_pa dp on p.dosen_pa_id = dp.dosen_id"
	rows, err := db.Query(query)
	for rows.Next() {
		var studentjoindosenpa entity.Join
		err = rows.Scan(
			&studentjoindosenpa.StudentId, &studentjoindosenpa.StudentName, &studentjoindosenpa.StudentIdentifier, &studentjoindosenpa.StudentEmail, &studentjoindosenpa.StudentAge, &studentjoindosenpa.StudentDosenPAId,
			&studentjoindosenpa.DosenPAId, &studentjoindosenpa.DosenPAName, &studentjoindosenpa.DosenPAIdentifier, &studentjoindosenpa.DosenPAEmail, &studentjoindosenpa.DosenPAAge,
		)
		studentJoinDosen = append(studentJoinDosen, studentjoindosenpa)
	}

	return
}

func StudentJoinDosenPAWithPage(limit int, offset int) (student []entity.Student) {
	db := utils.DBConn()
	defer db.Close()

	query := "SELECT p.student_id, p.name, p.identifier, p.email, p.age, p.dosen_pa_id, dp.dosen_id, dp.name, dp.identifier, dp.email, dp.age FROM student p JOIN dosen_pa dp on p.dosen_pa_id = dp.dosen_id LIMIT ? OFFSET ?"
	rows, err := db.Query(query, limit, offset)
	defer rows.Close()

	utils.PanicError(err)

	for rows.Next() {
		var std entity.Student
		err = rows.Scan(&std.Id, &std.Name, &std.Identifier, &std.Email, &std.Age, std.DosenPAID)
		student = append(student, std)
	}

	return
}