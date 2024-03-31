package models

type Transfer struct {
	ID               string     `json:"id"`
	OrderNo          string     `json:"order_no"`
	Type             string     `json:"type"`
	Category         string     `json:"category"`
	DocumentTypeName string     `json:"document_type_name"`
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
	DocumentTypeID   string     `json:"document_type_id"`
	Status			 string		`json:""`	
	Created_date	string		`json:""`
	Updated_date	string		`json:""`
	SaleItems        []SaleItem `json:"sale_items"`
}

type SaleItem struct {
	ID          string `json:"id"`
	TransfersID  string `json:"transfer_id"`
	SKU         string `json:"sku"`
	Amount      int    `json:"amount"`
	Discount    string    `json:"discount"`
	Free        int    `json:"free"`
	Warehouse   string `json:"warehouse"`
	ToWarehouse string `json:"to_warehouse"`
	Created_date	string		`json:""`
}

