import { Request, Response } from 'express';
import {
  createBook,
  updateBook,
} from '../../src/controllers/bookController';
import  {mockBookData, mockBookRequest}  from '../testData';

import Book from '../../src/models/book';
jest.mock('../../src/models/book');



describe('createBook', () => {
  test('should create a new book based on valid request body given', async () => {

    const req = {
      body: mockBookRequest,
    } as Request;

    const res = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn(),
    } as unknown as Response;

    (Book.findOne as jest.Mock).mockResolvedValue(null);

    (Book.create as jest.Mock).mockResolvedValue(mockBookData);

    // Call the controller function
    await createBook(req, res);

    // Assert that the response status and json methods were called with the expected values
    expect(res.status).toHaveBeenCalledWith(201);
    expect(res.json).toHaveBeenCalledWith(
      expect.objectContaining({
        id: mockBookData.id,
        isbn: mockBookData.isbn,
        title: mockBookData.title,
        language: mockBookData.language,
        totalCopy: mockBookData.totalCopy,
        shelfId: mockBookData.shelfId,
        categoryId: mockBookData.categoryId,
        description: mockBookData.description,
        created_at: expect.any(Date),
        updated_at: expect.any(Date),
        deleted_at: mockBookData.deleted_at,
      })
    );
  });
  test('should return 400 if book with the same ISBN already exists', async () => {
    // Mock request object with the necessary properties for creating a book
    const req = {
      body: mockBookRequest,
    } as Request;

    // Mock response object with status and json methods
    const res = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn(),
    } as unknown as Response;

    // Mock the behavior of Book.findOne to simulate an existing book with the same ISBN
    (Book.findOne as jest.Mock).mockResolvedValue(mockBookData);

    // Call the controller function
    await createBook(req, res);

    // Assert that the response status and json methods were called with the expected values
    expect(res.status).toHaveBeenCalledWith(400);
    expect(res.json).toHaveBeenCalledWith({
      error: 'Book with the same ISBN already exists',
    });

  })
});

describe('updateBook', () => {
  test('should update an existing book', async () => {
    const req = { params: { id: mockBookData.id }, body: {
      "title": "New Updated Title",
    } } as unknown as Request;
    const res = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn(),
    } as unknown as Response;

    (Book.findByPk as jest.Mock).mockResolvedValue(mockBookData);
    (Book.update as jest.Mock).mockResolvedValue({...mockBookData, title: "New Updated Title"});

    await updateBook(req, res);

    expect(res.status).toHaveBeenCalledWith(200);
    expect(res.json).toHaveBeenCalledWith(
      expect.objectContaining({
        id: mockBookData.id,
        title: "New Updated Title",
        language: mockBookData.language,
        totalCopy: mockBookData.totalCopy,
        shelfId: mockBookData.shelfId,
        categoryId: mockBookData.categoryId,
        description: mockBookData.description,
        created_at: expect.any(Date),
        updated_at: expect.any(Date),
        deleted_at: mockBookData.deleted_at,
      })
    );
  });

  test('should return 404 if book with the given ID is not found', async () => {
    const req = { params: { id: 1 }, body: mockBookRequest } as unknown as Request;
    const res = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn(),
    } as unknown as Response;

    (Book.findByPk as jest.Mock).mockResolvedValue(null);

    await updateBook(req, res);

    expect(res.status).toHaveBeenCalledWith(404);
    expect(res.json).toHaveBeenCalledWith({ error: 'Book not found' });
  });
});
// describe('getAllBooks', () => {
//     test('should get all books', async () => {
//       const res = {
//         status: jest.fn().mockReturnThis(),
//         json: jest.fn(),
//       } as unknown as Response;

//       await getAllBooks(res);

//       expect(res.status).toHaveBeenCalledWith(200);
//       expect(res.json).toHaveBeenCalledWith(expect.any(Array));
//     });

//     // Add more test cases for edge cases, error handling, etc.
//   });

// describe('getBookById', () => {
//     test('should get a book by ID', async () => {
//       const req = { params: { id: 1 } } as Request;
//       const res = {
//         status: jest.fn().mockReturnThis(),
//         json: jest.fn(),
//       } as unknown as Response;

//       await getBookById(req, res);

//       expect(res.status).toHaveBeenCalledWith(200);
//       expect(res.json).toHaveBeenCalledWith(expect.any(Object));
//     });

//     // Add more test cases for edge cases, error handling, etc.
//   });

// describe('deleteBook', () => {
//     test('should delete a book by ID', async () => {
//       const req = { params: { id: /* provide valid book ID */ } } as Request;
//       const res = {
//         status: jest.fn().mockReturnThis(),
//         send: jest.fn(),
//       } as unknown as Response;

//       await deleteBook(req, res);

//       expect(res.status).toHaveBeenCalledWith(204);
//       expect(res.send).toHaveBeenCalled();
//     });

//     // Add more test cases for edge cases, error handling, etc.
//   });