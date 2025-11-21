-- +goose Up
CREATE TABLE companies (
    ID UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(15) UNIQUE NOT NULL,
    description VARCHAR(3000),
    employee_count INT NOT NULL,
    registered BOOLEAN NOT NULL,
    company_type VARCHAR(20) NOT NULL,
    CONSTRAINT company_type_check CHECK (
        company_type IN (
            'Corporation', 'NonProfit', 'Cooperative', 'SoleProprietorship'
        )
    )
);

-- +goose Down
DROP TABLE companies;
