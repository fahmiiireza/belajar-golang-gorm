
const mockBookRequest = {
    isbn: '9780743273564',
    title: 'Mock Book',
    language: 'English',
    totalCopy: 1,
    shelfId: 1,
    categoryId: 1,
    description: 'Mock description',
  };

const mockBookData = {
    ...mockBookRequest,
    id: 1,
    created_at: new Date(),
    updated_at: new Date(),
    deleted_at: null,
  };

const userAuthData = {
    username: 'admin123',
    password: 'admin123',
  }
  export { mockBookRequest, mockBookData, userAuthData };