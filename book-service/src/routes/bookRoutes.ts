import { Router, Request, Response, NextFunction } from 'express';
import { createBook, getAllBooks, updateBook, getBookById, deleteBook } from '../controllers/bookController';
import { checkRole, CustomRequest } from '../middleware';

const router = Router();

// Apply the checkRole middleware before routes that require it
router.use(checkRole);

// Define routes
router.get('/', getAllBooks);
router.get('/:id', getBookById);
router.post('/', createBook);
router.patch('/:id', updateBook);
router.delete('/:id', deleteBook);

export default router;
