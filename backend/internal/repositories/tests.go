package repositories

import (
	"context"
	"database/sql"

	"github.com/C-dexTeam/codex/internal/domains"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TestRepository struct {
	db *sqlx.DB
}

type dbModelTests struct {
	ID        sql.NullString `db:"id"`
	ChapterID sql.NullString `db:"chapter_id"`
}

type dbModelInputs struct {
	ID     sql.NullString `db:"id"`
	TestID sql.NullString `db:"test_id"`
	Value  sql.NullString `db:"value"`
}

type dbModelOutputs struct {
	ID     sql.NullString `db:"id"`
	TestID sql.NullString `db:"test_id"`
	Value  sql.NullString `db:"value"`
}

func (r *TestRepository) dbModelFromAppModelTest(appModel domains.Test) (dbModel dbModelTests) {
	if appModel.GetID() != uuid.Nil {
		dbModel.ID.String = appModel.GetID().String()
		dbModel.ID.Valid = true
	}
	if appModel.ChapterID() != uuid.Nil {
		dbModel.ChapterID.String = appModel.ChapterID().String()
		dbModel.ChapterID.Valid = true
	}

	return
}

func (r *TestRepository) dbModelFromAppFilter(filter domains.TestFilter) (dbModel dbModelTests) {
	if filter.ID != uuid.Nil {
		dbModel.ID.String = filter.ID.String()
		dbModel.ID.Valid = true
	}
	if filter.ChapterID != uuid.Nil {
		dbModel.ChapterID.String = filter.ChapterID.String()
		dbModel.ChapterID.Valid = true
	}

	return
}

func (r *TestRepository) dbModelFromAppFilterInput(filter domains.GeneralFilter) (dbModel dbModelInputs) {
	if filter.ID != uuid.Nil {
		dbModel.ID.String = filter.ID.String()
		dbModel.ID.Valid = true
	}
	if filter.TestID != uuid.Nil {
		dbModel.TestID.String = filter.TestID.String()
		dbModel.TestID.Valid = true
	}

	return
}

func (r *TestRepository) dbModelFromAppFilterOutput(filter domains.GeneralFilter) (dbModel dbModelOutputs) {
	if filter.ID != uuid.Nil {
		dbModel.ID.String = filter.ID.String()
		dbModel.ID.Valid = true
	}
	if filter.TestID != uuid.Nil {
		dbModel.TestID.String = filter.TestID.String()
		dbModel.TestID.Valid = true
	}

	return
}

func (r *TestRepository) dbModelFromAppModelInput(appModel domains.Input) (dbModel dbModelInputs) {
	if appModel.GetTestID() != uuid.Nil {
		dbModel.TestID.String = appModel.GetTestID().String()
		dbModel.TestID.Valid = true
	}
	if appModel.GetValue() != "" {
		dbModel.Value.String = appModel.GetValue()
		dbModel.Value.Valid = true
	}

	return
}

func (r *TestRepository) dbModelFromAppModelOutput(appModel domains.Output) (dbModel dbModelOutputs) {
	if appModel.GetTestID() != uuid.Nil {
		dbModel.TestID.String = appModel.GetTestID().String()
		dbModel.TestID.Valid = true
	}
	if appModel.GetValue() != "" {
		dbModel.Value.String = appModel.GetValue()
		dbModel.Value.Valid = true
	}

	return
}

func (r *TestRepository) dbModelToAppModelTest(dbModel dbModelTests) (test domains.Test) {
	test.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		uuid.MustParse(dbModel.ChapterID.String),
		nil,
		nil,
	)

	return
}

func (r *TestRepository) dbModelToAppModelInput(dbModel dbModelInputs) (input domains.Input) {
	input.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		uuid.MustParse(dbModel.TestID.String),
		dbModel.Value.String,
	)

	return
}

func (r *TestRepository) dbModelToAppModelOutput(dbModel dbModelOutputs) (output domains.Output) {
	output.Unmarshal(
		uuid.MustParse(dbModel.ID.String),
		uuid.MustParse(dbModel.TestID.String),
		dbModel.Value.String,
	)

	return
}

