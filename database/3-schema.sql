CREATE ROLE vault_db_user LOGIN SUPERUSER PASSWORD 'mypass';
CREATE ROLE readonly NOINHERIT;

GRANT SELECT ON ALL TABLES IN SCHEMA public TO "readonly";
