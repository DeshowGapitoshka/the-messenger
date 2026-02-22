import { Pool } from 'pg';
import { DATABASE_CONFIG } from './databaseConfig.constants';

export const DATABASE = new Pool({
    user: DATABASE_CONFIG.USER,
    host: DATABASE_CONFIG.HOST,
    port: DATABASE_CONFIG.PORT,
    password: DATABASE_CONFIG.PASSWORD,
    database: DATABASE_CONFIG.DATABASE,
});
