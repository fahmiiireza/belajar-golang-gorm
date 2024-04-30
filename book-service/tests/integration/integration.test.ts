import supertest from 'supertest';
import testServer from './testServer';
import sequelize from '../../sequelize';
import {
  userAuthData,
  mockBookData,
  mockBookRequest,
  mockBookRequestForDelete,
  invalidBookData,
  mockBookRequestForUpdate,
} from '../testData';
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

  test('DELETE /books/:id deletes a book', async () => {
    const newBook = await Book.create(mockBookRequestForDelete);
    const response = await supertest(testServer)
      .delete(`/books/${newBook.id}`)
      .set('Authorization', `Bearer ${authToken}`);

    expect(response.status).toBe(204);
    expect(response.body).toEqual({});

    await Book.destroy({ where: { id: newBook.id }, force: true });
  });

  // test('GET /librarian/:id/books returns a list of books created by the librarian', async () => {
  //   const response = await supertest(testServer)
  //     .get('/librarian/1/books')
  //     .set('Authorization', `Bearer ${authToken}`);

  //   expect(response.status).toBe(200);
  //   expect(response.body).toBeInstanceOf(Array);
  // });
});

describe('CRUD Validation', () => {
  let authToken: string;
  beforeAll(async () => {
    const loginResponse = await supertest('http://user-service:8080')
      .post('/login')
      .send(userAuthData);
    authToken = loginResponse.body.token;
  });

  test('Get one book with invalid id', async () => {
    const response = await supertest(testServer)
      .get('/books/100')
      .set('Authorization', `Bearer ${authToken}`);

    expect(response.status).toBe(404);
    expect(response.body.error).toBe('Book not found');
  });
  test('Create book handles validation where theres same ISBN', async () => {
    const newBook = await Book.create(mockBookRequest);
    const response = await supertest(testServer)
      .post('/books')
      .set('Authorization', `Bearer ${authToken}`)
      .send(mockBookRequest);

    expect(response.status).toBe(400);
    expect(response.body.error).toBe('Book with the same ISBN already exists');

    await Book.destroy({ where: { isbn: newBook.isbn }, force: true });
  });
  test('Create book handles validation where theres missing field', async () => {
    const response = await supertest(testServer)
      .post('/books')
      .set('Authorization', `Bearer ${authToken}`)
      .send(invalidBookData);

    expect(response.status).toBe(400);
    expect(response.body.error).toBeDefined();
  });
  test('Create book handle validation of foreign key constraint', async () => {
    const response = await supertest(testServer)
      .post('/books')
      .set('Authorization', `Bearer ${authToken}`)
      .send({ ...mockBookRequest, shelfId: 1000 });

    expect(response.status).toBe(400);
    expect(response.body.error).toBeDefined();
  });

  test('Update book handles validation where it tries to update ISBN that already exists ', async () => {
    const createdBook = await Book.create(mockBookRequest);
    const bookToUpdate = await Book.create(mockBookRequestForUpdate);
    const response = await supertest(testServer)
      .patch(`/books/${bookToUpdate.id}`)
      .set('Authorization', `Bearer ${authToken}`)
      .send({ isbn: createdBook.isbn });

    await Book.destroy({ where: { id: createdBook.id }, force: true });
    await Book.destroy({ where: { id: bookToUpdate.id }, force: true });
    expect(response.status).toBe(400);
    expect(response.body.error).toBe('Cannot update ISBN, already exists');
  });
  test('Update book handles trying to update a non existent book', async () => {
    const response = await supertest(testServer)
      .patch('/books/1000')
      .set('Authorization', `Bearer ${authToken}`)
      .send(mockBookRequestForUpdate);

    expect(response.status).toBe(404);
    expect(response.body.error).toBe('Book not found');
  });

  test('Delete book handles trying to delete a non existent book', async () => {
    const response = await supertest(testServer)
      .delete('/books/1000')
      .set('Authorization', `Bearer ${authToken}`);

    expect(response.status).toBe(404);
    expect(response.body.error).toBe('Book not found');
  });

  afterAll(async () => {
    await sequelize.close();
    testServer.close();
  });
});
