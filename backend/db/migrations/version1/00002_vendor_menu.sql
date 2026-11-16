-- +goose Up
CREATE TABLE food_courts(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY ,
    name TEXT NOT NULL ,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE vendors (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY ,
    name TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    logourl TEXT,
    food_court_id BIGINT NOT NULL references food_courts (id),
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE menu_categories (
     id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY ,
     name TEXT NOT NULL ,
     vendor_id BIGINT NOT NULL REFERENCES vendors (id) ON DELETE CASCADE,
     created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    unique (vendor_id,name)
);

CREATE TABLE products(
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY ,
    name TEXT NOT NULL ,
    description TEXT NOT NULL,
    product_url TEXT,
    price int NOT NULL,
    menu_categories_id BIGINT NOT NUll REFERENCES menu_categories (id) ON delete cascade,
    vendor_id BIGINT NOT NULL REFERENCES vendors (id) ON DELETE CASCADE,
    is_available BOOLEAN NOT NULL DEFAULT true,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    unique (vendor_id,name)
);

CREATE TABLE vendor_staff (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    vendor_id BIGINT NOT NULL REFERENCES vendors (id) ON DELETE CASCADE,
    username TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE vendor_staff;
DROP TABLE products;
DROP TABLE menu_categories;
DROP TABLE vendors;
DROP TABLE food_courts;