package data

import (
	"strings"
)

func ClientFindByTitle(title string) (*Client, error) {
	var client Client
	result := DB.Where("lower(title) = ?", strings.ToLower(title)).Limit(1).Find(&client)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, nil
	}
	return &client, nil
}

func ClientCreate(title string, description *string) (*Client, error) {
	client := &Client{
		Title:       title,
		Description: description,
	}

	// Already in database?
	clientDB, err := ClientFindByTitle(title)
	if err == nil && clientDB != nil {
		return clientDB, err
	}

	// Create new and return
	tx := DB.Create(&client)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return client, nil
}
