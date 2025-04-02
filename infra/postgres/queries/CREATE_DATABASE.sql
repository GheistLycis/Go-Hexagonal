DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = '${NAME}') THEN  
        CREATE DATABASE ${NAME} WITH OWNER ${USER};
    END IF;
END $$;
