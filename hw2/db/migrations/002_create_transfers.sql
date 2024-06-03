CREATE TABLE transfers (
    id SERIAL PRIMARY KEY,
    from_entry_id INT NOT NULL,
    to_entry_id INT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (from_entry_id) REFERENCES entries(id),
    FOREIGN KEY (to_entry_id) REFERENCES entries(id)
);
