import supertest from 'supertest'; 
import  testServer  from './testServer';
import  sequelize  from '../../sequelize'; 
describe('Integration Tests', () => {
  beforeAll(async () => {
    await sequelize.authenticate();
  });
  afterEach(async () => {
    testServer.close();
  });
  test('GET /api/books returns a list of books', async () => {
    const response = await supertest(testServer).get('/books')
    .set('Accept', 'application/json');
    expect(response.status).toBe(200);
  });
  test('Creating a new book adds it to the database', async () => {
    const response = await supertest(testServer)
      .post('/books')
      .set('Accept', 'application/json')
      .send({
        title: 'New Book',
        author: 'John Doe',
        isbn: '1234567890',
        language: 'English',
        total_copy: 10,
        shelf_id: 1,
        category_id: 1,
      });
    expect(response.status).toBe(201);
  });

  afterAll(async () => {
    await sequelize.close();
   testServer.close(); 
  });
});

// docker exec -it 0400b9e49e99 /bin/bash 






  // // Add more test cases as needed
  