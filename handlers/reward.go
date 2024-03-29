package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/PiamNaJa/CourseZ_Backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateRewardItem(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardItem models.Reward_Item
		if err := c.BodyParser(&rewardItem); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := utils.Validate.Struct(rewardItem); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.Create(&rewardItem).Error; err != nil {
			return utils.Unexpected(err.Error())
		}

		return c.Status(fiber.StatusCreated).JSON(&rewardItem)
	}
}

func GetAllRewardItems(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardItems []models.Reward_Item
		db.Find(&rewardItems)

		return c.Status(fiber.StatusOK).JSON(&rewardItems)
	}
}

func GetRewardItemById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardItem models.Reward_Item
		if err := db.First(&rewardItem, c.Params("item_id")).Error; err != nil {
			return utils.HandleFindError(err)
		}

		return c.Status(fiber.StatusOK).JSON(&rewardItem)
	}
}

func DeleteRewardItemById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardItem models.Reward_Item
		if err := db.Delete(&rewardItem, c.Params("item_id")).Error; err != nil {
			return utils.HandleFindError(err)
		}

		return c.Status(fiber.StatusNoContent).JSON("Deleted")
	}
}

func UpdateRewardItem(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardItem models.Reward_Item
		var rewardItemInput models.Reward_Item
		if err := c.BodyParser(&rewardItemInput); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := utils.Validate.Struct(rewardItemInput); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := db.First(&rewardItem, c.Params("item_id")).Error; err != nil {
			return utils.HandleFindError(err)
		}

		rewardItem.Item_name = rewardItemInput.Item_name
		rewardItem.Item_picture = rewardItemInput.Item_picture
		rewardItem.Item_title = rewardItemInput.Item_title
		rewardItem.Item_cost = rewardItemInput.Item_cost
		if err := db.Save(&rewardItem).Error; err != nil {
			return utils.Unexpected(err.Error())
		}
		return c.Status(fiber.StatusCreated).JSON("Update Success")
	}
}

func CreateRewardInfo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("authorization")
		claims, err := utils.GetClaims(token)
		if err != nil {
			return utils.Unauthorized(err.Error())
		}
		var user models.User
		if err := db.Select("user_id", "point").First(&user, claims["user_id"]).Error; err != nil {
			return utils.HandleFindError(err)
		}

		var rewardInfo models.Reward_Info
		if err := c.BodyParser(&rewardInfo); err != nil {
			return utils.BadRequest(err.Error())
		}

		var rewardItem models.Reward_Item
		if err := db.First(&rewardItem, rewardInfo.ItemID).Error; err != nil {
			return utils.HandleFindError(err)
		}

		if rewardItem.Item_cost > user.Point {
			return utils.BadRequest("Not enough point")
		}
		user.Point -= rewardItem.Item_cost

		tx := db.Begin()

		if err := tx.Model(&user).Update("point", user.Point).Error; err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
		}

		rewardInfo.UserID = user.User_id

		if err := utils.Validate.Struct(rewardInfo); err != nil {
			return utils.BadRequest(err.Error())
		}

		if err := tx.Create(&rewardInfo).Error; err != nil {
			tx.Rollback()
			return utils.Unexpected(err.Error())
		}
		tx.Commit()
		return c.Status(fiber.StatusCreated).JSON(&rewardInfo)
	}
}

func GetAllRewardInfo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardInfo []models.Reward_Info
		if err := db.Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("fullname", "nickname", "user_id")
		}).Preload("Item").Find(&rewardInfo).Error; err != nil {
			return utils.HandleFindError(err)
		}

		return c.Status(fiber.StatusOK).JSON(&rewardInfo)
	}
}

func GetRewardInfoByUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardInfo []models.Reward_Info
		if err := db.Where("user_id = ?", c.Params("user_id")).Preload("Item").Find(&rewardInfo).Error; err != nil {
			return utils.HandleFindError(err)
		}
		if len(rewardInfo) == 0 {
			return utils.NotFound("Reward Item Not Found")
		}
		return c.Status(fiber.StatusOK).JSON(&rewardInfo)
	}
}

func GetRewardInfoByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardInfo models.Reward_Info
		if err := db.Where("reward_id", c.Params("reward_id")).Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("fullname", "nickname", "user_id")
		}).Preload("Item").Find(&rewardInfo).Error; err != nil {
			return utils.HandleFindError(err)
		}
		return c.Status(fiber.StatusOK).JSON(&rewardInfo)
	}
}

func DeleteRewardInfoById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardInfo models.Reward_Info
		if err := db.Delete(&rewardInfo, c.Params("reward_id")).Error; err != nil {
			return utils.HandleFindError(err)
		}

		return c.Status(fiber.StatusNoContent).JSON("Deleted")
	}
}
