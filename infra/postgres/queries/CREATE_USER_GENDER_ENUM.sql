DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_gender') THEN
        CREATE TYPE user_gender AS ENUM ('MASCULINO', 'FEMININO', 'OUTRO');
    END IF;
END $$;
