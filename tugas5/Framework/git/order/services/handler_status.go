package services

import (
	"context"
	"database/sql"
	"fmt"

	cm "tugas5/Framework/git/order/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) StatusHandler(ctx context.Context, req cm.StatusRequest) (res cm.StatusResponse) {
	defer panicRecovery()

	var db *sql.DB
	var err error

	host := cm.Config.Connection.Host
	port := cm.Config.Connection.Port
	user := cm.Config.Connection.User
	pass := cm.Config.Connection.Password
	data := cm.Config.Connection.Database

	var mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, data)

	db, err = sql.Open("mysql", mySQL)

	if err != nil {
		panic(err.Error())
	}

	res.TrxID = req.TrxID

	var status cm.StatusResponse

	sql := `SELECT
				
				IFNULL(merchant,''),
				IFNULL(bill_no,'') bill_no,
				IFFNULL(payment_reff,'') payment_reff,
				IFNULL(payment_date,'') payment_date,
				IFNULL(payment_status_date,'') payment_status_date,
				IFNULL(payment_status_desc,'') payment_status_desc,
				IFNULL(response_date,'') response_date,
				IFNULL(response_desc,'') response_desc
				FROM status_pembayaran WHERE trx_id = ?`

	result, err := db.Query(sql, req.TrxID)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		err := result.Scan(&status.Response, &status.MerchantID, &status.Merchant, &status.BillNO, &status.PaymentReff, &status.PaymentDate, &status.PaymentStatusCode, &status.PaymentStatusDesc, &status.ResponseCode, &status.ResponseDesc)

		if err != nil {
			panic(err.Error())
		}
	}

	status.Response = "Response"
	status.TrxID = req.TrxID
	status.MerchantID = req.MerchantID
	status.BillNO = req.BillNO
	status.Response = res.Response
	res = status
	return
}
