package controller

import (
	"database/sql"
	"fmt"
	"go_api/models"
	"log"
	"net/http"	
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gin-gonic/gin"
)

var (
	username = "mana"
	password = "manaa"
	dbname   = "rpa_dev"
	server   = "LAPTOP-OOENR37S"
	port     = 1433
)
var db *sql.DB

func Iconnect_toDB() {
	var err error
	fmt.Println("connecting to DB...")
	connString := fmt.Sprintf("server=%s;user id=%s; password=%s; port=%d; database=%s;", server, username, password, port, dbname)
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool : " + err.Error())
	}
	log.Printf("Connected!\n")

}

func CloseDB() {
	defer db.Close()
}

func ListTransfers(c *gin.Context) {
	fmt.Println("on List data")
	rows, err := db.Query("SELECT * FROM transfers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transfers []models.Transfer
	for rows.Next() {
		var transfer models.Transfer
		fmt.Println(rows)
		err := rows.Scan(
			&transfer.ID,
			&transfer.OrderNo,
			&transfer.Type,
			&transfer.Category,
			&transfer.DocumentTypeName,
			&transfer.SellerNo,
			&transfer.BuyerNo,
			&transfer.BuyerName,
			&transfer.SellerName,
			&transfer.CategoryCash,
			&transfer.Bank,
			&transfer.CashTransfer,
			&transfer.Cash,
			&transfer.Discount,
			&transfer.Retry,
			&transfer.Status,
			&transfer.DocumentTypeID,	
			&transfer.Created_date,
			&transfer.Updated_date,		
		)
		fmt.Printf("1")
		if err != nil {
			fmt.Println("2")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("3")
		// Fetch SaleItems for the current transfer
		saleItems, err := getSaleItems(transfer.ID)
		fmt.Printf("4")
		if err != nil {
			fmt.Println("5")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("6")
		transfer.SaleItems = saleItems
		transfers = append(transfers, transfer)
	}

	c.JSON(http.StatusOK, transfers)	
}

// getSaleItems fetches sale items for a given transfer ID
func getSaleItems(transferID string) ([]models.SaleItem, error) {
	fmt.Println("on get sale")
	rows, err := db.Query("SELECT * FROM sale_items WHERE transfer_id = @TransferID", sql.Named("TransferID", transferID))

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var saleItems []models.SaleItem
	for rows.Next() {
		var saleItem models.SaleItem
		err := rows.Scan(
			&saleItem.ID,
			&saleItem.TransfersID,
			&saleItem.SKU,
			&saleItem.Amount,
			&saleItem.Discount,
			&saleItem.Free,
			&saleItem.Warehouse,
			&saleItem.ToWarehouse,
			&saleItem.Created_date,
		)
		if err != nil {
			return nil, err
		}
		saleItems = append(saleItems, saleItem)
	}

	return saleItems, nil
}

func GetTransferBypennding(c *gin.Context) {
	fmt.Println("on List data")
	rows, err := db.Query("SELECT * FROM transfers where status = 'pending'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transfers []models.Transfer
	for rows.Next() {
		var transfer models.Transfer
		fmt.Println(rows)
		err := rows.Scan(
			&transfer.ID,
			&transfer.OrderNo,
			&transfer.Type,
			&transfer.Category,
			&transfer.DocumentTypeName,
			&transfer.SellerNo,
			&transfer.BuyerNo,
			&transfer.BuyerName,
			&transfer.SellerName,
			&transfer.CategoryCash,
			&transfer.Bank,
			&transfer.CashTransfer,
			&transfer.Cash,
			&transfer.Discount,
			&transfer.Retry,
			&transfer.Status,
			&transfer.DocumentTypeID,	
			&transfer.Created_date,
			&transfer.Updated_date,		
		)
		fmt.Printf("1")
		if err != nil {
			fmt.Println("2")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("3")
		// Fetch SaleItems for the current transfer
		saleItems, err := getSaleItems(transfer.ID)
		fmt.Printf("4")
		if err != nil {
			fmt.Println("5")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("6")
		transfer.SaleItems = saleItems
		transfers = append(transfers, transfer)
	}

	c.JSON(http.StatusOK, transfers)	
}

func GetTransferByOrderId(c *gin.Context, ) {
	fmt.Println("on get data by order id")
	orderID := c.Param("orderid")
	rows, err := db.Query("SELECT * FROM transfers WHERE order_no = @OrderNo", sql.Named("OrderNo", orderID))
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transfers []models.Transfer
	for rows.Next() {
		var transfer models.Transfer
		fmt.Println(rows)
		err := rows.Scan(
			&transfer.ID,
			&transfer.OrderNo,
			&transfer.Type,
			&transfer.Category,
			&transfer.DocumentTypeName,
			&transfer.SellerNo,
			&transfer.BuyerNo,
			&transfer.BuyerName,
			&transfer.SellerName,
			&transfer.CategoryCash,
			&transfer.Bank,
			&transfer.CashTransfer,
			&transfer.Cash,
			&transfer.Discount,
			&transfer.Retry,
			&transfer.Status,
			&transfer.DocumentTypeID,	
			&transfer.Created_date,
			&transfer.Updated_date,		
		)
		fmt.Printf("1")
		if err != nil {
			fmt.Println("2")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("3")
		// Fetch SaleItems for the current transfer
		saleItems, err := getSaleItems(transfer.ID)
		fmt.Printf("4")
		if err != nil {
			fmt.Println("5")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("6")
		transfer.SaleItems = saleItems
		transfers = append(transfers, transfer)
	}

	c.JSON(http.StatusOK, transfers)	
}

func GetTransferBybuyer_no(c *gin.Context, ) {
	fmt.Println("on get data by buyer no")
	buyerID := c.Param("buyerid")
	rows, err := db.Query("SELECT * FROM transfers WHERE buyer_no = @BuyerNo", sql.Named("BuyerNo", buyerID))	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transfers []models.Transfer
	for rows.Next() {
		var transfer models.Transfer
		fmt.Println(rows)
		err := rows.Scan(
			&transfer.ID,
			&transfer.OrderNo,
			&transfer.Type,
			&transfer.Category,
			&transfer.DocumentTypeName,
			&transfer.SellerNo,
			&transfer.BuyerNo,
			&transfer.BuyerName,
			&transfer.SellerName,
			&transfer.CategoryCash,
			&transfer.Bank,
			&transfer.CashTransfer,
			&transfer.Cash,
			&transfer.Discount,
			&transfer.Retry,
			&transfer.Status,
			&transfer.DocumentTypeID,	
			&transfer.Created_date,
			&transfer.Updated_date,		
		)
		fmt.Printf("1")
		if err != nil {
			fmt.Println("2")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("3")
		// Fetch SaleItems for the current transfer
		saleItems, err := getSaleItems(transfer.ID)
		fmt.Printf("4")
		if err != nil {
			fmt.Println("5")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("6")
		transfer.SaleItems = saleItems
		transfers = append(transfers, transfer)
	}

	c.JSON(http.StatusOK, transfers)	
}

func CreateTransfer(c *gin.Context) {
	fmt.Println("on create data transer")

	var transfer models.Transfer
	if err := c.BindJSON(&transfer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the transfer data into the database
	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tx.Commit()
	}()

	_, err = tx.Exec(
		"INSERT INTO transfers (id, order_no, type, category, document_type_name, seller_no, buyer_no, buyer_name, seller_name, category_cash, bank, cash_transfer, cash, discount, retry, document_type_id) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13, @p14, @p15, @p16)",
		transfer.ID,
		transfer.OrderNo,
		transfer.Type,
		transfer.Category,
		transfer.DocumentTypeName,
		transfer.SellerNo,
		transfer.BuyerNo,
		transfer.BuyerName,
		transfer.SellerName,
		transfer.CategoryCash,
		transfer.Bank,
		transfer.CashTransfer,
		transfer.Cash,
		transfer.Discount,
		transfer.Retry,
		transfer.DocumentTypeID,
	)
	if err != nil {
		return
	}

	// Insert sale items into the database
	for _, saleItem := range transfer.SaleItems {
		_, err := tx.Exec("INSERT INTO sale_items (id, transfer_id, sku, amount, discount, free, warehouse, to_warehouse) VALUES (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8)",
			saleItem.ID,
			saleItem.TransfersID,
			saleItem.SKU,
			saleItem.Amount,
			saleItem.Discount,
			saleItem.Free,
			saleItem.Warehouse,
			saleItem.ToWarehouse,
		)
		if err != nil {
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transfer created successfully"})
}
