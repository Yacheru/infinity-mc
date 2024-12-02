CREATE TABLE IF NOT EXISTS messages (
    player VARCHAR,
    player_uuid UUID,
    message TEXT,
    sent_at VARCHAR,
    PRIMARY KEY ( player, sent_at )
);
