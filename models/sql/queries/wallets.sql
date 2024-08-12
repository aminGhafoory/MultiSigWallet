-- name: CreateaWallet :exec
INSERT INTO wallets(
    user_id,
    created_at,
    wallet_id,
    wallet_name,
    wallet_keystore,
    wallet_address
)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
);

-- name: GetUserWallets :many
SELECT *
FROM
    users 
JOIN wallets ON 
    users.user_id = wallets.user_id 
WHERE 
    users.user_id=$1;


-- name: GetUserWalletByID :one
SELECT wallet_keystore
	FROM wallets
	INNER JOIN users ON
		wallets.user_id = users.user_id
	WHERE users.user_id=$1 AND wallet_id =$2;

-- name: DeleteWallet :exec
DELETE FROM wallets
WHERE user_id=$1 AND wallet_id=$2;