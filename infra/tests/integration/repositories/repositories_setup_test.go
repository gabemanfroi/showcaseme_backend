package repositories_test

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"showcaseme/infra/IoC"
	"showcaseme/infra/core"
	"showcaseme/infra/db"
	"showcaseme/internal/utils"
	"testing"
)

const (
	filename        = "../../../../.env.test"
	initialDataPath = "../../db/seeds/initial_data.sql"
)

type RepositoriesTestSuite struct {
	suite.Suite
	db *sql.DB
	m  *migrate.Migrate
}

func (s *RepositoriesTestSuite) SetupSuite() {
	loadInitialConfig()
	s.initDb()
	IoC.InitContainer()
	executeSeeds()
}

func executeSeeds() {
	_, filename, _, _ := runtime.Caller(0)

	seedsPath := filepath.Join(filepath.Dir(filename), initialDataPath)

	c, err := ioutil.ReadFile(seedsPath)
	utils.Check(err, "failed to read seeds file")
	sql := string(c)
	db.GetSqlInstance().Exec(sql)
}

func (s *RepositoriesTestSuite) initDb() {
	s.getConnection()
	s.runMigrations()
}

func (s *RepositoriesTestSuite) runMigrations() {
	var err error
	s.m, err = migrate.NewWithDatabaseInstance(getMigrationPath(), core.AppConfig.DbName, s.getDbDriver())
	require.NoError(s.T(), err)
	require.NoError(s.T(), s.m.Up())
}

func (s *RepositoriesTestSuite) getConnection() {
	var err error
	dbUrl := getDbUrl()
	s.db, err = sql.Open("postgres", dbUrl)
	require.NoError(s.T(), err)
}

func (s *RepositoriesTestSuite) getDbDriver() database.Driver {
	driver, err := postgres.WithInstance(s.db, &postgres.Config{DatabaseName: core.AppConfig.DbName})
	require.NoError(s.T(), err)
	return driver
}

func getMigrationPath() string {
	_, filename, _, _ := runtime.Caller(0)
	return "file://" + path.Join(filepath.Dir(filename), "../../../db/migrations")
}

func getDbUrl() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		core.AppConfig.DbUser,
		core.AppConfig.DbPassword,
		core.AppConfig.DbHost,
		core.AppConfig.DbPort,
		core.AppConfig.DbName,
	)
}

func loadInitialConfig() {

	utils.Check(godotenv.Load(filename), "error reading .env.test file...")
	core.LoadConfig()
}

func (s *RepositoriesTestSuite) TearDownSuite() {
	require.NoError(s.T(), s.m.Down())
}

func TestRepositories(t *testing.T) {
	suite.Run(t, new(RepositoriesTestSuite))
}
