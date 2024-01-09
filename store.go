package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Abstraction
type Storage interface {
	GetDataFilter(*RecordRequest) ([]*Record, error)
	GetDataAll(*RecordRequest) ([]*Record, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgres() (*PostgresStore, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")
	sslmode := os.Getenv("SSLMODE")
	connStr := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s", host, user, dbname, password, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresStore{db: db}, nil
}

// Migration table and creation of dummy data
func (p *PostgresStore) Init() error {
	err := p.CreateTableRecord()
	if err != nil {
		return err
	}
	err = p.generateDummy()
	return err
}

// method : create table record if exist
func (p *PostgresStore) CreateTableRecord() error {
	query := `CREATE TABLE IF NOT EXISTS records (
		id int,
		name varchar(50),
		marks int,
		createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := p.db.Exec(query)
	return err
}

// method : generate dummy data
func (p *PostgresStore) generateDummy() error {
	query := `
	delete from records;
	INSERT INTO records (id,name, marks, createdAt) values (1,'evan',90,'2024-01-01');
	INSERT INTO records (id,name, marks, createdAt) values (1,'evan',10,'2024-01-01');
	INSERT INTO records (id,name, marks, createdAt) values (1,'evan',80,'2024-01-01');
	INSERT INTO records (id,name, marks, createdAt) values (1,'evan',20,'2024-01-01');
	INSERT INTO records (id,name, marks, createdAt) values (2,'roy',50,'2024-01-04');
	INSERT INTO records (id,name, marks, createdAt) values (2,'roy',10,'2024-01-04');
	INSERT INTO records (id,name, marks, createdAt) values (2,'roy',50,'2024-01-04');
	INSERT INTO records (id,name, marks, createdAt) values (2,'roy',30,'2024-01-04');
	INSERT INTO records (id,name, marks, createdAt) values (3,'gusman',30,'2024-01-04');
	
	`
	_, err := p.db.Exec(query)
	return err
}

// Get filtered Data
func (p *PostgresStore) GetDataFilter(req *RecordRequest) ([]*Record, error) {
	records := []*Record{}
	query := `select sum(marks) totalMarks,id,createdat 
	from records 
	where createdat between $1 and $2
	group by id,createdat
	order by id,createdat`
	rows, err := p.db.Query(query, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		record := &Record{}
		err := rows.Scan(&record.Marks, &record.Id, &record.CreatedAt)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

// Get All Data inside the Table
func (p *PostgresStore) GetDataAll(req *RecordRequest) ([]*Record, error) {
	records := []*Record{}
	query := `select * from records`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		record := &Record{}
		err := rows.Scan(&record.Id, &record.Name, &record.Marks, &record.CreatedAt)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}
