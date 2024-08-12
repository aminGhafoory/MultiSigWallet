-- name: CreateaContract :one
INSERT INTO
    contracts (
        user_id,
        created_at,
        contract_id,
        contract_name)
VALUES
    ($1,
    $2,
    $3,
    $4) 
RETURNING *;

-- name: GetUserContracts :many
SELECT
    *
FROM
    users
    JOIN contracts ON users.user_id = contracts.user_id
WHERE
    users.user_id = $1;

-- name: EditContractName :one
UPDATE contracts
SET
    contract_name = $1
WHERE
    contract_id = $2 RETURNING *;

-- name: DeleteContract :exec
DELETE FROM contracts
WHERE
    contract_id = $2
    AND user_id = $1;