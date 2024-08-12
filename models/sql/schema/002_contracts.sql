-- +goose Up 
CREATE TABLE
    contracts (
        user_id UUID NOT NULL,
        created_at TIMESTAMP NOT NULL,
        contract_id varchar(128) NOT NULL UNIQUE,
        contract_name varchar(128) NOT NULL DEFAULT 'Untiteld',
        CONSTRAINT fk_contracts FOREIGN KEY (user_id) REFERENCES users (user_id)
    );

-- +goose Down
DROP TABLE IF EXISTS contracts;