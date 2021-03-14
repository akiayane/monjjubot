package main

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"monjjubot/database"
	"time"
)

type SnippetModel struct {
	DB *pgxpool.Pool
}

type BookPack struct {
	book_id   int
	course_id int
	subject string
	book_name string
	book_link string
}

func (m *SnippetModel) getByCourse(course_id int) ([]*database.BookPack, error) {
	stmt := "SELECT * FROM books WHERE course_id = $1"

	rows, err := m.DB.Query(context.Background(), stmt, course_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*database.BookPack

	for rows.Next() {

		s := &database.BookPack{}

		err = rows.Scan(&s.BookId, &s.CourseId, &s.Subject, &s.BookName, &s.BookLink)
		if err != nil {
			return nil, err
		}

		books = append(books, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (m *SnippetModel) getBySubject(course_id int, subject string) ([]*database.BookPack, error) {
	stmt := "SELECT * FROM books WHERE course_id = $1 and subject = $2"

	rows, err := m.DB.Query(context.Background(), stmt, course_id, subject)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []*database.BookPack

	for rows.Next() {

		s := &database.BookPack{}

		err = rows.Scan(&s.BookId, &s.CourseId, &s.Subject, &s.BookName, &s.BookLink)
		if err != nil {
			return nil, err
		}

		books = append(books, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (m *SnippetModel) registerUser(chat_id string, email string, vkey string) (bool, error) {
	stmt := "INSERT INTO users (chat_id, email, verifyed, vkey, create_time) VALUES ($1, $2, $3, $4, $5)"

	rows, err := m.DB.Query(context.Background(), stmt, chat_id, email, false, vkey, time.Now())
	if err != nil {
		return false, err
	}

	defer rows.Close()

	return true, nil
}

func (m *SnippetModel) userExist(chat_id string, email string) (bool) {
	stmt := "SELECT * FROM users WHERE chat_id = $1 and email = $2"
	rows, err := m.DB.Query(context.Background(), stmt, chat_id, email)
	if err != nil {
		return false
	}
	defer rows.Close()
	if rows.Next(){
		return true
	}else{
		return false
	}
}

func (m *SnippetModel) checkVerification(chat_id string) bool {
	stmt := "SELECT verified FROM users WHERE chat_id = $1"
	rows, err := m.DB.Query(context.Background(), stmt, chat_id)
	if err != nil {
		return false
	}
	status := rows.Next()
	defer rows.Close()
	return status
}

func (m *SnippetModel) confirmRegister(vkey string) (bool) {
	stmt := "UPDATE users SET verified = true WHERE vkey = $1"

	rows, err := m.DB.Query(context.Background(), stmt, vkey)
	if err != nil {
		return false
	}

	defer rows.Close()

	return true
}

// This will insert a new snippet into the database_server.
/*func (m *SnippetModel) Insert(title, content, expires string) (int, error) {

	stmt := "INSERT INTO snippets (title, content, created, expires) VALUES ($1, $2, $3, $4) RETURNING id"

	var id uint64
	days, err := strconv.Atoi(expires)
	if err != nil {
		return 0, err
	}
	result := m.DB.QueryRow(context.Background(), stmt, title, content, time.Now(), time.Now().AddDate(0, 0, days))
	err = result.Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil

}

// This will return a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets
			WHERE expires > now() AND id = $1`

	row := m.DB.QueryRow(context.Background(), stmt, id)

	s := &models.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil

}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := "SELECT id, title, content, created, expires FROM snippets WHERE expires > $1 ORDER BY created DESC LIMIT 10"

	rows, err := m.DB.Query(context.Background(), stmt, time.Now())
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {

		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}*/
