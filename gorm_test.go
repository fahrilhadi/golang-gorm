package golang_gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:@tcp(localhost:3306)/db_golang-gorm?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T)  {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T)  {
	err := db.Exec("insert into sample(id, name) values (?, ?)", "1", "Fahril").Error
	assert.Nil(t, err)

	err2 := db.Exec("insert into sample(id, name) values (?, ?)", "2", "Hadi").Error
	assert.Nil(t, err2)

	err3 := db.Exec("insert into sample(id, name) values (?, ?)", "3", "Abu").Error
	assert.Nil(t, err3)

	err4 := db.Exec("insert into sample(id, name) values (?, ?)", "4", "Hanif").Error
	assert.Nil(t, err4)
}

type Sample struct {
	Id string
	Name string
}

func TestRawSQL(t *testing.T)  {
	var sample Sample
	err := db.Raw("select id, name from sample where id = ?", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "Fahril", sample.Name)

	var samples []Sample
	err2 := db.Raw("select id, name from sample").Scan(&samples).Error
	assert.Nil(t, err2)
	assert.Equal(t, 4, len(samples))
}

func TestSQLRow(t *testing.T)  {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		var id string
		var name string

		err2 := rows.Scan(&id, &name)
		assert.Nil(t, err2)

		samples = append(samples, Sample{
			Id: id,
			Name: name,
		})
	}
	assert.Equal(t, 4, len(samples))
}

func TestScanRow(t *testing.T)  {
	rows, err := db.Raw("select id, name from sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next() {
		err2 := db.ScanRows(rows, &samples)
		assert.Nil(t, err2)
	}
	assert.Equal(t, 4, len(samples))
}