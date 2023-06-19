package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Note struct {
	ID      int
	Title   string
	Content string
}

var notes = []Note{
	{ID: 1, Title: "Note 1", Content: "Content of Note 1"},
	{ID: 2, Title: "Note 2", Content: "Content of Note 2"},
	{ID: 3, Title: "Note 3", Content: "Content of Note 3"},
}

func main() {
	app := fiber.New()
	app.Get("/notes", getNotes)
	app.Get("/notes/:id", getNoteByID)
	app.Post("/notes", createNote)
	app.Put("/notes/:id", updateNote)
	app.Delete("/notes/:id", deleteNote)

	fmt.Println("Listining on Port 3000")
	log.Fatal(app.Listen(":3000"))
}

func getNotes(c *fiber.Ctx) error {
	return c.JSON(notes)
}

func getNoteByID(c *fiber.Ctx) error {
	idParams := c.Params("id")
	id, err := strconv.Atoi(idParams)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	for _, note := range notes {
		if note.ID == id {
			return c.JSON(note)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
}

func createNote(c *fiber.Ctx) error {
	note := new(Note)

	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	note.ID = len(notes) + 1
	notes = append(notes, *note)

	return c.Status(fiber.StatusCreated).JSON(note)
}

func updateNote(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	note := new(Note)

	if err := c.BodyParser(note); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request Payload"})
	}

	for i := range notes {
		if notes[i].ID == id {
			notes[i].Title = note.Title
			notes[i].Content = note.Content
			return c.Status(fiber.StatusOK).JSON(notes[i])
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
}

func deleteNote(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid note ID"})
	}

	for i := range notes {
		if notes[i].ID == id {
			notes = append(notes[:i], notes[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Note not found"})
}
