CREATE EXTENSION "uuid-ossp" SCHEMA public;
CREATE TYPE equity_type AS ENUM('Stock', 'Option', 'Crypto');
CREATE TYPE side AS ENUM('Buy', 'Sell');
CREATE TYPE option_type AS ENUM('Call', 'Put');

CREATE TABLE account (
    id uuid DEFAULT uuid_generate_v4(),
    provider_account_id text NOT NULL,
    nickname text NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE balance (
    id uuid DEFAULT uuid_generate_v4(),
    account_id uuid NOT NULL,
    date timestamp NOT NULL,
    amount float,
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES account(id)
);


CREATE TABLE stock_info (
    id uuid DEFAULT uuid_generate_v4(),
    qt_id text NOT NULL,
    ticker text NOT NULL,
    description text NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE option_info (
    id uuid DEFAULT uuid_generate_v4(),
    qt_id text NOT NULL,
    ticker text NOT NULL,
    description text NOT NULL,
    expiration timestamp NOT NULL,
    strike float NOT NULL,
    type option_type NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE crypto_info (
    id uuid DEFAULT uuid_generate_v4(),
    qt_id text NOT NULL,
    ticker text NOT NULL,
    description text NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE trade_group (
    id uuid DEFAULT uuid_generate_v4(),
    account_id uuid NOT NULL,
    option_id uuid,
    stock_id uuid,
    crypto_id uuid,
    type equity_type NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (account_id) REFERENCES account(id),
    FOREIGN KEY (option_id) REFERENCES option_info(id),
    FOREIGN KEY (stock_id) REFERENCES stock_info(id),
    FOREIGN KEY (crypto_id) REFERENCES crypto_info(id)
);

CREATE TABLE trade (
    id uuid DEFAULT uuid_generate_v4(),
    trade_group_id uuid NOT NULL,
    date timestamp NOT NULL,
    price float NOT NULL,
    quantity float NOT NULL,
    commission float NOT NULL,
    gross_amount float NOT NULL,
    net_amount float NOT NULL,
    side side NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (trade_group_id) REFERENCES trade_group(id)
);






