CREATE TABLE IF NOT EXISTS public.user_profile (
    id serial PRIMARY KEY CONSTRAINT id_is_positive CHECK (id > 0),
    username TEXT UNIQUE NOT NULL CONSTRAINT valid_username CHECK (
        length(username) > 0
        AND length(username) <= 150
    ),
    password_hash bytea NOT NULL,
    salt bytea NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

---- create above / drop below ----
DROP TABLE IF EXISTS public.user_profile;