func NewTestRepository(db *sqlx.DB) domains.ITestRepository {
	return &TestRepository{db: db}
}

func (r *TestRepository) FilterTest(ctx context.Context, filter domains.TestFilter, limit, page int64) (tests []domains.Test, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilter(filter)
	dbResult := []dbModelTests{}

	query := `
	SELECT
		r.*
	FROM 
		t_tests r
	WHERE
		($1::uuid IS NULL OR r.id = $1::uuid) AND
		($2::uuid IS NULL OR r.chapter_id = $2::uuid)
	LIMIT $3 OFFSET $4
	`

	err = r.db.SelectContext(
		ctx,
		&dbResult,
		query,
		dbFilter.ID,
		dbFilter.ChapterID,
		limit,
		(page-1)*limit,
	)
	if err != nil {
		return
	}
	for _, dbModel := range dbResult {
		tests = append(tests, r.dbModelToAppModelTest(dbModel))
	}
	return
}

func (r *TestRepository) FilterInput(ctx context.Context, filter domains.GeneralFilter, limit, page int64) (inputs []domains.Input, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilterInput(filter)
	dbResult := []dbModelInputs{}

	query := `
	SELECT
		*
	FROM 
		t_inputs r
	WHERE
		($1::uuid IS NULL OR r.id = $1::uuid) AND
		($2::uuid IS NULL OR r.test_id = $2::uuid)
	LIMIT $3 OFFSET $4
	`

	err = r.db.SelectContext(
		ctx,
		&dbResult,
		query,
		dbFilter.ID,
		dbFilter.TestID,
		limit,
		(page-1)*limit,
	)
	if err != nil {
		return
	}
	for _, dbModel := range dbResult {
		inputs = append(inputs, r.dbModelToAppModelInput(dbModel))
	}
	return
}

func (r *TestRepository) FilterOutput(ctx context.Context, filter domains.GeneralFilter, limit, page int64) (outputs []domains.Output, dataCount int64, err error) {
	dbFilter := r.dbModelFromAppFilterOutput(filter)
	dbResult := []dbModelOutputs{}

	query := `
	SELECT
		*
	FROM 
		t_outputs r
	WHERE
		($1::uuid IS NULL OR r.id = $1::uuid) AND
		($2::uuid IS NULL OR r.test_id = $2::uuid)
	LIMIT $3 OFFSET $4
	`

	err = r.db.SelectContext(
		ctx,
		&dbResult,
		query,
		dbFilter.ID,
		dbFilter.TestID,
		limit,
		(page-1)*limit,
	)
	if err != nil {
		return
	}
	for _, dbModel := range dbResult {
		outputs = append(outputs, r.dbModelToAppModelOutput(dbModel))
	}
	return
}

func (r *TestRepository) AddTest(ctx context.Context, input *domains.Test) (uuid.UUID, error) {
	dbModel := r.dbModelFromAppModelTest(*input)
	query := `
		INSERT INTO
			t_tests (chapter_id)
		VALUES
			($1)
		RETURNING id
	`

	var id uuid.UUID
	err := r.db.QueryRowxContext(ctx, query, dbModel.ChapterID).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *TestRepository) AddInput(ctx context.Context, input *domains.Input) (uuid.UUID, error) {
	dbModel := r.dbModelFromAppModelInput(*input)
	query := `
		INSERT INTO
			t_inputs (test_id, value)
		VALUES
			($1, $2)
		RETURNING id
	`

	var id uuid.UUID
	err := r.db.QueryRowxContext(ctx, query, dbModel.TestID, dbModel.Value).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *TestRepository) AddOutput(ctx context.Context, output *domains.Output) (uuid.UUID, error) {
	dbModel := r.dbModelFromAppModelOutput(*output)
	query := `
		INSERT INTO
			t_outputs (test_id, value)
		VALUES
			($1, $2)
		RETURNING id
	`

	var id uuid.UUID
	err := r.db.QueryRowxContext(ctx, query, dbModel.TestID, dbModel.Value).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
