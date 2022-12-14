package versions

import (
	"database/sql"
	"time"
)

type VersionRepository interface {
	FindAll() ([]Version, error)
}

type VersionRepo struct {
	DB *sql.DB
}

func NewVersionRepo(db *sql.DB) *VersionRepo {
	return &VersionRepo{DB: db}
}

func (v *VersionRepo) FindAll() ([]Version, error) {

	result := []Version{}

	v1 := Version{
		Id:        1,
		VerNumber: OCPI_V_2_1_1,
		URL:       "https://localhost:9090/ocpi/emsp/versions/2.1.1",
		PartyId:   "VOL",
		CreateAt:  time.Now().Local(),
		UpdateAt:  time.Now().Local(),
	}

	result = append(result, v1)

	v2 := Version{
		Id:        2,
		VerNumber: OCPI_V_2_2_1,
		URL:       "https://localhost:9090/ocpi/emsp/versions/2.2.1",
		PartyId:   "VOL",
		CreateAt:  time.Now().Local(),
		UpdateAt:  time.Now().Local(),
	}

	result = append(result, v2)
	return result, nil
}
