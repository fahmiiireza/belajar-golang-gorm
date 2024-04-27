
const mockBookRequest = {
    id: 1,
    isbn: '9780743273564',
    title: 'Mock Book',
    language: 'English',
    totalCopy: 1,
    shelfId: 1,
    categoryId: 1,
    description: 'Mock description',
    createdAt: new Date(),
    updatedAt: new Date(),
    deletedAt: null,
  };

const mockBookData = {
    ...mockBookRequest,
    id: 1,
    createdAt: new Date(),
    updatedAt: new Date(),
    deletedAt: null,
  };

export { mockBookRequest, mockBookData };  
