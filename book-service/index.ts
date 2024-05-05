// index.ts
import express, { Request, Response } from 'express';
import sequelize from './sequelize';
import bookRoutes from './src/routes/bookRoutes'; // Import book routes
import { verifyToken } from './src/middleware';
import { exec } from 'child_process';
const app = express();
const port = process.env.PORT || 3000;

app.use(express.json());

app.get('/test', (req: Request, res: Response) => {
  res.send('Test');
});
app.use(verifyToken);
app.use('/books', bookRoutes);



console.log('Attempting to authenticate with the database...');
sequelize
  .authenticate()
  .then(() => {
    console.log(
      'Connection to the database has been established successfully.'
    );
    exec('npx sequelize-cli db:migrate', (error, stdout, stderr) => {
      if (error) {
        console.error(`Error running migration: ${error}`);
        return;
      }
      console.log(`Migration output: ${stdout}`);

      exec(
        'npx sequelize-cli db:seed:all',
        (seedError, seedStdout, seedStderr) => {
          try {
            if (seedError) {
              throw seedError;
            }
            console.log(`Seeding output: ${seedStdout}`);
          } catch (error) {
            console.error(`Error running seeding: ${error}`);
          }
        }
      );

      app.listen(port, () => {
        console.log(`Server is running on port ${port}`);
      });
    });
  })
  .catch((error) => {
    console.error('Unable to connect to the database:', error);
    process.exit(1); // Exit the process if unable to connect to the database
  });

export { app };
