package repositories

import (
	"grocery/models/entity"
	"grocery/utils"
)

func IndexWithPageDosenPA(limit int, offset int) (dosenPA []entity.DosenPA) {
	db := utils.DBConn()
	defer db.Close()

	query := "SELECT dosen_id,name,identifier,email,age FROM dosen_pa LIMIT ? OFFSET ?"
	rows, err := db.Query(query, limit, offset)
	defer rows.Close()

	utils.PanicError(err)

	for rows.Next() {
		var pa entity.DosenPA
		err = rows.Scan(&pa.Id, &pa.Name, &pa.Identifier, &pa.Email, &pa.Age)
		dosenPA = append(dosenPA, pa)
	}

	return
}

func CountDosenPA() int {
	db := utils.DBConn()
	defer db.Close()

	var count int
	query := "SELECT COUNT(*) FROM dosen_pa"
	row := db.QueryRow(query)
	_ = row.Scan(&count)
	return count
}

func CreateDosenPA(dosenPA entity.DosenPA) (id int64, queryErr error) {
	db := utils.DBConn()
	defer db.Close()

	query := "INSERT INTO dosen_pa (name, identifier, email, age) VALUES (?, ?, ?, ?)"
	stmt, stmtErr := db.Prepare(query)
	utils.PanicError(stmtErr)

	res, queryErr := stmt.Exec(dosenPA.Name, dosenPA.Identifier, dosenPA.Email, dosenPA.Age)
	utils.PanicError(queryErr)

	id, getLastInsertIdErr := res.LastInsertId()
	utils.PanicError(getLastInsertIdErr)

 	return
}

func FindDosenPAById(id int64) (dosenPA entity.DosenPA, err error) {
	db := utils.DBConn()
	defer db.Close()

	query := "SELECT dosen_id, name, identifier, email, age FROM dosen_pa WHERE dosen_id = ?"

	row := db.QueryRow(query, id)
	err = row.Scan(&dosenPA.Id, &dosenPA.Name, &dosenPA.Identifier, &dosenPA.Email, &dosenPA.Age)

	return
}


func PutDosenPAByID(id int64, dosenPA entity.DosenPA) (entity.DosenPA, error) {
	db := utils.DBConn()
	defer db.Close()

	query := "UPDATE dosen_pa SET name = ?, identifier = ? , email = ?, age = ? WHERE dosen_id = ?"
	stmt, stmtErr := db.Prepare(query)
	utils.PanicError(stmtErr)

	_, queryErr := stmt.Exec(dosenPA.Name, dosenPA.Identifier, dosenPA.Email, &dosenPA.Age, id)
	utils.PanicError(queryErr)

	dosenPA.Id = id
	return dosenPA, queryErr
}

func DeleteDosenByID(dosenPA entity.DosenPA) (queryErr error) {
	db := utils.DBConn()
	defer db.Close()

	query := "DELETE FROM dosen_pa WHERE dosen_id = ?"
	stmt, stmtErr := db.Prepare(query)
	utils.PanicError(stmtErr)

	_, queryErr = stmt.Exec(dosenPA.Id)
	utils.PanicError(queryErr)

	return
}