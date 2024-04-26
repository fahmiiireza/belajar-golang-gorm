// index.ts

import express, { Request, Response } from 'express';
import sequelize from './sequelize';
import bookRoutes from './src/routes/bookRoutes'; // Import book routes
import { CustomRequest, verifyToken } from './src/middlewares/authMiddleware';
const app = express();
const port = process.env.PORT || 3000;


app.use(express.json());
app.use(verifyToken)
app.use('/books', bookRoutes);

app.get('/', (req: Request, res: Response) => {
  res.send('Test');
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});

sequelize
  .authenticate()
  .then(() => {
    console.log(
      'Connection to the database has been established successfully.'
    );
    // If you have associations or any other setup code, you can place it here
  })
  .catch((error) => {
    console.error('Unable to connect to the database:', error);
    process.exit(1); // Exit the process if unable to connect to the database
  });
