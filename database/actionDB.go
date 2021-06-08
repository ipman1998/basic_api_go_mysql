package dal

import (
	"api/model"
)
//InsertMessage to OrderDB

func InsertMessage(message model.Message) (int64, int64, error) {
	GetConnection()

	sqlQuery := "INSERT Message SET id=?, message=?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(message.Id, message.Message)
	if err != nil {
		return 0, 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	lastInsertId, err := res.LastInsertId()
	return rowsAffected, lastInsertId, err
}

// UpdateMessage in OrderDB
func UpdateMessage(message model.Message) (int64, error) {
	db := GetConnection()

	sqlQuery := "UPDATE Message SET message=? WHERE id=?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(message.Message, message.Id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

//DeleteMessage in OrderDB with parameter is MessageId
func DeleteMessage(id int) (int64, error) {
	db := GetConnection()
	defer db.Close()
	sqlQuery := "DELETE FROM Message WHERE id=?"

	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

//GetMessage from MessageId
func GetMessage(id int) (model.Message) {
	db := GetConnection()

	sqlQuery := "SELECT id, message FROM Message WHERE id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var message model.Message
	if err != nil {
		return message
	}
	res, err := stmt.Query(id)

	if res.Next() {
		res.Scan(&message.Id, &message.Message)
	}
	return message
}

func GetAllMessage() ([] model.Message) {
	db := GetConnection()

	sqlQuery := "SELECT * FROM Message"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var message model.Message
	var messages []model.Message
	if err != nil {
		return messages
	}
	res, err := stmt.Query()

	for ;res.Next()==true; {
		res.Scan(&message.Id, &message.Message)
		messages = append(messages, message)
	}
	return messages
}
