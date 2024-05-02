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
import { verifyToken } from './src/middleware';
import { exec } from 'child_process';
const app = express();
const port = process.env.PORT || 3000;

app.use(express.json());
app.use(verifyToken);
app.use('/books', bookRoutes);

app.get('/', (req: Request, res: Response) => {
  res.send('Test');
});

console.log('Attempting to authenticate with the database...');
sequelize
  .authenticate()
  .then(() => {
    console.log(
      'Connection to the database has been established successfully.'
    );
    // If you have associations or any other setup code, you can place it here

    // Run Sequelize migration
    exec('npx sequelize-cli db:migrate', (error, stdout, stderr) => {
      if (error) {
        console.error(`Error running migration: ${error}`);
        return;
      }
      console.log(`Migration output: ${stdout}`);

      // Once migration is complete, run Sequelize seeding
      exec(
        'npx sequelize-cli db:seed:all',
        (seedError, seedStdout, seedStderr) => {
          if (seedError) {
            console.error(`Error running seeding: ${seedError}`);
            return;
          }
          console.log(`Seeding output: ${seedStdout}`);

          // Start the server after migration and seeding are complete
          app.listen(port, () => {
            console.log(`Server is running on port ${port}`);
          });
        }
      );
    });
  })
  .catch((error) => {
    console.error('Unable to connect to the database:', error);
    process.exit(1); // Exit the process if unable to connect to the database
  });

export { app };
