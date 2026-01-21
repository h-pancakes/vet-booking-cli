CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    phone TEXT NOT NULL,
    email TEXT NOT NULL,
    CONSTRAINT users_email_unique_ci UNIQUE (LOWER(email))
);

CREATE TABLE appointments (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    pet_name TEXT NOT NULL,
    pet_species TEXT NOT NULL,
    pet_age INTEGER NOT NULL,
    pet_weight REAL NOT NULL,
    vaccinated BOOLEAN NOT NULL,
    appointment_type TEXT NOT NULL,
    vet_name TEXT NOT NULL,
    appointment_time TIMESTAMPTZ NOT NULL,

    CONSTRAINT pet_age_positive CHECK (pet_age >= 0),
    CONSTRAINT pet_weight_positive CHECK (pet_weight > 0),
    CONSTRAINT appointment_in_future CHECK (appointment_time > now())
);
