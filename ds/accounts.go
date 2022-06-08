package ds

import (
	"atms/models"
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func GetUserAccounts(db *sql.DB) ([]models.Account, error) {
	rows, err := db.Query("SELECT account_id, username, status FROM user_account")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()

	var accounts []models.Account
	for rows.Next() {
		var account models.Account
		err = rows.Scan(&account.AccountID, &account.UserName, &account.Status)
		if err != nil {
			panic(err.Error())
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func GetUserWithIDAccount(db *sql.DB, username string) (models.Account, error) {
	rows := db.QueryRow("SELECT account_id, username, status FROM user_account WHERE username = ? ", username)

	var account models.Account
	err := rows.Scan(&account.AccountID, &account.UserName, &account.Status)
	if err != nil {
		panic(err.Error())
	}

	return account, nil
}

func CreateUserAccount(db *sql.DB, user models.Account) error {
	fmt.Println("user.UserName, user.Password, user.Status ", user.UserName, user.Password, user.Status)
	rows, err := db.Exec("INSERT INTO user_account(username, password, status) VALUES (?,?,?) ", user.UserName, user.Password, user.Status)
	if err != nil {
		return err
	}

	rowCount, err := rows.RowsAffected()
	if int(rowCount) < 1 || err != nil {
		return errors.New("couldn't insert user account")
	}

	return nil
}

func UpdateUserAccount(db *sql.DB, userID int, user models.Account) error {
	rows, err := db.Exec("UPDATE user_account SET username = ?, password = ?, status = ? WHERE account_id = ? ", user.UserName, user.Password, user.Status, userID)
	if err != nil {
		return err
	}

	rowCount, err := rows.RowsAffected()
	if int(rowCount) < 1 || err != nil {
		return errors.New("couldn't update user account for accountID " + strconv.Itoa(user.AccountID))
	}

	return nil
}

func DeleteUserAccount(db *sql.DB, accountID int) error {
	rows, err := db.Exec("DELETE FROM user_account where account_id = ?", accountID)
	if err != nil {
		panic(err.Error())
	}

	rowCount, err := rows.RowsAffected()
	if int(rowCount) < 1 || err != nil {
		return errors.New("couldn't delete user account for accountID " + strconv.Itoa(accountID))
	}

	return nil
}

func AuthenticateUserAccount(db *sql.DB, auth models.Authenticate) (bool, error) {
	rows := db.QueryRow("SELECT COUNT(*) FROM user_account WHERE username = ? AND password = ? ", auth.UserName, auth.Password)

	var count int

	err := rows.Scan(&count)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("count : ", count)
	validUser := false
	if count == 1 {
		validUser = true
		return validUser, nil
	}

	return validUser, nil
}
