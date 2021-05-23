package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/gofiber/fiber"
)

type fruitInterface struct {
	Page   int
	Fruits []string
}

// Home ...
func Home(c *fiber.Ctx) error {
	// Create html template from multiple files
	page, err := template.ParseFiles("html/pages/index.html", "html/components/header/index.html")
	if err != nil {
		return err
	}

	// Execute html template
	var template bytes.Buffer
	err = page.Execute(&template, nil)
	if err != nil {
		return err
	}

	// Send html template string
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(template.String())
}

// Routes ...
func Routes(c *fiber.Ctx) error {
	// Create html template from multiple files
	page, err := template.ParseFiles("html/pages/route/index.html", "html/components/header/index.html")
	if err != nil {
		return err
	}

	// Execute html template
	var template bytes.Buffer
	err = page.Execute(&template, nil)
	if err != nil {
		return err
	}

	// Send html template string
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return c.SendString(template.String())
}

// API ...
func API(c *fiber.Ctx) error {
	// Send API path
	path := fmt.Sprintf("API path: %s", c.Params("*"))
	c.Send([]byte(path))

	// Handle path if it is equal to "fruits"
	// **YOU COULD ALSO NOT USE A WILDCARD HERE, AND INSTEAD USE A PARAMETER**
	if c.Params("*") == "fruits" {

		// Response info
		response := fruitInterface{
			Page:   1,
			Fruits: []string{"apple", "peach", "pear"},
		}

		// Write response to json
		responseJSON, _ := json.Marshal(response)

		// Send response json
		return c.Send(responseJSON)
	}

	return nil
}
