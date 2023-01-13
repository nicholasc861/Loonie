CREATE TABLE authentication(
    id uuid DEFAULT uuid_generate_v4(),
    date timestamp DEFAULT NOW() NOT NULL,
    access_token text NOT NULL,
    refres_token text NOT NULL

)