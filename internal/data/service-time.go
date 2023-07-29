package data

import (
	"fmt"
	"strings"
	"time"
)

func TimeFind(id string) (*Time, error) {
	var time Time
	result := DB.
		Preload("Client").
		Preload("Project").
		Where("lower(id) = ?", strings.ToLower(id)).Limit(1).
		Find(&time)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, nil
	}

	return &time, nil
}

func TimeFindRunning() (*Time, error) {
	var time Time
	result := DB.
		Preload("Client").
		Preload("Project").
		Where("stopped_at IS NULL").Limit(1).
		Find(&time)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected < 1 {
		return nil, nil
	}

	return &time, nil
}

func TimeFindBetween(startDate time.Time, stopDate time.Time) ([]Time, error) {
	var times []Time
	result := DB.
		Preload("Client").
		Preload("Project").
		Where("started_at BETWEEN ? AND ?", startDate, stopDate).Order("started_at ASC").
		Find(&times)

	if result.Error != nil {
		return nil, result.Error
	}

	return times, nil
}

func TimeDelete(id string) error {
	time, err := TimeFind(id)
	if err != nil {
		return err
	}

	DB.Where("lower(id) = ?", strings.ToLower(time.ID)).Delete(&Time{})
	return nil
}

func TimeSearch(client string, project string, code string, comment string, startDate time.Time, stopDate time.Time) ([]Time, error) {
	var times []Time
	query := DB.InnerJoins("Client").InnerJoins("Project")

	// Client
	if len(client) > 0 {
		query.Where("Client.title LIKE ?", fmt.Sprintf("%%%s%%", client))
	}

	// Project
	if len(project) > 0 {
		query.Where("Project.title LIKE ?", fmt.Sprintf("%%%s%%", project))
	}

	// Code
	if len(code) > 0 {
		query.Where("code LIKE ?", fmt.Sprintf("%%%s%%", code))
	}

	// Comment
	if len(comment) > 0 {
		query.Where("comment LIKE ?", fmt.Sprintf("%%%s%%", comment))
	}

	// Dates
	query.Where("started_at BETWEEN ? AND ?", startDate, stopDate).Order("started_at ASC")

	result := query.Find(&times)
	if result.Error != nil {
		return nil, result.Error
	}

	return times, nil
}

func TimeCreate(clientID string, projectID string, startedAt time.Time, stoppedAt *time.Time, code *string, comment *string) (*Time, error) {
	time := &Time{
		ClientID:  clientID,
		ProjectID: projectID,
		StartedAt: &startedAt,
		StoppedAt: stoppedAt,
		Code:      code,
		Comment:   comment,
	}

	// Create new and return
	tx := DB.Create(&time)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return time, nil
}

func TimeUpdate(id string, clientID string, projectID string, startedAt time.Time, stoppedAt *time.Time, code *string, comment *string) (*Time, error) {
	time, err := TimeFind(id)
	if err != nil {
		return nil, err
	}
	if time == nil {
		return nil, nil
	}

	time.ClientID = clientID
	time.ProjectID = projectID
	time.StartedAt = &startedAt
	time.StoppedAt = stoppedAt
	time.Code = code
	time.Comment = comment

	// Update and return
	tx := DB.Omit("Client", "Project").Save(&time)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return time, nil
}

// Stop all runnings timers
func TimeStopAll(stoppedAt time.Time) error {
	t := stoppedAt.Format("2006-01-02 15:04:05+00:00")
	tx := DB.Model(&Time{}).Where("stopped_at IS NULL").Update("stopped_at", t)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
