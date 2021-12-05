package repository

import (
	"database/sql"
	"example/web-service-gin/internal/business/domain"
	"example/web-service-gin/internal/business/gateway"
	"example/web-service-gin/internal/infraestructure/delivery/webapi/utils/apierrors"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySqlItemsRepository() gateway.ItemsRepository {

	dbUsername := "TBD"
	dbPassword := "TBD"
	dbHost := "TBD"
	dbName := "TBD"

	port := os.Getenv("PORT")
	if port == "" {
		dbUsername = "root"
		dbPassword = ""
		dbHost = "127.0.0.1"
		dbName = "web-service-gin"
	}
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUsername, dbPassword, dbHost, dbName)
	database, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}

	return &mysqlItemsRepository{
		database: database,
	}
}

type mysqlItemsRepository struct {
	database *sql.DB
}

func (repository *mysqlItemsRepository) AddItem(item *domain.Item) error {

	insert, err := repository.database.Prepare("INSERT INTO `items` (`title`, `price`, `date_created`) " +
		"VALUES (?, ?, ?)")
	if err != nil {
		return apierrors.NewInternalServerApiError("error preparing to save in db", err)
	}

	currentTime := time.Now()
	_, errExec := insert.Exec(item.Title, item.Price, currentTime.Format("2006-01-02 15:04:05"))

	if errExec != nil {
		return apierrors.NewInternalServerApiError("error saving item in db", err)
	}

	return nil
}

func (repository *mysqlItemsRepository) GetItemById(itemID string) (*domain.Item, error) {
	if itemID != "1234" {
		return nil, apierrors.NewNotFoundApiError(fmt.Sprintf("item_id %s was not found", itemID))
	}
	now := time.Now()
	return &domain.Item{
		ID:          1,
		Title:       "Harry Potter and the Philosopher’s Stone",
		Price:       102.34,
		DateCreated: &now,
	}, nil
}

func (repository *mysqlItemsRepository) GetItems() ([]*domain.Item, error) {

	registers, err := repository.database.Query("SELECT * FROM `items`")
	if err != nil {
		return nil, apierrors.NewInternalServerApiError("error getting items", err)
	}
	var items []*domain.Item
	for registers.Next() {
		var id int
		var title string
		var price float32
		var dateCreatedString, dateUpdatedString string
		errScanning := registers.Scan(&id, &title, &price, &dateCreatedString, &dateUpdatedString)
		if errScanning != nil {
			return nil, apierrors.NewInternalServerApiError("error scanning items in database", errScanning)
		}

		dateCreated, errDateCreated := time.Parse("2006-01-02 15:04:05", dateCreatedString)
		if errDateCreated != nil {
			return nil, apierrors.NewInternalServerApiError("error converting dateCreated", errDateCreated)
		}

		dateUpdated, errDateUpdated := time.Parse("2006-01-02 15:04:05", dateUpdatedString)
		if errDateUpdated != nil {
			return nil, apierrors.NewInternalServerApiError("error converting dateCreated", errDateUpdated)
		}

		var item = domain.Item{
			ID:          id,
			Title:       title,
			Price:       price,
			DateCreated: &dateCreated,
			DateUpdated: &dateUpdated,
		}
		items = append(items, &item)
	}

	return items, nil
}
