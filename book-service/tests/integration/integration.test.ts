import supertest from 'supertest';
import testServer from './testServer';
import sequelize from '../../sequelize';
import { userAuthData, mockBookData, mockBookRequest } from '../testData';
import Book from '../../src/models/book';

describe('Book CRUD', () => {
  let authToken: string;
  beforeAll(async () => {
    await sequelize.authenticate();

    const loginResponse = await supertest('http://user-service:8080')
      .post('/login')
      .send(userAuthData);
    authToken = loginResponse.body.token;
  });

  test('GET /books returns a list of books', async () => {
    const response = await supertest(testServer)
      .get('/books')
      .set('Authorization', `Bearer ${authToken}`);

    expect(response.status).toBe(200);
    expect(response.body).toBeInstanceOf(Array);
  });

  test('POST /books create a new book', async () => {
    const response = await supertest(testServer)
      .post('/books')
      .set('Authorization', `Bearer ${authToken}`)
      .send(mockBookRequest);

    expect(response.status).toBe(201);
    expect(response.body).toEqual(
      expect.objectContaining({
        description: mockBookData.description,
        isbn: mockBookData.isbn,
        title: mockBookData.title,
      })
    );
    await Book.destroy({ where: { isbn: mockBookData.isbn }, force: true });
  });
  test('PATCH /books/:id updates a book', async () => {});

  afterAll(async () => {
    await sequelize.close();
    testServer.close();
  });
});

// docker exec -it 0400b9e49e99 /bin/bash

// // Add more test cases as needed
