package models

import (
	"time"

	"example.com/eventico/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserID      int64     `json:"user_id"`
}

func (e* Event) Save() error {
  query := `
  INSERT INTO events(name, description, location, dateTime, user_id)
  VALUES (?, ?, ?, ?, ?)`
  stmt, err := db.DB.Prepare(query)
  if err != nil {
      return err
  }

  defer stmt.Close()
  result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
  if err != nil {
      return err
  }

  id, err := result.LastInsertId()
  e.ID = id
  return err
}

func GetAllEvents() ([]Event, error) {
  query := `SELECT * FROM events`
  rows, err  := db.DB.Query(query)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  events := []Event{}
  for rows.Next() {
    var e Event
    err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

    if err != nil { 
      return nil, err
    }

    events = append(events, e)
  }
  return events, nil
}

func GetEventById(id int64) (Event, error) {
  query := `SELECT * FROM events WHERE id = ?`
  row := db.DB.QueryRow(query, id)
  var e Event
  err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

  if err != nil {
    return e, err
  }
  
  return e, nil
}

func (e Event) Update() error {
  query := `
  UPDATE events 
  SET name = ?, description = ?, location = ?, dateTime = ? 
  WHERE id = ?
  `
  stmt, err := db.DB.Prepare(query)
  if err != nil {
    return err
  }
  
  defer stmt.Close()

  _, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)

  if err != nil {
    return err
  }

  return  nil
}

func (e Event) Delete() error {
  query := `DELETE FROM events WHERE id = ?`
  stmt, err := db.DB.Prepare(query)
  if err != nil {
    return err
  }

  defer stmt.Close()

  _, err = stmt.Exec(e.ID)

  if err != nil {
    return err
  }

  return nil
}

func (e Event) Register(userId int64) error {
  query := `INSERT INTO registrations(user_id, event_id) VALUES (?, ?)`
  stmt, err := db.DB.Prepare(query)
  if err != nil {
    return err
  }

  defer stmt.Close()

  _, err = stmt.Exec(userId, e.ID)

  return err
}