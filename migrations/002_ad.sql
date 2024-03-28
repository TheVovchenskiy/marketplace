CREATE TABLE IF NOT EXISTS public.ad (
    id serial PRIMARY KEY CONSTRAINT id_is_positive CHECK (id > 0),
    "name" TEXT NOT NULL CONSTRAINT valid_name CHECK (
        length("name") > 0
        AND length("name") <= 150
    ),
    description TEXT NOT NULL CONSTRAINT valid_description CHECK (length(description) > 0),
    cents_price BIGINT NOT NULL CONSTRAINT valid_price CHECK (cents_price >= 0),
    picture_url TEXT NOT NULL CONSTRAINT valid_picture_url CHECK (length(picture_url) > 0)
);

---- create above / drop below ----
DROP TABLE IF EXISTS public.ad;
