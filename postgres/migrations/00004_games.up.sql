CREATE TABLE games (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    game_id UUID NOT NULL,
    winner UUID NOT NULL,
    player1_id UUID NOT NULL,
    player2_id UUID NOT NULL,
    player1_cast VARCHAR(255) NOT NULL,
    player2_cast VARCHAR(255) NOT NULL,
    datetime TIMESTAMP NOT NULL,
    FOREIGN KEY (player1_id) REFERENCES players(id),
    FOREIGN KEY (player2_id) REFERENCES players(id),
    FOREIGN KEY (winner) REFERENCES players(id)
);