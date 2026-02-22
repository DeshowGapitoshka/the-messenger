import Route from 'express';
import { MessageController } from './message.controller';

export const messageRoutes = Route();

messageRoutes.post('/message', MessageController.createMessage);
