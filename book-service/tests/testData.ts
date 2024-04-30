const mockBookRequest = {
  isbn: '9780743273564',
  title: 'Mock Book',
  language: 'English',
  totalCopy: 1,
  shelfId: 1,
  categoryId: 1,
  description: 'Mock description',
};

const mockBookRequestForDelete = {
  isbn: '9780743564',
  title: 'Mock Book for delete',
  language: 'English',
  totalCopy: 1,
  shelfId: 1,
  categoryId: 1,
  description: 'Mock description',
};
const mockBookRequestForUpdate = {
  isbn: '97807435623232',
  title: 'Mock Book for Update',
  language: 'Indonesia',
  totalCopy: 1,
  shelfId: 1,
  categoryId: 1,
  description: 'Mock Update',
}
const invalidBookData = {
  // Missing required fields: isbn, title
  language: 'English',
  totalCopy: 10,
  shelfId: 'shelf_1',
  categoryId: 'category_1',
  description: 'Invalid book description',
};

const mockBookData = {
  ...mockBookRequest,
  id: 1,
  created_at: new Date(),
  updated_at: new Date(),
  deleted_at: null,
};
const mockBooks = [mockBookData];

const userAuthData = {
  username: 'librarian123',
  password: 'librarian123',
};
export {
  mockBookRequest,
  mockBookData,
  userAuthData,
  mockBookRequestForDelete,
  invalidBookData,
  mockBooks,
  mockBookRequestForUpdate,
};
