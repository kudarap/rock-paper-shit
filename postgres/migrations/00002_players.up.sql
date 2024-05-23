CREATE TABLE players (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ranking INT NOT NULL,
    wins INT NOT NULL,
    losses INT NOT NULL,
    draws INT NOT NULL,
    play_count INT NOT NULL
);