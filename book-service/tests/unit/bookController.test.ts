import { Request, Response } from 'express';
import {
  createBook,
  updateBook,
  deleteBook,
  getAllBooks,
  getBookById,
} from '../../src/controllers/bookController';
import { mockBookData, mockBookRequest, mockBooks } from '../testData';

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

    await createBook(req, res);

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
    const req = {
      body: mockBookRequest,
    } as Request;

    const res = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn(),
    } as unknown as Response;

    (Book.findOne as jest.Mock).mockResolvedValue(mockBookData);

    await createBook(req, res);

    expect(res.status).toHaveBeenCalledWith(400);
    expect(res.json).toHaveBeenCalledWith({
      error: 'Book with the same ISBN already exists',
    });
  });
});

describe('getAllBooks', () => {
  test('should return a list of books', async () => {
    const req = {} as unknown as Request;

    const res = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn(),
    } as unknown as Response;

    (Book.findAll as jest.Mock).mockResolvedValue(mockBooks);

    await getAllBooks(req, res);

    expect(res.status).toHaveBeenCalledWith(200);
    expect(res.json).toHaveBeenCalledWith(mockBooks);
  });
});

describe('getBookById', () => {
  test('should return a single book by its ID', async () => {
    const req = { params: { id: mockBookData.id } } as unknown as Request;

    const res = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn(),
    } as unknown as Response;

    (Book.findByPk as jest.Mock).mockResolvedValue(mockBookData);

    await getBookById(req, res);

    expect(res.status).toHaveBeenCalledWith(200);
    expect(res.json).toHaveBeenCalledWith(mockBookData);
  });

  test('should return 404 if book with the given ID is not found', async () => {
    const req = { params: { id: 100 } } as unknown as Request;

    const res = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn(),
    } as unknown as Response;

    (Book.findByPk as jest.Mock).mockResolvedValue(null);

    await getBookById(req, res);

    expect(res.status).toHaveBeenCalledWith(404);
    expect(res.json).toHaveBeenCalledWith({ error: 'Book not found' });
  });
});

describe('updateBook', () => {
  test('should update an existing book', async () => {
    const req = {
      params: { id: mockBookData.id },
      body: {
        title: 'New Updated Title',
      },
    } as unknown as Request;
    const res = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn(),
    } as unknown as Response;

    (Book.findByPk as jest.Mock).mockResolvedValue(mockBookData);
    (Book.update as jest.Mock).mockResolvedValue({
      ...mockBookData,
      title: 'New Updated Title',
    });

    await updateBook(req, res);

    expect(res.status).toHaveBeenCalledWith(200);
    expect(res.json).toHaveBeenCalledWith(
      expect.objectContaining({
        id: mockBookData.id,
        title: 'New Updated Title',
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
    const req = {
      params: { id: 1 },
      body: mockBookRequest,
    } as unknown as Request;
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

describe('deleteBook', () => {
  test('should delete an existing book', async () => {
    const req = { params: { id: mockBookData.id } } as unknown as Request;
    const res = {
      status: jest.fn().mockReturnThis(),
      send: jest.fn(),
    } as unknown as Response;

    (Book.findByPk as jest.Mock).mockResolvedValue(mockBookData);

    (Book.destroy as jest.Mock).mockResolvedValue(1);

    await deleteBook(req, res);

    expect(res.status).toHaveBeenCalledWith(204);
    expect(res.send).toHaveBeenCalled();
  });

  test('should return 404 if book with the given ID is not found', async () => {
    const req = { params: { id: 1 } } as unknown as Request;
    const res = {
      status: jest.fn().mockReturnThis(),
      json: jest.fn(),
    } as unknown as Response;

    (Book.findByPk as jest.Mock).mockResolvedValue(null);

    await deleteBook(req, res);

    expect(res.status).toHaveBeenCalledWith(404);
    expect(res.json).toHaveBeenCalledWith({ error: 'Book not found' });
  });
});
