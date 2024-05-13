-- name: SelectItem :one
SELECT
    *
FROM
    Items
WHERE
    id = ?
LIMIT
    1;

-- name: ListItems :many
SELECT
    *
FROM
    Items
ORDER BY
    name;

-- name: CountItems :many
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