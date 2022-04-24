package dao

import (
	"github.com/ProjectTravelPartner/accmgtms/models"
	"github.com/ProjectTravelPartner/dbclient"
)

func AccountCreate(data models.Account) (uint64, error) {
	qry := `
			INSERT INTO accounts
				(id,name,DOB,email,contact,pwd)
			VALUES
				(?,?,?,?,?,?)`
	id, err := dbclient.ExecGetID(qry, data.ID, data.Name, data.DOB, data.EMail, data.Contact, data.Password)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func AccountGet(id uint64) (models.Account, error) {
	qry := `
			SELECT 
				*
			FROM
				accounts
			WHERE
				id=?`
	row := dbclient.QueryRow(qry, id)
	var data models.Account
	row.Scan(
		&data.ID,
		&data.Name,
		&data.DOB,
		&data.EMail,
		&data.Contact,
		&data.Password,
	)
	return data, nil

}

func VerifyUser(email string, pwd string) uint64 {
	qry := `SELECT 
				id
			FROM
				accounts
			WHERE
				email=?
			AND
				pwd=?
			limit 1`
	rows := dbclient.QueryRow(qry, email, pwd)
	var out uint64
	rows.Scan(&out)
	return out

}
