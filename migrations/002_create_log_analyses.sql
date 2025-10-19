DROP TABLE IF EXISTS log_analysis CASCADE;

CREATE TABLE log_analysis (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    filename TEXT,
    total_requests INT DEFAULT 0,
    unique_ips INT DEFAULT 0,
    error_count INT DEFAULT 0,
    average_response FLOAT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- Auto-update updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_log_analysis_updated_at 
    BEFORE UPDATE ON log_analysis 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();