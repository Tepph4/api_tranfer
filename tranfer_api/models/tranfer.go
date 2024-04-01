package models
import(
	"time"
	
)
type Transfer struct {
	ID               string     `json:"id" binding:"required"`
	OrderNo          string     `json:"order_no" binding:"required"`
	Type             string     `json:"type" binding:"required"`
	Category         string     `json:"category" binding:"required"`
	DocumentTypeName string     `json:"document_type_name" binding:"required"`
	SellerNo         string     `json:"seller_no"`
	BuyerNo          string     `json:"buyer_no"`
	BuyerName        string     `json:"buyer_name"`
	SellerName       string     `json:"seller_name"`
	CategoryCash     string     `json:"category_cash"`
	Bank             string     `json:"bank"`
	CashTransfer     string     `json:"cash_transfer"`
	Cash             string     `json:"cash"`
	Discount         string     `json:"discount"`
	Retry            int        `json:"retry"`
	DocumentTypeID   string     `json:"document_type_id" binding:"required"`
	Status			 string		`json:""`	
	Created_date	 time.Time		`json:""`
	Updated_date	 time.Time		`json:""`
	SaleItems        []SaleItem `json:"sale_items" binding:"required"`
}

type SaleItem struct {
	ID          	string `json:"id" binding:"required"`
	TransfersID  	string `json:"transfer_id" binding:"required"`
	SKU         	string `json:"sku" binding:"required"`
	Amount      	int    `json:"amount" binding:"required"`
	Discount    	string    `json:"discount" binding:"required"`
	Free       		int    `json:"free" binding:"required"`
	Warehouse   	string `json:"warehouse" binding:"required"`
	ToWarehouse 	string `json:"to_warehouse" binding:"required"`
	Created_date	time.Time	`json:""`
}

