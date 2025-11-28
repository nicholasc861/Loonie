CREATE TABLE "user" (
    id uuid DEFAULT uuidv7() PRIMARY KEY,
    name text NOT NULL,
    email_address text UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE category (
    id uuid PRIMARY KEY DEFAULT uuidv7(),
    user_id uuid REFERENCES "user"(id),
    description text NOT NULL,
    parent_id uuid REFERENCES category(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE(user_id, description, parent_id)
);

CREATE TABLE payment_method (
    id uuid DEFAULT uuidv7() PRIMARY KEY,
    user_id uuid REFERENCES "user"(id),
    label varchar,
    type varchar NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE rewards_program (
    id uuid PRIMARY KEY DEFAULT uuidv7(),
    issuer varchar,
    issuer_type varchar, --- i.e. Airline, Bank Rewards Program, Hotel
    name varchar,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE credit_card (
    id uuid DEFAULT uuidv7() PRIMARY KEY,
    rewards_program_id uuid REFERENCES rewards_program(id),
    issuer varchar NOT NULL,
    name varchar NOT NULL,
    network varchar NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE (issuer, name)
);

CREATE TABLE user_credit_card (
    payment_method_id uuid REFERENCES payment_method(id),
    credit_card_id uuid REFERENCES credit_card(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE transaction (
    id uuid DEFAULT uuidv7() PRIMARY KEY,
    category_id uuid REFERENCES category(id),
    user_id uuid REFERENCES "user"(id) NOT NULL,
    payment_method_id uuid REFERENCES payment_method(id),
    credit_card_id uuid REFERENCES credit_card(id),
    merchant_name varchar,
    transaction_date DATE DEFAULT CURRENT_DATE,
    description varchar,
    amount NUMERIC(1000, 2),
    currency_code varchar NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE cashback_rate (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    category_id uuid REFERENCES category(id),
    credit_card_id uuid REFERENCES credit_card(id), --- Type of points / cashback earned
    multiplier NUMERIC,
    from_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
