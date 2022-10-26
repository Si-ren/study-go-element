package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"mysql/model"
)

const (
	insertResourceSQL = `INSERT INTO resource (
		id,vendor,region,zone,create_at,expire_at,category,type,instance_id,
		name,description,status,update_at,sync_at,sync_accout,public_ip,
		private_ip,pay_type,describe_hash,resource_hash
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	insertHostSQL = `INSERT INTO host (
		resource_id,cpu,memory,gpu_amount,gpu_spec,os_type,os_name,
		serial_number,image_id,internet_max_bandwidth_out,
		internet_max_bandwidth_in,key_pair_name,security_groups
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`
	updateResourceSQL = `UPDATE resource SET 
		expire_at=?,category=?,type=?,name=?,description=?,
		status=?,update_at=?,sync_at=?,sync_accout=?,
		public_ip=?,private_ip=?,pay_type=?,describe_hash=?,resource_hash=?
	WHERE id = ?`
	updateHostSQL = `UPDATE host SET 
		cpu=?,memory=?,gpu_amount=?,gpu_spec=?,os_type=?,os_name=?,
		image_id=?,internet_max_bandwidth_out=?,
		internet_max_bandwidth_in=?,key_pair_name=?,security_groups=?
	WHERE resource_id = ?`

	queryHostSQL      = `SELECT * FROM host `
	deleteHostSQL     = `DELETE FROM host WHERE resource_id = ?;`
	deleteResourceSQL = `DELETE FROM resource WHERE id = ?;`
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:13306)/cmdb?charset=utf8mb4")
	if err != nil {
		logrus.Error("open db failed : ", err)
		return
	}
	fmt.Println(queryHostSQL + " where resource_id > ?")
	//stmt, err := db.Prepare(queryHostSQL + " where id>?")
	stmt, err := db.Prepare(queryHostSQL + " where resource_id > ?")
	res, err := stmt.Query(5)

	if err != nil {
		logrus.Error("stmt query error")
		return
	}
	fmt.Println(res)
	hs := model.NewHostSet()
	h := model.NewHost()
	for res.Next() {
		res.Scan(&h.ResourceId, &h.CPU, &h.Memory, &h.GPUAmount, &h.GPUSpec, &h.OSType, &h.OSName, &h.SerialNumber)
		fmt.Println(h.Describe)
		hs.Add(h)
	}
	fmt.Println(hs.Items)

	stmt2, err := db.Prepare(deleteResourceSQL)
	stmt2.Exec(10)
}
