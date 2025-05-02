package main

import (
	"backend/config"
	"backend/internal/db/models"
	"backend/pkg/db"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	config, err := config.NewParsedConfig()
	if err != nil {
		log.Fatalf("cannot get config, err:\n %+v", err)
	}

	dbConfig := db.DBConfig{
		Host:     config.Database.Host,
		Port:     config.Database.Port,
		Username: config.Database.Username,
		Password: config.Database.Password,
		Name:     config.Database.Name,
	}

	db, err := db.PgxConnect(ctx, dbConfig)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer db.Close(ctx)

	inputFile := flag.String("input", "example.json", "Path to input JSON file")
	flag.Parse()

	jsonData, err := os.ReadFile(*inputFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %+v", err)
	}

	// users := []models.User{}
	// if err = json.Unmarshal(jsonData, &users); err != nil {
	// 	log.Fatalf("Error parsing from json to struct: %+v", err)
	// }
	// products := []models.Product{}
	// if err = json.Unmarshal(jsonData, &products); err != nil {
	// 	log.Fatalf("Error parsing from json to struct: %+v", err)
	// }
	// sales := []models.Sale{}
	// if err = json.Unmarshal(jsonData, &sales); err != nil {
	// 	log.Fatalf("Error parsing from json to struct: %+v", err)
	// }
	// salesSummary := []models.SaleSummary{}
	// if err = json.Unmarshal(jsonData, &salesSummary); err != nil {
	// 	log.Fatalf("Error parsing from json to struct: %+v", err)
	// }
	// purchases := []models.Purchase{}
	// if err = json.Unmarshal(jsonData, &purchases); err != nil {
	// 	log.Fatalf("Error parsing from json to struct: %+v", err)
	// }
	// purchaseSummary := []models.PurchaseSummary{}
	// if err = json.Unmarshal(jsonData, &purchaseSummary); err != nil {
	// 	log.Fatalf("Error parsing from json to struct: %+v", err)
	// }
	// expenses := []models.Expense{}
	// if err = json.Unmarshal(jsonData, &expenses); err != nil {
	// 	log.Fatalf("Error parsing from json to struct: %+v", err)
	// }
	// expenseSummary := []models.ExpenseSummary{}
	// if err = json.Unmarshal(jsonData, &expenseSummary); err != nil {
	// 	log.Fatalf("Error parsing from json to struct: %+v", err)
	// }
	expenseByCategory := []models.ExpenseByCategory{}
	if err = json.Unmarshal(jsonData, &expenseByCategory); err != nil {
		log.Fatalf("Error parsing from json to struct: %+v", err)
	}

	//err = BulkInsertCopy(ctx, db, users)
	//err = BulkInsertCopy(ctx, db, "Users", users)
	//err = BulkInsertCopy(ctx, db, "Products", products)
	//err = BulkInsertCopy(ctx, db, "Sales", sales)
	//err = BulkInsertCopy(ctx, db, "SalesSummary", salesSummary)
	// err = BulkInsertCopy(ctx, db, "Purchases", purchases)
	// err = BulkInsertCopy(ctx, db, "PurchaseSummary", purchaseSummary)
	// err = BulkInsertCopy(ctx, db, "Expenses", expenses)
	// err = BulkInsertCopy(ctx, db, "ExpenseSummary", expenseSummary)
	err = BulkInsertCopy(ctx, db, "ExpenseByCategory", expenseByCategory)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	fmt.Println("Data insertion completed")
}

func BulkInsertCopy[T any](ctx context.Context, conn *pgx.Conn, tableName string, items []T) error {
	if len(items) == 0 {
		return nil
	}

	var columns []string //contain column names
	var rows [][]any     //contain data in each row

	t := reflect.TypeOf(items[0])

	//nếu t là pointer thì gán về struct để thao tác
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return errors.New("BulkInsertCopy only supports struct slices")
	}

	//get and add column name from tag db to columns
	for i := range t.NumField() {
		dbTag := t.Field(i).Tag.Get("db")
		if dbTag != "" && dbTag != "-" {
			columns = append(columns, dbTag)
		}
	}

	//get and add data in each row to rows
	for _, item := range items {
		val := reflect.ValueOf(item)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		var row []any
		for i := range t.NumField() {
			dbTag := t.Field(i).Tag.Get("db")
			if dbTag == "" || dbTag == "-" {
				continue
			}
			row = append(row, val.Field(i).Interface())
		}
		rows = append(rows, row)
	}

	_, err := conn.CopyFrom(
		ctx,
		pgx.Identifier{tableName},
		columns,
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return fmt.Errorf("bulk insert %s err:\n %+v", tableName, err)
	}

	return nil
}

// func BulkInsertCopy(ctx context.Context, conn *pgx.Conn, users []models.User) error {
// 	rows := make([][]any, len(users))
// 	for i, user := range users {
// 		rows[i] = []any{user.UserID, user.Name, user.Email}
// 	}

// 	_, err := conn.CopyFrom(
// 		ctx,
// 		pgx.Identifier{"Users"},
// 		[]string{"userId", "name", "email"},
// 		pgx.CopyFromRows(rows),
// 	)
// 	if err != nil {
// 		return fmt.Errorf("bulk insert err:\n %+v", err)
// 	}

// 	return nil
// }
