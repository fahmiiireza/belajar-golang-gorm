import { Request, Response } from 'express';
import Book from '../models/book'; 
import Author from '../models/author';
import Shelf from '../models/shelf';
import Category from '../models/category';
import {
  ValidationError,
  UniqueConstraintError,
  ForeignKeyConstraintError,
} from 'sequelize';

 async function createBook(req: Request, res: Response) {
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
    console.log('start');
    // Check if a book with the same ISBN already exists
    const existingBook = await Book.findOne({ where: { isbn } });
    if (existingBook) {
      console.log('Book with the same ISBN already exists');
      console.log(existingBook);
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
    console.log(error);
    if (
      error instanceof ValidationError ||
      error instanceof UniqueConstraintError ||
      error instanceof ForeignKeyConstraintError
    ) {
      return res.status(400).json({ error: error.message });
    } else {
        console.log(error);
      return res.status(500).send('Internal server error');
    }
  }
}


 async function updateBook(req: Request, res: Response) {
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
    //   if(!req.params.id){
    //     return res.status(400).json({ error: 'Book ID is required' });
    //   }

      const bookToUpdate = await Book.findByPk(req.params.id);
      if (!bookToUpdate) {
        return res.status(404).json({ error: 'Book not found' });
      }
  
      await bookToUpdate.update({
        isbn,
        title,
        language,
        totalCopy: total_copy,
        shelfId: shelf_id,
        categoryId: category_id,
        description,
      });
  
      res.status(200).json(bookToUpdate);
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

 async function getAllBooks(req: Request,res: Response) {
  try {
  const books = await Book.findAll();
    res.status(200).json(books);
  } catch (error) {
    console.error('Error getting all books:', error);
  }
}

async function getBookById(req: Request, res: Response) {
    try {
        // const book = await Book.findByPk(req.params.id,);
        const book = await Book.findOne({ where: { id: req.params.id }, 
          include: [
            { model: Category },
            { model: Author },
            { model: Shelf }
          ]});
        if (!book) {
        return res.status(404).json({ error: 'Book not found' });
        }
        res.status(200).json(book);
    } catch (error) {
        console.error('Error getting book by ID:', error);
    }
    
}


async function deleteBook(req: Request, res: Response) {
    try {
        const bookToDelete = await Book.findByPk(req.params.id);
        if (!bookToDelete) {
        return res.status(404).json({ error: 'Book not found' });
        }
        await bookToDelete.destroy();
        res.status(204).send();
    } catch (error) {
        console.error('Error deleting book:', error);
    }

}
export { createBook, getAllBooks, updateBook, getBookById, deleteBook };
