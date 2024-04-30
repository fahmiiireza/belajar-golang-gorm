// testServer.ts

import express from 'express';
import { verifyToken } from '../../src/middleware';

import bookRoutes from '../../src/routes/bookRoutes'; 
const testApp = express();
const testPort = 4000; 

testApp.use(express.json());
testApp.use(verifyToken);
testApp.use('/books', bookRoutes);
testApp.get('/', (req, res) => {
  res.send('Test server is running');
});

const testServer = testApp.listen(testPort, () => {
  console.log(`Test server is running on port ${testPort}`);
});

export default testServer;
