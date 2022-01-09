CREATE TABLE post
(
    ident      BIGSERIAL PRIMARY KEY,
    caption    VARCHAR(4000) NOT NULL,
    images     VARCHAR(120)[],
    created_at TIMESTAMP DEFAULT now()
);
