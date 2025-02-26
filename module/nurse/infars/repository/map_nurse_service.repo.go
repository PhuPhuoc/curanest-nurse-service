package nurserepository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (repo *nurseRepo) MapNurseService(ctx context.Context, nurseID uuid.UUID, serviceIDs []uuid.UUID) (err error) {
	currentServiceIDs, err := repo.GetNurseService(ctx, nurseID)
	if err != nil {
		return fmt.Errorf("cannot get current services for nurse %v: %w", nurseID, err)
	}

	servicesToAdd := differenceUUIDs(serviceIDs, currentServiceIDs)
	servicesToRemove := differenceUUIDs(currentServiceIDs, serviceIDs)

	tx, err := repo.db.Beginx()
	if err != nil {
		return fmt.Errorf("cannot start transaction for mapping nurse services: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("cannot commit transaction: %w", commitErr)
		}
	}()

	insertQuery := `INSERT INTO nursing_service (nursing_id, service_id) VALUES (?, ?)`
	for _, serviceID := range servicesToAdd {
		if _, err = tx.Exec(insertQuery, nurseID, serviceID); err != nil {
			return fmt.Errorf("failed to add record for nurse %v, service %v: %w", nurseID, serviceID, err)
		}
	}

	deleteQuery := `DELETE FROM nursing_service WHERE nursing_id = ? AND service_id = ?`
	for _, serviceID := range servicesToRemove {
		if _, err = tx.Exec(deleteQuery, nurseID, serviceID); err != nil {
			return fmt.Errorf("failed to delete record for nurse %v, service %v: %w", nurseID, serviceID, err)
		}
	}

	return nil
}

func differenceUUIDs(base, exclude []uuid.UUID) []uuid.UUID {
	var diff []uuid.UUID
	for _, id := range base {
		exist := false
		for _, exID := range exclude {
			if id == exID {
				exist = true
				break
			}
		}
		if !exist {
			diff = append(diff, id)
		}
	}
	return diff
}
