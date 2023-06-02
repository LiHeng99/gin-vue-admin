package db_tools

import (
	"database/sql"
	"fmt"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/lib/pq"

	"github.com/flipped-aurora/gin-vue-admin/server/model/db_tools"
)

var dbMap = make(map[string]*sql.DB)

var dbMapMutex sync.RWMutex

func SetDB(key string, db *sql.DB) {
	dbMapMutex.Lock()
	defer dbMapMutex.Unlock()
	dbMap[key] = db
}

func GetDB(key string) (*sql.DB, bool) {
	dbMapMutex.RLock()
	defer dbMapMutex.RUnlock()

	db, ok := dbMap[key]
	return db, ok
}

type Database interface {
	Connect() error
	GetDB() *sql.DB
}

type MySQL struct {
	dsn string
	db  *sql.DB
}

func (m *MySQL) Connect() error {
	var err error
	m.db, err = sql.Open("mysql", m.dsn)
	return err
}

func (m *MySQL) GetDB() *sql.DB {
	return m.db
}

type PostgreSQL struct {
	dsn string
	db  *sql.DB
}

func (p *PostgreSQL) Connect() error {
	var err error
	p.db, err = sql.Open("postgres", p.dsn)
	return err
}

func (p *PostgreSQL) GetDB() *sql.DB {
	return p.db
}

func DatabaseFactory(dbType, dsn string) Database {
	switch dbType {
	case "mysql":
		return &MySQL{dsn: dsn}
	case "postgres":
		return &PostgreSQL{dsn: dsn}
	default:
		panic("Invalid database type")
	}
}

func OpenDb(dbInfo db_tools.DbInfo) *sql.DB {
	//dbMap = make(map[string]*sql.DB)
	// Assume the database driver is MySQL
	// Please replace the DSN string with your own settings
	db, err := sql.Open("mysql", dbInfo.DbUrl)
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
		return nil
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
		return nil
	}

	SetDB(strconv.FormatUint(uint64(dbInfo.ID), 10)+dbInfo.DbName, db)
	return db
}

type DatabaseManager struct {
	dbs map[string]Database
	mu  sync.Mutex
}

func NewDatabaseManager() *DatabaseManager {
	return &DatabaseManager{
		dbs: make(map[string]Database),
	}
}

func (m *DatabaseManager) Add(dbType, id, dsn string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	var db Database
	switch dbType {
	case "mysql":
		db = &MySQL{dsn: dsn}
	case "postgres":
		db = &PostgreSQL{dsn: dsn}
	default:
		return fmt.Errorf("invalid database type")
	}

	err := db.Connect()
	if err != nil {
		return err
	}

	m.dbs[id] = db
	return nil
}

func (m *DatabaseManager) Get(id string) (Database, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	db, ok := m.dbs[id]
	return db, ok
}
