import { Router, Request, Response, NextFunction } from 'express';
import { createBook, getAllBooks, updateBook, getBookById, deleteBook } from '../controllers/bookController';
import { checkLibrarianRole, CustomRequest } from '../middleware';

const router = Router();


// Define routes
router.get('/', getAllBooks);
router.get('/:id', getBookById);
router.post('/',checkLibrarianRole, createBook);
router.patch('/:id',checkLibrarianRole, updateBook);
router.delete('/:id',checkLibrarianRole, deleteBook);

export default router;
