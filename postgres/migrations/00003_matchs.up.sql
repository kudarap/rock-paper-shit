CREATE TABLE matches (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    player_id UUID NOT NULL,
    status VARCHAR(255) NOT NULL,
    datetime TIMESTAMP NOT NULL,
    FOREIGN KEY (player_id) REFERENCES players(id)
);