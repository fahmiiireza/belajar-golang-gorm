'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up(queryInterface, Sequelize) {
    await queryInterface.bulkInsert('authors', [
      {
        biography: 'Biography 1',
        nationality: 'Nationality 1',
        created_at: new Date(),
        updated_at: new Date(),
      },
      {
        biography: 'Biography 2',
        nationality: 'Nationality 2',
        created_at: new Date(),
        updated_at: new Date(),
      },
      // Add more author data as needed
    ]);

    await queryInterface.bulkInsert('shelves', [
      {
        floor: 1,
        section: 1,
        book_capacity: 10,
        created_at: new Date(),
        updated_at: new Date(),
      },
      {
        floor: 2,
        section: 2,
        book_capacity: 20,
        created_at: new Date(),
        updated_at: new Date(),
      },
      // Add more shelf data as needed
    ]);

    await queryInterface.bulkInsert('categories', [
      {
        name: 'Category 1',
        description: 'Category description 1',
        created_at: new Date(),
        updated_at: new Date(),
      },
      {
        name: 'Category 2',
        description: 'Category description 2',
        created_at: new Date(),
        updated_at: new Date(),
      },
      // Add more category data as needed
    ]);

    await queryInterface.bulkInsert('books', [
      {
        isbn: 'ISBN1',
        title: 'Book Title 1',
        language: 'Language 1',
        total_copy: 5,
        description: 'Description 1',
        shelf_id: 1, // Assuming shelf_id references an existing shelf
        category_id: 1, // Assuming category_id references an existing category
        created_at: new Date(),
        updated_at: new Date(),
      },
      {
        isbn: 'ISBN2',
        title: 'Book Title 2',
        language: 'Language 2',
        total_copy: 10,
        description: 'Description 2',
        shelf_id: 2, // Assuming shelf_id references an existing shelf
        category_id: 2, // Assuming category_id references an existing category
        created_at: new Date(),
        updated_at: new Date(),
      },
      // Add more book data as needed
    ]);
  },

  async down(queryInterface, Sequelize) {
    // Add logic to revert seed data if needed
    await queryInterface.bulkDelete('authors', null, {});
    await queryInterface.bulkDelete('shelves', null, {});
    await queryInterface.bulkDelete('categories', null, {});
    await queryInterface.bulkDelete('books', null, {});
  },
};
