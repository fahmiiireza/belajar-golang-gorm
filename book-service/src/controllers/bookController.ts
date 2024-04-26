import { Request, Response } from 'express';
import Book from '../models/book'; // Import your Book model
import {
  ValidationError,
  UniqueConstraintError,
  ForeignKeyConstraintError,
} from 'sequelize';

export async function createBook(req: Request, res: Response) {
  try {
    const {
      isbn,
      title,
      language,
      total_copy,
      shelf_id,
      category_id,
      description,
    } = req.body;

    // Check if a book with the same ISBN already exists
    const existingBook = await Book.findOne({ where: { isbn } });
    if (existingBook) {
      return res
        .status(400)
        .json({ error: 'Book with the same ISBN already exists' });
    }

    // Create a new book
    const book = await Book.create({
      isbn,
      title,
      language,
      totalCopy: total_copy,
      shelfId: shelf_id,
      categoryId: category_id,
      description,
    });

    res.status(201).json(book);
  } catch (error: any) {
    if (
      error instanceof ValidationError ||
      error instanceof UniqueConstraintError ||
      error instanceof ForeignKeyConstraintError
    ) {
      return res.status(400).json({ error: error.message });
    } else {
      return res.status(500).send('Internal server error');
    }
  }
}

export async function getAllBooks(req: Request, res: Response) {
  try {
    const books = await Book.findAll();
    res.status(200).json(books);
  } catch (error) {
    console.error('Error getting all books:', error);
  }
}
