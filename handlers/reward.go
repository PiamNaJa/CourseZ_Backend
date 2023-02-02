package handlers

import (
	"github.com/PiamNaJa/CourseZ_Backend/constants"
	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateRewardItem(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardItem *models.Reward_Item
		if err := c.BodyParser(&rewardItem); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := constants.Validate.Struct(rewardItem); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		if err := db.Model(&models.Reward_Item{}).Create(&rewardItem).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&rewardItem)
	}
}

func GetAllRewardItems(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardItems *[]models.Reward_Item
		if err := db.Model(&models.Reward_Item{}).Find(&rewardItems).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&rewardItems)
	}
}

func GetRewardItemById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardItem *models.Reward_Item
		if err := db.Model(&models.Reward_Item{}).First(&rewardItem, c.Params("item_id")).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&rewardItem)
	}
}

func DeleteRewardItemById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardItem *models.Reward_Item
		if err := db.Model(&models.Reward_Item{}).Delete(&rewardItem, c.Params("item_id")).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON("Deleted")
	}
}

func UpdateRewardItem(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardItem *models.Reward_Item
		var rewardItemInput *models.Reward_Item
		if err := c.BodyParser(&rewardItemInput); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := constants.Validate.Struct(rewardItemInput); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		if err := db.Model(&models.Reward_Item{}).First(&rewardItem, c.Params("item_id")).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		rewardItem.Item_name = rewardItemInput.Item_name
		rewardItem.Item_picture = rewardItemInput.Item_picture
		rewardItem.Item_title = rewardItemInput.Item_title
		rewardItem.Item_cost = rewardItemInput.Item_cost
		if err := db.Save(&rewardItem).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON("Update Success")
	}
}

func CreateRewardInfo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardInfo *models.Reward_Info
		if err := c.BodyParser(&rewardInfo); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if err := constants.Validate.Struct(rewardInfo); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if err := db.Model(&models.Reward_Info{}).Create(&rewardInfo).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(&rewardInfo)
	}
}

func GetAllRewardInfo(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardInfo *[]models.Reward_Info
		if err := db.Model(&models.Reward_Info{}).Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("fullname", "nickname", "user_id")
		}).Preload("Item").Find(&rewardInfo).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&rewardInfo)
	}
}

func GetRewardInfoByUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardInfo *[]models.Reward_Info
		var rewardItem []*models.Reward_Item
		if err := db.Model(&models.Reward_Info{}).Where("user_id = ?", c.Params("user_id")).Preload("Item").Find(&rewardInfo).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		for _, rI := range *rewardInfo {
			rewardItem = append(rewardItem, rI.Item)
		}
		return c.Status(fiber.StatusOK).JSON(&rewardItem)
	}
}

func GetRewardInfoByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardInfo *models.Reward_Info
		if err := db.Model(&models.Reward_Info{}).Where("reward_id", c.Params("reward_id")).Preload("User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("fullname", "nickname", "user_id")
		}).Preload("Item").Find(&rewardInfo).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&rewardInfo)
	}
}

func DeleteRewardInfoById(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var rewardInfo *models.Reward_Info
		if err := db.Model(&models.Reward_Info{}).Delete(&rewardInfo, c.Params("reward_id")).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON("Deleted")
	}
}