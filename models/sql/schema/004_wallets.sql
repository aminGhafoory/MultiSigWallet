-- +goose Up 
CREATE TABLE
    wallets (
        user_id UUID NOT NULL,
        created_at TIMESTAMP NOT NULL,
        wallet_id UUID UNIQUE NOT NULL,
        wallet_name varchar(128) NOT NULL DEFAULT 'Untiteld',
        wallet_keystore TEXT NOT NULL,
        wallet_address TEXT NOT NULL UNIQUE,
        CONSTRAINT fk_wallets FOREIGN KEY (user_id) REFERENCES users (user_id)
    );

-- +goose Down
DROP TABLE IF EXISTS wallets;