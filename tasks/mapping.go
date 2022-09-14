package tasks

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	. "treehole_migration/models"
)

func modifyMapping() {
	offset := 0
	size := 1000

	err := DB.Transaction(func(tx *gorm.DB) error {
		for {
			holes := make([]Hole, 0)
			result := tx.Table("hole").
				Where("id > ? and mapping->>'$' != '{}'", offset).
				Limit(size).Find(&holes)
			if result.Error != nil {
				return result.Error
			}
			if len(holes) == 0 {
				break
			}

			anonynameMapping := make([]AnonynameMapping, 0)
			for _, hole := range holes {
				mappingOld := make(map[string]string)
				err := json.Unmarshal([]byte(hole.Mapping), &mappingOld)
				if err != nil {
					return result.Error
				}

				for k, v := range mappingOld {
					userID, err := strconv.Atoi(k)
					if err != nil {
						return err
					}
					anonynameMapping = append(anonynameMapping, AnonynameMapping{
						HoleID:    hole.ID,
						UserID:    userID,
						Anonyname: v,
					})
				}
			}

			result = tx.Clauses(
				clause.OnConflict{
					DoNothing: true,
				}).Create(&anonynameMapping)
			if result.Error != nil {
				return result.Error
			}

			offset = holes[len(holes)-1].ID

			fmt.Printf("hole: (%v - %v)\n", holes[0].ID, offset)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}
