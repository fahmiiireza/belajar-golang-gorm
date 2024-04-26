import { Request, Response, Router } from 'express';
import Book from '../models/book'; // Import the Book model



export async function createBook(req:Request,res:Response) {
    try {
        const {
          isbn,
          title,
          language,
          total_copy,
          shelf_id,
          category_id,
          description,
          created_at,
          updated_at,
          deleted_at
        } = req.body;
    
        // Create a new book
        const book = await Book.create({
          isbn,
          title,
          language,
          totalCopy: total_copy,
          shelfId: shelf_id,
          categoryId: category_id,
          description,
          createdAt: created_at,
          updatedAt: updated_at,
          deletedAt: deleted_at
        });
    
        res.status(201).json(book); // Respond with the created book
      } catch (error) {
        console.error('Error creating book:', error);
        res.status(500).send('Internal server error');
      }
}

export async function getAllBooks(req:Request,res:Response){
    try {
        const books = await Book.findAll();
        res.status(200).json(books);
      } catch (error) {
        console.error('Error getting all books:', error);
      }
}