DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_status') THEN
        CREATE TYPE user_status AS ENUM ('ATIVO', 'INATIVO', 'EM ANÁLISE');
    END IF;
END $$;
