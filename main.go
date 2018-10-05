package main

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

var mySql *sqlx.DB
var sqlConn *sqlx.DB

func init() {
	sqlConn = ConnectSQL()
	mySql = ConnectMySql()
}

func ConnectSQL() (msdb *sqlx.DB) {
	db_host := "192.168.0.7"
	db_name := "bcnp"
	db_user := "vbuser"
	db_pass := "132"
	port := "1433"

	dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", db_host, db_user, db_pass, port, db_name)
	msdb = sqlx.MustConnect("mssql", dsn)
	if msdb.Ping() != nil {
		fmt.Println("Error")
	}

	fmt.Println("msdb = ", msdb.DriverName())
	return msdb
}

func ConnectMySql() (mydb *sqlx.DB) {
	dsn := "root:[ibdkifu88@tcp(nopadol.net:3306)/" + "npdl" + "?parseTime=true&charset=utf8&loc=Local" //DriveThru Server
	mydb = sqlx.MustConnect("mysql", dsn)
	if (mydb.Ping() != nil) {
		fmt.Println("Error")
	}
	fmt.Println("mysql = ", mydb.DriverName(), "dsn = ", dsn)
	return mydb
}

type Item struct {
	code         string  `json:"code" db:"code"`
	itemName     string  `json:"item_name" db:"item_name"`
	barCode      string  `json:"bar_code" db:"bar_code"`
	shortName    string  `json:"short_name" db:"short_name"`
	unitCode     string  `json:"unit_code" db:"unit_code"`
	price        float64 `json:"price" db:"price"`
	buyUnit      string  `json:"buy_unit" db:"buy_unit"`
	stockType    int64   `json:"stock_type" db:"stock_type"`
	itemStatus   int64   `json:"item_status" db:"item_status"`
	stockQty     float64 `json:"stock_qty" db:"stock_qty"`
	averageCost  float64 `json:"average_cost" db:"average_cost"`
	activeStatus int64   `json:"active_status" db:"active_status"`
	createBy     string  `json:"create_by" db:"create_by"`
	createTime   string  `json:"create_time" db:"create_time"`
	editBy       string  `json:"edit_by" db:"edit_by"`
	editTime     string  `json:"edit_time" db:"edit_time"`
}

