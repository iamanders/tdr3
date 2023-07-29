package data

import (
	"strings"
)

func ProjectFindByTitle(title string) (*Project, error) {
	var project Project
	result := DB.Where("lower(title) = ?", strings.ToLower(title)).Limit(1).Find(&project)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, nil
	}
	return &project, nil
}

func ProjectCreate(title string, description *string) (*Project, error) {
	project := &Project{
		Title:       title,
		Description: description,
	}

	// Already in database?
	projectDB, err := ProjectFindByTitle(title)
	if err == nil && projectDB != nil {
		return projectDB, err
	}

	// Create new and return
	tx := DB.Create(&project)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return project, nil
}
