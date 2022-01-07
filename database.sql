CREATE DATABASE billing;

CREATE TABLE medicines
(
    id         SERIAL       NOT NULL,
    name       VARCHAR(100) NOT NULL,
    price      NUMERIC      NOT NULL,
    location   VARCHAR(100) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT medicines_id_pk PRIMARY KEY (id)
);

CREATE TABLE promotions
(
    id          SERIAL       NOT NULL,
    description VARCHAR(300) NOT NULL,
    percentage  INTEGER      NOT NULL,
    start_date  TIMESTAMP    NOT NULL,
    finish_date TIMESTAMP    NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP,
    CONSTRAINT promotions_id_pk PRIMARY KEY (id)
);

CREATE TABLE invoices
(
    id            SERIAL    NOT NULL,
    total_price   NUMERIC   NOT NULL,
    promotion_id  INTEGER   NOT NULL,
    medicines_ids JSON      NOT NULL,
    created_at    TIMESTAMP NOT NULL DEFAULT now(),
    updated_at    TIMESTAMP,
    CONSTRAINT invoices_id_pk PRIMARY KEY (id),
    CONSTRAINT invoices_promotion_id_fk FOREIGN KEY (promotion_id) REFERENCES promotions (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
);

CREATE INDEX invoices_created_at_ix ON invoices (created_at);

