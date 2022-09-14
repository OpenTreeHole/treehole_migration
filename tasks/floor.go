package tasks

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	. "treehole_migration/models"
)

func modifyFloor() {
	offset := 0
	size := 1000

	err := DB.Transaction(func(tx *gorm.DB) error {
		for {
			floors := make([]Floor, 0)
			result := tx.Table("floor").
				Where("id > ? and (`like` != 0 or history->>'$' <> '[]')", offset).
				Limit(size).Find(&floors)
			if result.Error != nil {
				return result.Error
			}

			if len(floors) == 0 {
				break
			}

			floorLikes := make([]FloorLike, 0)
			floorHistory := make([]FloorHistory, 0)
			for _, floor := range floors {
				var err error
				if floor.Like != 0 {
					likeData := make([]int, 0)
					err = json.Unmarshal([]byte(floor.LikeData), &likeData)
					if err != nil {
						return err
					}
					for _, userID := range likeData {
						floorLikes = append(floorLikes, FloorLike{
							FloorID:  floor.ID,
							UserID:   userID,
							LikeData: 1,
						})
					}
				}

				if floor.History != "[]" {
					historyData := make([]FloorHistoryOld, 0)
					err = json.Unmarshal([]byte(floor.History), &historyData)
					if err != nil {
						return err
					}

					for _, history := range historyData {
						floorHistory = append(floorHistory, FloorHistory{
							BaseModel: BaseModel{
								CreatedAt: history.AlteredTime,
								UpdatedAt: history.AlteredTime,
							},
							Content: history.Content,
							Reason:  "",
							FloorID: floor.ID,
							UserID:  history.AlteredBy,
						})
					}
				}
			}

			result = tx.Clauses(
				clause.OnConflict{
					DoNothing: true,
				}).Create(&floorLikes)
			if result.Error != nil {
				return result.Error
			}

			result = tx.Clauses(
				clause.OnConflict{
					DoNothing: true,
				}).Create(&floorHistory)
			if result.Error != nil {
				return result.Error
			}

			offset = floors[len(floors)-1].ID

			fmt.Printf("floor: (%v - %v)\n", floors[0].ID, offset)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
}
