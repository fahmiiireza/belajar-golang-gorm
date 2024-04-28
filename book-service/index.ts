// index.ts
// TODO: make failed tests
//  make tests for all routes
// LEARN ABOUT TESTING
// LEARN ABOUT SPA VS MPA
// UPDATE ROOT MAKEFILE TO ALSO HANDLE THE BOOK-SERVICE
// LEARN ABOUT MICROSERVICES
// UPDATE NOTION TO FILL WITH ALL OF THIS TASK
import express, { Request, Response } from 'express';
import sequelize from './sequelize';
import bookRoutes from './src/routes/bookRoutes'; // Import book routes
import { verifyToken } from './src/middlewares/authMiddleware';
const app = express();
const port = process.env.PORT || 3000;

app.use(express.json());
app.use(verifyToken)
app.use('/books', bookRoutes);

app.get('/', (req: Request, res: Response) => {
  res.send('Test');
});

sequelize
  .authenticate()
  .then(() => {
    console.log(
      'Connection to the database has been established successfully.'
    );
    // If you have associations or any other setup code, you can place it here
    
    // Start the server after database connection is established
    app.listen(port, () => {
      console.log(`Server is running on port ${port}`);
    });
  })
  .catch((error) => {
    console.error('Unable to connect to the database:', error);
    process.exit(1); // Exit the process if unable to connect to the database
  });

export { app };
