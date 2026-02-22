import { IDatabaseConfig } from '../types/IDatabaseConfig.types';
import dotenv from 'dotenv';

dotenv.config();

export const DATABASE_CONFIG: IDatabaseConfig = {
    USER: process.env.DATABASE_USER ? process.env.DATABASE_USER : '',
    PORT: process.env.DATABASE_PORT ? Number(process.env.DATABASE_PORT) : 0,
    DATABASE: process.env.DATABASE_DB ? process.env.DATABASE_DB : '',
    HOST: process.env.DATABASE_HOST ? process.env.DATABASE_HOST : '',
    PASSWORD: process.env.DATABASE_PASSWORD ? process.env.DATABASE_PASSWORD : '',
};
