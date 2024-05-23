CREATE TABLE games (
      id serial,
      player_id_1 text,
      player_id_2 text,
      player_cast_1 text,
      player_cast_2 text,
      created_at timestamp NOT NULL,
      PRIMARY KEY(id)
);
