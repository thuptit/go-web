// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: org.sql

package datastore

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createOrg = `-- name: CreateOrg :execrows
INSERT INTO org (org_id, org_extl_id, org_name, org_description, org_kind_id, create_app_id, create_user_id,
                 create_timestamp, update_app_id, update_user_id, update_timestamp)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
`

type CreateOrgParams struct {
	OrgID           uuid.UUID
	OrgExtlID       string
	OrgName         string
	OrgDescription  string
	OrgKindID       uuid.UUID
	CreateAppID     uuid.UUID
	CreateUserID    uuid.NullUUID
	CreateTimestamp time.Time
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
}

func (q *Queries) CreateOrg(ctx context.Context, arg CreateOrgParams) (int64, error) {
	result, err := q.db.Exec(ctx, createOrg,
		arg.OrgID,
		arg.OrgExtlID,
		arg.OrgName,
		arg.OrgDescription,
		arg.OrgKindID,
		arg.CreateAppID,
		arg.CreateUserID,
		arg.CreateTimestamp,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const createOrgKind = `-- name: CreateOrgKind :execrows
insert into org_kind (org_kind_id, org_kind_extl_id, org_kind_desc, create_app_id, create_user_id, create_timestamp,
                      update_app_id, update_user_id, update_timestamp)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9)
`

type CreateOrgKindParams struct {
	OrgKindID       uuid.UUID
	OrgKindExtlID   string
	OrgKindDesc     string
	CreateAppID     uuid.UUID
	CreateUserID    uuid.NullUUID
	CreateTimestamp time.Time
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
}

func (q *Queries) CreateOrgKind(ctx context.Context, arg CreateOrgKindParams) (int64, error) {
	result, err := q.db.Exec(ctx, createOrgKind,
		arg.OrgKindID,
		arg.OrgKindExtlID,
		arg.OrgKindDesc,
		arg.CreateAppID,
		arg.CreateUserID,
		arg.CreateTimestamp,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const deleteOrg = `-- name: DeleteOrg :execrows
DELETE
FROM org
WHERE org_id = $1
`

func (q *Queries) DeleteOrg(ctx context.Context, orgID uuid.UUID) (int64, error) {
	result, err := q.db.Exec(ctx, deleteOrg, orgID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const findOrgByExtlID = `-- name: FindOrgByExtlID :one
SELECT o.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       o.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc
FROM org o
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
WHERE org_extl_id = $1
`

type FindOrgByExtlIDRow struct {
	OrgID          uuid.UUID
	OrgExtlID      string
	OrgName        string
	OrgDescription string
	OrgKindID      uuid.UUID
	OrgKindExtlID  string
	OrgKindDesc    string
}

func (q *Queries) FindOrgByExtlID(ctx context.Context, orgExtlID string) (FindOrgByExtlIDRow, error) {
	row := q.db.QueryRow(ctx, findOrgByExtlID, orgExtlID)
	var i FindOrgByExtlIDRow
	err := row.Scan(
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.OrgKindID,
		&i.OrgKindExtlID,
		&i.OrgKindDesc,
	)
	return i, err
}

const findOrgByExtlIDWithAudit = `-- name: FindOrgByExtlIDWithAudit :one
SELECT o.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       ok.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc,
       o.create_app_id,
       a.org_id           create_app_org_id,
       a.app_extl_id      create_app_extl_id,
       a.app_name         create_app_name,
       a.app_description  create_app_description,
       o.create_user_id,
       cu.first_name      create_user_first_name,
       cu.last_name       create_user_last_name,
       o.create_timestamp,
       o.update_app_id,
       a2.org_id          update_app_org_id,
       a2.app_extl_id     update_app_extl_id,
       a2.app_name        update_app_name,
       a2.app_description update_app_description,
       o.update_user_id,
       uu.first_name      update_user_first_name,
       uu.last_name       update_user_last_name,
       o.update_timestamp
FROM org o
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
         INNER JOIN app a on a.app_id = o.create_app_id
         INNER JOIN app a2 on a2.app_id = o.update_app_id
         INNER JOIN users cu on cu.user_id = o.create_user_id
         INNER JOIN users uu on uu.user_id = o.update_user_id
WHERE o.org_extl_id = $1
`

type FindOrgByExtlIDWithAuditRow struct {
	OrgID                uuid.UUID
	OrgExtlID            string
	OrgName              string
	OrgDescription       string
	OrgKindID            uuid.UUID
	OrgKindExtlID        string
	OrgKindDesc          string
	CreateAppID          uuid.UUID
	CreateAppOrgID       uuid.UUID
	CreateAppExtlID      string
	CreateAppName        string
	CreateAppDescription string
	CreateUserID         uuid.NullUUID
	CreateUserFirstName  string
	CreateUserLastName   string
	CreateTimestamp      time.Time
	UpdateAppID          uuid.UUID
	UpdateAppOrgID       uuid.UUID
	UpdateAppExtlID      string
	UpdateAppName        string
	UpdateAppDescription string
	UpdateUserID         uuid.NullUUID
	UpdateUserFirstName  string
	UpdateUserLastName   string
	UpdateTimestamp      time.Time
}

func (q *Queries) FindOrgByExtlIDWithAudit(ctx context.Context, orgExtlID string) (FindOrgByExtlIDWithAuditRow, error) {
	row := q.db.QueryRow(ctx, findOrgByExtlIDWithAudit, orgExtlID)
	var i FindOrgByExtlIDWithAuditRow
	err := row.Scan(
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.OrgKindID,
		&i.OrgKindExtlID,
		&i.OrgKindDesc,
		&i.CreateAppID,
		&i.CreateAppOrgID,
		&i.CreateAppExtlID,
		&i.CreateAppName,
		&i.CreateAppDescription,
		&i.CreateUserID,
		&i.CreateUserFirstName,
		&i.CreateUserLastName,
		&i.CreateTimestamp,
		&i.UpdateAppID,
		&i.UpdateAppOrgID,
		&i.UpdateAppExtlID,
		&i.UpdateAppName,
		&i.UpdateAppDescription,
		&i.UpdateUserID,
		&i.UpdateUserFirstName,
		&i.UpdateUserLastName,
		&i.UpdateTimestamp,
	)
	return i, err
}

const findOrgByID = `-- name: FindOrgByID :one
SELECT o.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       o.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc
FROM org o
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
WHERE o.org_id = $1
`

type FindOrgByIDRow struct {
	OrgID          uuid.UUID
	OrgExtlID      string
	OrgName        string
	OrgDescription string
	OrgKindID      uuid.UUID
	OrgKindExtlID  string
	OrgKindDesc    string
}

func (q *Queries) FindOrgByID(ctx context.Context, orgID uuid.UUID) (FindOrgByIDRow, error) {
	row := q.db.QueryRow(ctx, findOrgByID, orgID)
	var i FindOrgByIDRow
	err := row.Scan(
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.OrgKindID,
		&i.OrgKindExtlID,
		&i.OrgKindDesc,
	)
	return i, err
}

const findOrgByIDWithAudit = `-- name: FindOrgByIDWithAudit :one
SELECT o.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       ok.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc,
       o.create_app_id,
       a.org_id           create_app_org_id,
       a.app_extl_id      create_app_extl_id,
       a.app_name         create_app_name,
       a.app_description  create_app_description,
       o.create_user_id,
       cu.first_name      create_user_first_name,
       cu.last_name       create_user_last_name,
       o.create_timestamp,
       o.update_app_id,
       a2.org_id          update_app_org_id,
       a2.app_extl_id     update_app_extl_id,
       a2.app_name        update_app_name,
       a2.app_description update_app_description,
       o.update_user_id,
       uu.first_name      update_user_first_name,
       uu.last_name       update_user_last_name,
       o.update_timestamp
FROM org o
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
         INNER JOIN app a on a.app_id = o.create_app_id
         INNER JOIN app a2 on a2.app_id = o.update_app_id
         INNER JOIN users cu on cu.user_id = o.create_user_id
         INNER JOIN users uu on uu.user_id = o.update_user_id
WHERE o.org_id = $1
`

type FindOrgByIDWithAuditRow struct {
	OrgID                uuid.UUID
	OrgExtlID            string
	OrgName              string
	OrgDescription       string
	OrgKindID            uuid.UUID
	OrgKindExtlID        string
	OrgKindDesc          string
	CreateAppID          uuid.UUID
	CreateAppOrgID       uuid.UUID
	CreateAppExtlID      string
	CreateAppName        string
	CreateAppDescription string
	CreateUserID         uuid.NullUUID
	CreateUserFirstName  string
	CreateUserLastName   string
	CreateTimestamp      time.Time
	UpdateAppID          uuid.UUID
	UpdateAppOrgID       uuid.UUID
	UpdateAppExtlID      string
	UpdateAppName        string
	UpdateAppDescription string
	UpdateUserID         uuid.NullUUID
	UpdateUserFirstName  string
	UpdateUserLastName   string
	UpdateTimestamp      time.Time
}

func (q *Queries) FindOrgByIDWithAudit(ctx context.Context, orgID uuid.UUID) (FindOrgByIDWithAuditRow, error) {
	row := q.db.QueryRow(ctx, findOrgByIDWithAudit, orgID)
	var i FindOrgByIDWithAuditRow
	err := row.Scan(
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.OrgKindID,
		&i.OrgKindExtlID,
		&i.OrgKindDesc,
		&i.CreateAppID,
		&i.CreateAppOrgID,
		&i.CreateAppExtlID,
		&i.CreateAppName,
		&i.CreateAppDescription,
		&i.CreateUserID,
		&i.CreateUserFirstName,
		&i.CreateUserLastName,
		&i.CreateTimestamp,
		&i.UpdateAppID,
		&i.UpdateAppOrgID,
		&i.UpdateAppExtlID,
		&i.UpdateAppName,
		&i.UpdateAppDescription,
		&i.UpdateUserID,
		&i.UpdateUserFirstName,
		&i.UpdateUserLastName,
		&i.UpdateTimestamp,
	)
	return i, err
}

const findOrgByName = `-- name: FindOrgByName :one
SELECT o.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       o.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc
FROM org o
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
WHERE o.org_name = $1
`

type FindOrgByNameRow struct {
	OrgID          uuid.UUID
	OrgExtlID      string
	OrgName        string
	OrgDescription string
	OrgKindID      uuid.UUID
	OrgKindExtlID  string
	OrgKindDesc    string
}

func (q *Queries) FindOrgByName(ctx context.Context, orgName string) (FindOrgByNameRow, error) {
	row := q.db.QueryRow(ctx, findOrgByName, orgName)
	var i FindOrgByNameRow
	err := row.Scan(
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.OrgKindID,
		&i.OrgKindExtlID,
		&i.OrgKindDesc,
	)
	return i, err
}

const findOrgByNameWithAudit = `-- name: FindOrgByNameWithAudit :one
SELECT o.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       ok.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc,
       o.create_app_id,
       a.org_id           create_app_org_id,
       a.app_extl_id      create_app_extl_id,
       a.app_name         create_app_name,
       a.app_description  create_app_description,
       o.create_user_id,
       cu.first_name      create_user_first_name,
       cu.last_name       create_user_last_name,
       o.create_timestamp,
       o.update_app_id,
       a2.org_id          update_app_org_id,
       a2.app_extl_id     update_app_extl_id,
       a2.app_name        update_app_name,
       a2.app_description update_app_description,
       o.update_user_id,
       uu.first_name      update_user_first_name,
       uu.last_name       update_user_last_name,
       o.update_timestamp
FROM org o
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
         INNER JOIN app a on a.app_id = o.create_app_id
         INNER JOIN app a2 on a2.app_id = o.update_app_id
         INNER JOIN users cu on cu.user_id = o.create_user_id
         INNER JOIN users uu on uu.user_id = o.update_user_id
WHERE o.org_name = $1
`

type FindOrgByNameWithAuditRow struct {
	OrgID                uuid.UUID
	OrgExtlID            string
	OrgName              string
	OrgDescription       string
	OrgKindID            uuid.UUID
	OrgKindExtlID        string
	OrgKindDesc          string
	CreateAppID          uuid.UUID
	CreateAppOrgID       uuid.UUID
	CreateAppExtlID      string
	CreateAppName        string
	CreateAppDescription string
	CreateUserID         uuid.NullUUID
	CreateUserFirstName  string
	CreateUserLastName   string
	CreateTimestamp      time.Time
	UpdateAppID          uuid.UUID
	UpdateAppOrgID       uuid.UUID
	UpdateAppExtlID      string
	UpdateAppName        string
	UpdateAppDescription string
	UpdateUserID         uuid.NullUUID
	UpdateUserFirstName  string
	UpdateUserLastName   string
	UpdateTimestamp      time.Time
}

func (q *Queries) FindOrgByNameWithAudit(ctx context.Context, orgName string) (FindOrgByNameWithAuditRow, error) {
	row := q.db.QueryRow(ctx, findOrgByNameWithAudit, orgName)
	var i FindOrgByNameWithAuditRow
	err := row.Scan(
		&i.OrgID,
		&i.OrgExtlID,
		&i.OrgName,
		&i.OrgDescription,
		&i.OrgKindID,
		&i.OrgKindExtlID,
		&i.OrgKindDesc,
		&i.CreateAppID,
		&i.CreateAppOrgID,
		&i.CreateAppExtlID,
		&i.CreateAppName,
		&i.CreateAppDescription,
		&i.CreateUserID,
		&i.CreateUserFirstName,
		&i.CreateUserLastName,
		&i.CreateTimestamp,
		&i.UpdateAppID,
		&i.UpdateAppOrgID,
		&i.UpdateAppExtlID,
		&i.UpdateAppName,
		&i.UpdateAppDescription,
		&i.UpdateUserID,
		&i.UpdateUserFirstName,
		&i.UpdateUserLastName,
		&i.UpdateTimestamp,
	)
	return i, err
}

const findOrgKindByExtlID = `-- name: FindOrgKindByExtlID :one
SELECT org_kind_id, org_kind_extl_id, org_kind_desc, create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp
FROM org_kind
WHERE org_kind_extl_id = $1
`

func (q *Queries) FindOrgKindByExtlID(ctx context.Context, orgKindExtlID string) (OrgKind, error) {
	row := q.db.QueryRow(ctx, findOrgKindByExtlID, orgKindExtlID)
	var i OrgKind
	err := row.Scan(
		&i.OrgKindID,
		&i.OrgKindExtlID,
		&i.OrgKindDesc,
		&i.CreateAppID,
		&i.CreateUserID,
		&i.CreateTimestamp,
		&i.UpdateAppID,
		&i.UpdateUserID,
		&i.UpdateTimestamp,
	)
	return i, err
}

const findOrgKinds = `-- name: FindOrgKinds :many

SELECT org_kind_id, org_kind_extl_id, org_kind_desc, create_app_id, create_user_id, create_timestamp, update_app_id, update_user_id, update_timestamp
FROM org_kind
`

// ---------------------------------------------------------------------------------------------------------------------
// Org Kind
// ---------------------------------------------------------------------------------------------------------------------
func (q *Queries) FindOrgKinds(ctx context.Context) ([]OrgKind, error) {
	rows, err := q.db.Query(ctx, findOrgKinds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []OrgKind
	for rows.Next() {
		var i OrgKind
		if err := rows.Scan(
			&i.OrgKindID,
			&i.OrgKindExtlID,
			&i.OrgKindDesc,
			&i.CreateAppID,
			&i.CreateUserID,
			&i.CreateTimestamp,
			&i.UpdateAppID,
			&i.UpdateUserID,
			&i.UpdateTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findOrgs = `-- name: FindOrgs :many
SELECT o.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       o.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc
FROM org o
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
ORDER BY org_name
`

type FindOrgsRow struct {
	OrgID          uuid.UUID
	OrgExtlID      string
	OrgName        string
	OrgDescription string
	OrgKindID      uuid.UUID
	OrgKindExtlID  string
	OrgKindDesc    string
}

func (q *Queries) FindOrgs(ctx context.Context) ([]FindOrgsRow, error) {
	rows, err := q.db.Query(ctx, findOrgs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindOrgsRow
	for rows.Next() {
		var i FindOrgsRow
		if err := rows.Scan(
			&i.OrgID,
			&i.OrgExtlID,
			&i.OrgName,
			&i.OrgDescription,
			&i.OrgKindID,
			&i.OrgKindExtlID,
			&i.OrgKindDesc,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findOrgsByKindExtlID = `-- name: FindOrgsByKindExtlID :many
SELECT o.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       ok.org_kind_extl_id,
       ok.org_kind_desc
FROM org o
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
WHERE ok.org_kind_extl_id = $1
`

type FindOrgsByKindExtlIDRow struct {
	OrgID          uuid.UUID
	OrgExtlID      string
	OrgName        string
	OrgDescription string
	OrgKindExtlID  string
	OrgKindDesc    string
}

func (q *Queries) FindOrgsByKindExtlID(ctx context.Context, orgKindExtlID string) ([]FindOrgsByKindExtlIDRow, error) {
	rows, err := q.db.Query(ctx, findOrgsByKindExtlID, orgKindExtlID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindOrgsByKindExtlIDRow
	for rows.Next() {
		var i FindOrgsByKindExtlIDRow
		if err := rows.Scan(
			&i.OrgID,
			&i.OrgExtlID,
			&i.OrgName,
			&i.OrgDescription,
			&i.OrgKindExtlID,
			&i.OrgKindDesc,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const findOrgsWithAudit = `-- name: FindOrgsWithAudit :many
SELECT o.org_id,
       o.org_extl_id,
       o.org_name,
       o.org_description,
       ok.org_kind_id,
       ok.org_kind_extl_id,
       ok.org_kind_desc,
       o.create_app_id,
       a.org_id           create_app_org_id,
       a.app_extl_id      create_app_extl_id,
       a.app_name         create_app_name,
       a.app_description  create_app_description,
       o.create_user_id,
       cu.first_name      create_user_first_name,
       cu.last_name       create_user_last_name,
       o.create_timestamp,
       o.update_app_id,
       a2.org_id          update_app_org_id,
       a2.app_extl_id     update_app_extl_id,
       a2.app_name        update_app_name,
       a2.app_description update_app_description,
       o.update_user_id,
       uu.first_name      update_user_first_name,
       uu.last_name       update_user_last_name,
       o.update_timestamp
FROM org o
         INNER JOIN org_kind ok on ok.org_kind_id = o.org_kind_id
         INNER JOIN app a on a.app_id = o.create_app_id
         INNER JOIN app a2 on a2.app_id = o.update_app_id
         INNER JOIN users cu on cu.user_id = o.create_user_id
         INNER JOIN users uu on uu.user_id = o.update_user_id
`

type FindOrgsWithAuditRow struct {
	OrgID                uuid.UUID
	OrgExtlID            string
	OrgName              string
	OrgDescription       string
	OrgKindID            uuid.UUID
	OrgKindExtlID        string
	OrgKindDesc          string
	CreateAppID          uuid.UUID
	CreateAppOrgID       uuid.UUID
	CreateAppExtlID      string
	CreateAppName        string
	CreateAppDescription string
	CreateUserID         uuid.NullUUID
	CreateUserFirstName  string
	CreateUserLastName   string
	CreateTimestamp      time.Time
	UpdateAppID          uuid.UUID
	UpdateAppOrgID       uuid.UUID
	UpdateAppExtlID      string
	UpdateAppName        string
	UpdateAppDescription string
	UpdateUserID         uuid.NullUUID
	UpdateUserFirstName  string
	UpdateUserLastName   string
	UpdateTimestamp      time.Time
}

func (q *Queries) FindOrgsWithAudit(ctx context.Context) ([]FindOrgsWithAuditRow, error) {
	rows, err := q.db.Query(ctx, findOrgsWithAudit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindOrgsWithAuditRow
	for rows.Next() {
		var i FindOrgsWithAuditRow
		if err := rows.Scan(
			&i.OrgID,
			&i.OrgExtlID,
			&i.OrgName,
			&i.OrgDescription,
			&i.OrgKindID,
			&i.OrgKindExtlID,
			&i.OrgKindDesc,
			&i.CreateAppID,
			&i.CreateAppOrgID,
			&i.CreateAppExtlID,
			&i.CreateAppName,
			&i.CreateAppDescription,
			&i.CreateUserID,
			&i.CreateUserFirstName,
			&i.CreateUserLastName,
			&i.CreateTimestamp,
			&i.UpdateAppID,
			&i.UpdateAppOrgID,
			&i.UpdateAppExtlID,
			&i.UpdateAppName,
			&i.UpdateAppDescription,
			&i.UpdateUserID,
			&i.UpdateUserFirstName,
			&i.UpdateUserLastName,
			&i.UpdateTimestamp,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrg = `-- name: UpdateOrg :execrows
UPDATE org
SET org_name         = $1,
    org_description  = $2,
    update_app_id    = $3,
    update_user_id   = $4,
    update_timestamp = $5
WHERE org_id = $6
`

type UpdateOrgParams struct {
	OrgName         string
	OrgDescription  string
	UpdateAppID     uuid.UUID
	UpdateUserID    uuid.NullUUID
	UpdateTimestamp time.Time
	OrgID           uuid.UUID
}

func (q *Queries) UpdateOrg(ctx context.Context, arg UpdateOrgParams) (int64, error) {
	result, err := q.db.Exec(ctx, updateOrg,
		arg.OrgName,
		arg.OrgDescription,
		arg.UpdateAppID,
		arg.UpdateUserID,
		arg.UpdateTimestamp,
		arg.OrgID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