type BCItem struct {
	Code            string  `json:"code" db:"Code"`
	Name1           string  `json:"name" db:"Name1"`
	BarCode         string  `json:"bar_code" db:"BarCode"`
	ShortName       string  `json:"short_name" db:"ShortName"`
	DefSaleUnitCode string  `json:"def_sale_unit_code" db:"DefSaleUnitCode"`
	SalePrice1      float64 `json:"sale_price_1" db:"SalePrice1"`
	DefBuyUnitCode  string  `json:"def_buy_unit_code" db:"DefBuyUnitCode"`
	StockType       int64   `json:"stock_type" db:"StockType"`
	ItemStatus      int64   `json:"item_status" db:"ItemStatus"`
	StockQty        float64 `json:"stock_qty" db:"StockQty"`
	AverageCost     float64 `json:"average_cost" db:"AverageCost"`
	ActiveStatus    int64   `json:"active_status" db:"ActiveStatus"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
}

type Customer struct {
	code         string `json:"code" db:"code"`
	name         string `json:"name" db:"name"`
	address      string `json:"address" db:"address"`
	telephone    string `json:"telephone" db:"telephone"`
	fax          string `json:"fax" db:"fax"`
	tax_no       string `json:"tax_no" db:"tax_no"`
	activeStatus int64  `json:"active_status" db:"active_status"`
	createBy     string `json:"create_by" db:"create_by"`
	createTime   string `json:"create_time" db:"create_time"`
	editBy       string `json:"edit_by" db:"edit_by"`
	editTime     string `json:"edit_time" db:"edit_time"`
}

type BCCustomer struct {
	Code           string `json:"code" db:"Code"`
	Name1          string `json:"name1" db:"Name1"`
	BillAddress    string `json:"bill_address" db:"BillAddress"`
	Telephone      string `json:"telephone" db:"Telephone"`
	Fax            string `json:"fax" db:"Fax"`
	TaxNo          string `json:"tax_no" db:"TaxNo"`
	ActiveStatus   int64  `json:"active_status" db:"ActiveStatus"`
	CreatorCode    string `json:"create_code" db:"CreatorCode"`
	CreateDateTime string `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode string `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT  string `json:"last_edit_date_t" db:"LastEditDateT"`
}

type Vendor struct {
	code         string `json:"code" db:"code"`
	name         string `json:"name" db:"name"`
	address      string `json:"address" db:"address"`
	telephone    string `json:"telephone" db:"telephone"`
	fax          string `json:"fax" db:"fax"`
	activeStatus int64  `json:"active_status" db:"active_status"`
	createBy     string `json:"create_by" db:"create_by"`
	createTime   string `json:"create_time" db:"create_time"`
	editBy       string `json:"edit_by" db:"edit_by"`
	editTime     string `json:"edit_time" db:"edit_time"`
}

type BCSupplier struct {
	Code           string `json:"code" db:"Code"`
	Name1          string `json:"name1" db:"Name1"`
	Address        string `json:"address" db:"Address"`
	Telephone      string `json:"telephone" db:"Telephone"`
	Fax            string `json:"fax" db:"Fax"`
	ActiveStatus   int64  `json:"active_status" db:"ActiveStatus"`
	CreatorCode    string `json:"create_code" db:"CreatorCode"`
	CreateDateTime string `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode string `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT  string `json:"last_edit_date_t" db:"LastEditDateT"`
}

type BCBarcode struct {
	ItemCode     string `json:"item_code" db:"ItemCode"`
	BarCode      string `json:"bar_code" db:"BarCode"`
	UnitCode     string `json:"unit_code" db:"UnitCode"`
	ActiveStatus int64  `json:"active_status" db:"ActiveStatus"`
}

type BCPrice struct {
	ItemCode     string  `json:"item_code" db:"ItemCode"`
	BarCode      string  `json:"bar_code" db:"BarCode"`
	UnitCode     string  `json:"unit_code" db:"UnitCode"`
	SalePrice1   float64 `json:"sale_price_1" db:"SalePrice1"`
	SalePrice2   float64 `json:"sale_price_2" db:"SalePrice2"`
}

func main() {
	//var item_exist int
	//var cust_exist int
	//var vendor_exist int
	//var bar_exist int
	var price_exist int

	//items := []BCItem{}
	//custs := []BCCustomer{}
	//vendors := []BCSupplier{}
	//bars := []BCBarcode{}
	prices := []BCPrice{}

	//sql := `set dateformat dmy     select Code, Name1, isnull(ItemBarCode,'') as BarCode, isnull(ShortName,'') as ShortName, isnull(DefSaleUnitCode,'') as DefSaleUnitCode, isnull(SalePrice1,0) as SalePrice1, isnull(DefBuyUnitCode,'') as DefBuyUnitCode, isnull(StockType,0) as StockType, isnull(ItemStatus,0) as ItemStatus, isnull(StockQty,0) as StockQty, isnull(AverageCost,0) as AverageCost, ActiveStatus, isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT from dbo.BCItem where activestatus = 1 order by code`
	//err := sqlConn.Select(&items, sql)
	//if err != nil {
	//	fmt.Println("error from sale= ", err)
	//}
	//var number1 int
	//for _, list := range items {
	// number1 = number1+1
	//	fmt.Println(strconv.Itoa(number1) +"."+list.Code)
	//	sql := `select count(code) as vCount from Item where code = ?`
	//	err := mySql.Get(&item_exist, sql, list.Code)
	//	if err != nil {
	//		fmt.Println("error count = ", err)
	//	}
	//
	//	if (item_exist == 0) {
	//		sql := `insert into Item(code,item_name,bar_code,short_name,unit_code,price,buy_unit,stock_type,item_status,stock_qty,average_cost,active_status,create_by,create_time,edit_by,edit_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	//		_, err := mySql.Exec(sql, list.Code, list.Name1, list.BarCode, list.ShortName, list.DefSaleUnitCode, list.SalePrice1, list.DefBuyUnitCode, list.StockType, list.ItemStatus, list.StockQty, list.AverageCost, list.ActiveStatus, list.CreatorCode, list.CreateDateTime, list.LastEditorCode, list.LastEditDateT)
	//		if err != nil {
	//			fmt.Println("error from user = ", err)
	//		}
	//	}
	//}

	//sql_cust := `set dateformat dmy     select Code, Name1, isnull(BillAddress,'') as BillAddress, isnull(Telephone,'') as Telephone, isnull(Fax,'') as Fax, isnull(TaxNo,'') as TaxNo, ActiveStatus, isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT from dbo.BCAR where activestatus = 1 order by code`
	//err := sqlConn.Select(&custs, sql_cust)
	//if err != nil {
	//	fmt.Println("error from customer= ", err)
	//}
	//
	//var number2 int
	//
	//for _, list := range custs {
	//	number2 = number2+1
	//	fmt.Println(strconv.Itoa(number2) +"."+list.Code)
	//	sql := `select count(code) as vCount from Customer where code = ?`
	//	err := mySql.Get(&cust_exist, sql, list.Code)
	//	if err != nil {
	//		fmt.Println("error count = ", err)
	//	}
	//
	//	if (cust_exist == 0) {
	//		sql := `insert into Customer(code,name,address,telephone,fax,tax_no,active_status,create_by,create_time,edit_by,edit_time) values(?,?,?,?,?,?,?,?,?,?,?)`
	//		_, err := mySql.Exec(sql, list.Code,list.Name1,list.BillAddress,list.Telephone,list.Fax,list.TaxNo, list.ActiveStatus, list.CreatorCode, list.CreateDateTime, list.LastEditorCode, list.LastEditDateT)
	//		if err != nil {
	//			fmt.Println("error from user = ", err)
	//		}
	//	}
	//}

	//sql_vendor := `set dateformat dmy     select Code, Name1, isnull(Address,'') as Address, isnull(Telephone,'') as Telephone, isnull(Fax,'') as Fax, ActiveStatus, isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT from dbo.BCAP where activestatus = 1 order by code`
	//err := sqlConn.Select(&vendors, sql_vendor)
	//if err != nil {
	//	fmt.Println("error from customer= ", err)
	//}
	//
	//var number3 int
	//
	//for _, list := range vendors {
	//	number3 = number3 + 1
	//	fmt.Println(strconv.Itoa(number3) + "." + list.Code)
	//	sql := `select count(code) as vCount from Customer where code = ?`
	//	err := mySql.Get(&vendor_exist, sql, list.Code)
	//	if err != nil {
	//		fmt.Println("error count = ", err)
	//	}
	//
	//	if (vendor_exist == 0) {
	//		sql := `insert into Vendor(code,name,address,telephone,fax,active_status,create_by,create_time,edit_by,edit_time) values(?,?,?,?,?,?,?,?,?,?)`
	//		_, err := mySql.Exec(sql, list.Code, list.Name1, list.Address, list.Telephone, list.Fax, list.ActiveStatus, list.CreatorCode, list.CreateDateTime, list.LastEditorCode, list.LastEditDateT)
	//		if err != nil {
	//			fmt.Println("error from user = ", err)
	//		}
	//	}
	//}

	sql_price := `set dateformat dmy     select ItemCode, UnitCode, SalePrice1, SalePrice2 from dbo.BCPriceList where itemcode in (select code from dbo.BCItem where activestatus = 1) and saletype = 0 and transporttype = 0 order by itemcode`
	err := sqlConn.Select(&prices, sql_price)
	if err != nil {
		fmt.Println("error = ", err)
	}

	var number4 int

	for _, list := range prices {
		number4 = number4 + 1
		fmt.Println(strconv.Itoa(number4) + "." + list.ItemCode)
		sql := `select count(item_code) as vCount from Price where item_code = ? and unit_code = ?`
		err := mySql.Get(&price_exist, sql, list.ItemCode, list.UnitCode)
		if err != nil {
			fmt.Println("error count = ", err)
		}

		if (price_exist == 0) {
			sql := `insert into Price(item_code,unit_code,sale_price_1,sale_price_2) values(?,?,?,?)`
			_, err := mySql.Exec(sql, list.ItemCode, list.UnitCode, list.SalePrice1, list.SalePrice2)
			if err != nil {
				fmt.Println("error from user = ", err)
			}
		}
	}
}
