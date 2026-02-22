import dotenv from 'dotenv';
import { IServerConfig } from '../types/IServerConfig.types';

dotenv.config();

// Need to type the ProcessEnv
export const SERVER_CONFIG: IServerConfig = {
    PORT: process.env.SERVER_PORT ? process.env.SERVER_PORT : '',
} as const;
