package repositories

import (
	"grocery/models/entity"
	"grocery/utils"
)

func IndexWithPage(limit int, offset int) []entity.ShoppingList {
	db := utils.DBConn()
	defer db.Close()

	query := "SELECT id, name, qty, unit FROM shopping_list LIMIT ? OFFSET ?"
	rows, err := db.Query(query, limit, offset)
	defer rows.Close()

	utils.PanicError(err)

	var shoppingLists []entity.ShoppingList
	for rows.Next() {
		var sl entity.ShoppingList

		err = rows.Scan(&sl.Id, &sl.Name, &sl.Qty, &sl.Unit)
		shoppingLists = append(shoppingLists, sl)
	}

	return shoppingLists
}

func Count() int {
	db := utils.DBConn()
	defer db.Close()

	var count int
	query := "SELECT COUNT(*) FROM shopping_list"
	row := db.QueryRow(query)
	_ = row.Scan(&count)
	return count
}

func Create(shoppingList entity.ShoppingList) (int64, error) {
	db := utils.DBConn()
	defer db.Close()

	query := "INSERT INTO shopping_list (name, qty, unit) VALUES(?, ?, ?);"
	stmt, stmtErr := db.Prepare(query)
	utils.PanicError(stmtErr)

	res, queryErr := stmt.Exec(shoppingList.Name, shoppingList.Qty, shoppingList.Unit)
	utils.PanicError(queryErr)

	id, getLastInsertIdErr := res.LastInsertId()
	utils.PanicError(getLastInsertIdErr)

	return id, queryErr
}

func FindById(id int64) (entity.ShoppingList, error) {
	var shoppingList entity.ShoppingList
	db := utils.DBConn()
	defer db.Close()

	query := "SELECT id, name, qty, unit FROM shopping_list WHERE id = ?;"

	row := db.QueryRow(query, id)
	_ = row.Scan(&shoppingList.Id, &shoppingList.Name, &shoppingList.Qty, &shoppingList.Unit)

	return shoppingList, nil
}


func Put(id int64, shoppingList entity.ShoppingList) (entity.ShoppingList, error) {
	db := utils.DBConn()
	defer db.Close()

	query := "UPDATE shopping_list SET name = ?, qty = ?, unit = ? WHERE id = ?"
	stmt, stmtErr := db.Prepare(query)
	utils.PanicError(stmtErr)

	_, queryErr := stmt.Exec(shoppingList.Name, shoppingList.Qty, shoppingList.Unit, id)
	utils.PanicError(queryErr)

	shoppingList.Id = id
	return shoppingList, queryErr
}

func Delete(shoppingList entity.ShoppingList) error {
	db := utils.DBConn()
	defer db.Close()

	query := "DELETE FROM shopping_list WHERE id = ?"
	stmt, stmtErr := db.Prepare(query)
	utils.PanicError(stmtErr)

	_, queryErr := stmt.Exec(shoppingList.Id)
	utils.PanicError(queryErr)

	return queryErr
}