CREATE TABLE IF NOT EXISTS pokemons(
  id bigserial PRIMARY KEY,
  name VARCHAR(60) NOT NULL,
  type VARCHAR(100) []
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
  updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW();
);
