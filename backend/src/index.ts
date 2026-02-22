import express, { json } from 'express';

import { SERVER_CONFIG } from './utils/constants/serverConfig.constants';
import { messageRoutes } from './modules/message/message.routes';

const server = express();

server.use(json());
server.use('/', messageRoutes);

server.listen(SERVER_CONFIG.PORT, () => {
    console.log('The server have been started corretly');
    console.log(`The port: ${SERVER_CONFIG.PORT}`);
});
