package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"notes/pkg/models"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

var notes []models.Note
var lastId int

func saveNotesToFile() error {
	file, err := os.Create("./pkg/notes.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(notes)
	if err != nil {
		return err
	}

	return nil
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var newNote models.Note
	err := json.NewDecoder(r.Body).Decode(&newNote)
	if err != nil {
		log.Println(err)
		return
	}

	lastId++
	newNote.Id = lastId
	notes = append(notes, newNote)

	err = saveNotesToFile()
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newNote)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteId := vars["id"]

	id, err := strconv.Atoi(noteId)
	if err != nil {
		log.Println(err)
		return
	}

	var rangeId models.Note
	for _, note := range notes {
		if note.Id == id {
			rangeId = note
			break
		}
	}
	json.NewEncoder(w).Encode(rangeId)
}

func EditNotes(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	noteIdStr := vars["id"]

	noteId, err := strconv.Atoi(noteIdStr)
	if err != nil {
		log.Println(err)
		return
	}

	var editNote models.Note
	err = json.NewDecoder(r.Body).Decode(&editNote)
	if err != nil {
		log.Println(err)
		return
	}

	var rangeNote *models.Note
	for i := range notes {
		if notes[i].Id == noteId {
			rangeNote = &notes[i]
			break
		}
	}

	rangeNote.Content = editNote.Content
	rangeNote.CreatedData = editNote.CreatedData

	saveNotesToFile()

	json.NewEncoder(w).Encode(rangeNote)
}

func DeleteNotes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteIdStr := vars["id"]

	noteId, err := strconv.Atoi(noteIdStr)
	if err != nil {
		log.Println(err)
		return
	}

	var rangeIndex = -1
	for i, note := range notes {
		if note.Id == noteId {
			rangeIndex = i
			break
		}
	}

	notes = append(notes[:rangeIndex], notes[rangeIndex+1:]...)

	saveNotesToFile()
}

func GetAllNotes(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(notes)
}
