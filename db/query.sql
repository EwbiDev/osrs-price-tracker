-- name: SelectItem :one
SELECT
    *
FROM
    Items
WHERE
    id = ?
LIMIT
    1;

-- name: SelectItems :many
SELECT
    *
FROM
    Items
WHERE
    (
        COALESCE(sqlc.narg (id), 0) = 0
        OR id = sqlc.narg (id)
    )
    AND (
        COALESCE(sqlc.narg (name), 0) = 0
        OR name LIKE CONCAT ("%", sqlc.narg (name), "%")
    )
    AND (
        COALESCE(sqlc.narg (icon), 0) = 0
        OR icon = sqlc.narg (icon)
    )
    AND (
        COALESCE(sqlc.narg (trade_limit), 0) = 0
        OR trade_limit = sqlc.narg (trade_limit)
    )
    AND (
        COALESCE(sqlc.narg (members), -1) = -1
        OR members = sqlc.narg (members)
    )
    AND (
        COALESCE(sqlc.narg (item_value), 0) = 0
        OR item_value = sqlc.narg (item_value)
    )
    AND (
        COALESCE(sqlc.narg (low_alch), 0) = 0
        OR low_alch = sqlc.narg (low_alch)
    )
    AND (
        COALESCE(sqlc.narg (high_alch), 0) = 0
        OR high_alch = sqlc.narg (high_alch)
    )
    AND (
        COALESCE(sqlc.narg (created_at), 0) = 0
        OR created_at = sqlc.narg (created_at)
    )
    AND (
        COALESCE(sqlc.narg (updated_at), 0) = 0
        OR updated_at = sqlc.narg (updated_at)
    );

-- name: ListItems :many
SELECT
    *
FROM
    Items
ORDER BY
    name;

-- name: CountItems :one
SELECT
    COUNT(*)
FROM
    Items;

-- name: InsertItem :one
INSERT INTO
    Items (
        id,
        name,
        icon,
        trade_limit,
        members,
        item_value,
        low_alch,
        high_alch
    )
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?) RETURNING *;

-- name: UpdateItem :one
UPDATE Items
SET
    name = ?,
    icon = ?,
    trade_limit = ?,
    members = ?,
    item_value = ?,
    low_alch = ?,
    high_alch = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = ? RETURNING *;

-- name: InsertOfficialPrice :one
INSERT INTO
    Official_Prices (
        item_id,
        price,
        last_price,
        volume,
        jagex_timestamp
    )
VALUES
    (?, ?, ?, ?, ?) RETURNING *;