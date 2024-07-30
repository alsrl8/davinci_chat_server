package handlers

import (
	"davinci-chat/consts"
	"davinci-chat/types"
	"davinci-chat/utils"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func GetUserEmailsByNameHandler(c *fiber.Ctx) error {
	userName := c.Query("name")

	var emails []string
	for conn, valid := range connections {
		if !valid {
			continue
		} else if connUserMap[conn].UserName != userName {
			continue
		}

		isGuest, err := utils.GetIsGuest(conn)
		if err != nil || isGuest {
			continue
		}

		emails = append(emails, connUserMap[conn].UserEmail)
	}

	if len(emails) == 0 {
		return c.JSON(fiber.Map{"emails": []string{}})
	}

	return c.JSON(fiber.Map{"emails": emails})
}

func SendInvitation(c *fiber.Ctx) error {
	var sendInvitationRequest types.SendInvitationRequest
	if err := c.BodyParser(&sendInvitationRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Failed to parse request"})
	}

	conn, has := userEmailConnMap[sendInvitationRequest.UserEmail]
	if !has {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "there is no user with such email now"})
	}

	// TODO Invitation Message 고도화
	msg := types.Message{
		User:        "ADMIN",
		Message:     "You got an invitation",
		Time:        "",
		MessageType: 0,
		UserType:    consts.Admin,
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to marshal message"})
	}

	err = conn.WriteMessage(int(consts.TextMessage), jsonData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "failed to send invitation"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "success"})
}
