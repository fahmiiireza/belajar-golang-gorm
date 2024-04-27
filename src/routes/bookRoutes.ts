import {Router } from 'express';
import {createBook,getAllBooks, updateBook, getBookById, deleteBook} from '../controllers/bookController';
const router = Router();

router.post('/',createBook);
router.get('/',getAllBooks);
router.patch('/:id',updateBook);
router.get('/:id',getBookById);
router.delete('/:id',deleteBook);

export default router;
