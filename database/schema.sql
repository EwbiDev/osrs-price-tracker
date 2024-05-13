CREATE TABLE
    IF NOT EXISTS Items (
        id INTEGER PRIMARY KEY NOT NULL,
        name TEXT NOT NULL,
        icon TEXT NOT NULL,
        trade_limit INTEGER NOT NULL,
        members BOOLEAN NOT NULL,
        item_value INTEGER NOT NULL,
        low_alch INTEGER NOT NULL,
        high_alch INTEGER NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
    );

CREATE TABLE
    IF NOT EXISTS Official_Prices (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        item_id INTEGER NOT NULL,
        price INTEGER NOT NULL,
        last_price INTEGER NOT NULL,
        volume INTEGER NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        FOREIGN KEY (item_id) REFERENCES Items (id)
    );

CREATE TABLE
    IF NOT EXISTS Wiki_Prices (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        item_id INTEGER NOT NULL,
        avg_high_price INTEGER NOT NULL,
        high_price_volume INTEGER NOT NULL,
        avg_low_price INTEGER NOT NULL,
        low_price_volume INTEGER NOT NULL,
        timescale TEXT CHECK (timescale IN ('5m', '1h', '6h', '24h')) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        FOREIGN KEY (item_id) REFERENCES Items (id)
    );

CREATE TRIGGER IF NOT EXISTS update_item_timestamp BEFORE
UPDATE ON Items FOR EACH ROW BEGIN
UPDATE Items
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = OLD.id;

END;

CREATE TRIGGER IF NOT EXISTS update_official_price_timestamp BEFORE
UPDATE ON Official_Prices FOR EACH ROW BEGIN
UPDATE Official_Prices
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = OLD.id;

END;

CREATE TRIGGER IF NOT EXISTS update_wiki_price_timestamp BEFORE
UPDATE ON Wiki_Prices FOR EACH ROW BEGIN
UPDATE Wiki_Prices
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = OLD.id;

END;