package models

import (
    "errors"
    "sync"
    "time"
)

type Note struct {
    ID        int       `json:"id"`
    Content   string    `json:"content" binding:"required"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

var (
    notes  []Note
    nextID int
    mutex  sync.Mutex
)

func GetAllNotes() []Note {
    mutex.Lock()
    defer mutex.Unlock()
    return notes
}

func AddNote(content string) (Note, error) {
    if content == "" {
        return Note{}, errors.New("content cannot be empty")
    }

    mutex.Lock()
    defer mutex.Unlock()
    note := Note{
        ID:        nextID,
        Content:   content,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }
    notes = append(notes, note)
    nextID++
    return note, nil
}

func UpdateNote(id int, content string) (bool, error) {
    if content == "" {
        return false, errors.New("content cannot be empty")
    }

    mutex.Lock()
    defer mutex.Unlock()
    for i, note := range notes {
        if note.ID == id {
            notes[i].Content = content
            notes[i].UpdatedAt = time.Now()
            return true, nil
        }
    }
    return false, nil
}

func GetNote(id int) (Note, bool) {
    mutex.Lock()
    defer mutex.Unlock()
    for _, note := range notes {
        if note.ID == id {
            return note, true
        }
    }
    return Note{}, false
}

func DeleteNote(id int) bool {
    mutex.Lock()
    defer mutex.Unlock()
    for i, note := range notes {
        if note.ID == id {
            notes = append(notes[:i], notes[i+1:]...)
            return true
        }
    }
    return false
}