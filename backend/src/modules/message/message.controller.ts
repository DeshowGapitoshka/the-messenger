import { Request, Response } from 'express';
import { classMembersString, IMessage } from './message.type';

import { crypto } from '../../utils/functions/crypto.functions';
import { DATABASE } from '../../utils/constants/database.constants';

// ???
type MessageRequest = Request<any, any, IMessage, any>;

export class MessageController {
    static async createMessage(req: MessageRequest, res: Response) {
        try {
            const message = req.body;
            const cryptoText = crypto(message.createdAt, message.text);

            await DATABASE.query(`INSERT INTO messages ${classMembersString} VALUES ($1, $2, $3, $4, $5)`, [
                cryptoText,
                message.createdAt,
                message.createdBy,
                message.status,
                message.isChanged,
            ]);

            res.status(200).json('Message was created');
        } catch (err) {
            res.status(500).json('ERROR. Message was not created');
            console.log(err);
        }
    }
}
