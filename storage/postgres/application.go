package postgres

import (
	"errors"
	"fmt"
	"go_boilerplate/go_boilerplate_modules/application_service"
	"go_boilerplate/storage/repo"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
)

type applicationRepo struct {
	db *sqlx.DB
}

// NewApplicationRepo ...
func NewApplicationRepo(db *sqlx.DB) repo.ApplicationStorageI {
	return &applicationRepo{db: db}
}

func (r *applicationRepo) Create(entity application_service.CreateApplicationModel) (res application_service.ApplicationCreatedModel, err error) {
	insertQuery := `INSERT INTO application (
		id,
		body
	) VALUES (
		$1,
		$2
	)`

	_, err = r.db.Exec(insertQuery,
		entity.ID,
		entity.Body,
	)

	if err != nil {
		return res, err
	}

	res.ID = entity.ID

	return res, nil
}

func (r *applicationRepo) GetList(queryParam application_service.ApplicationQueryParamModel) (res application_service.ApplicationListModel, err error) {
	res.Applications = []application_service.ApplicationModel{}
	params := make(map[string]interface{})
	query := `SELECT
		id,
		body,
		created_at,
		updated_at
	FROM
		application`
	filter := " WHERE 1=1"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND (body ILIKE '%' || :search || '%')"
	}

	if len(queryParam.Order) > 0 {
		valid := regexp.MustCompile("^[A-Za-z0-9_]+$")
		if valid.MatchString(queryParam.Order) {
			order = fmt.Sprintf(" ORDER BY %s", queryParam.Order)
		} else {
			return res, errors.New("wrong order query param")
		}
	}

	switch strings.ToUpper(queryParam.Arrangement) {
	case "DESC":
		arrangement = " DESC"
	case "ASC":
		arrangement = " ASC"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	cQ := "SELECT count(1) FROM application" + filter
	row, err := r.db.NamedQuery(cQ, params)
	if err != nil {
		return res, err
	}
	defer row.Close()
	if row.Next() {
		err = row.Scan(
			&res.Count,
		)
		if err != nil {
			return res, err
		}
	}

	q := query + filter + order + arrangement + offset + limit
	rows, err := r.db.NamedQuery(q, params)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var application application_service.ApplicationModel
		err = rows.Scan(
			&application.ID,
			&application.Body,
			&application.CreatedAt,
			&application.UpdatedAt,
		)
		if err != nil {
			return res, err
		}
		res.Applications = append(res.Applications, application)
	}

	return res, nil
}

func (r *applicationRepo) GetByID(id string) (res application_service.ApplicationModel, err error) {
	query := `SELECT
		id,
		body,
		created_at,
		updated_at
	FROM
		application
	WHERE
		id = $1`

	row, err := r.db.Query(query, id)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(
			&res.ID,
			&res.Body,
			&res.CreatedAt,
			&res.UpdatedAt,
		)
		if err != nil {
			return res, err
		}
	} else {
		return res, errors.New("not found")
	}

	return res, nil
}

func (r *applicationRepo) Update(entity application_service.UpdateApplicationModel) (rowsAffected int64, err error) {
	query := `UPDATE application SET
		body = :body,
		updated_at = now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":   entity.ID,
		"body": entity.Body,
	}

	result, err := r.db.NamedExec(query, params)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (r *applicationRepo) Delete(id string) (rowsAffected int64, err error) {
	query := `DELETE FROM application WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
