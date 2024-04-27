'use strict';

module.exports = {
  up: async (queryInterface, Sequelize) => {
    // Create Author table
    await queryInterface.createTable('authors', {
      id: {
        type: Sequelize.DataTypes.BIGINT,
        primaryKey: true,
        autoIncrement: true,
      },
      biography: {
        type: Sequelize.DataTypes.STRING,
      },
      nationality: {
        type: Sequelize.DataTypes.STRING,
        allowNull: false,
      },
      created_at: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      updated_at: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      deleted_at: {
        type: Sequelize.DataTypes.DATE,
      },
    });

    // Create Shelf table
    await queryInterface.createTable('shelves', {
      id: {
        type: Sequelize.DataTypes.BIGINT,
        primaryKey: true,
        autoIncrement: true,
      },
      floor: {
        type: Sequelize.DataTypes.INTEGER,
        allowNull: false,
      },
      section: {
        type: Sequelize.DataTypes.INTEGER,
        allowNull: false,
      },
      book_capacity: {
        type: Sequelize.DataTypes.INTEGER,
        allowNull: false,
      },
      created_at: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      updated_at: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      deleted_at: {
        type: Sequelize.DataTypes.DATE,
      },
    });

    // Create Category table
    await queryInterface.createTable('categories', {
      id: {
        type: Sequelize.DataTypes.BIGINT,
        primaryKey: true,
        autoIncrement: true,
      },
      name: {
        type: Sequelize.DataTypes.STRING,
        allowNull: false,
      },
      description: {
        type: Sequelize.DataTypes.STRING,
      },
      created_at: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      updated_at: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      deleted_at: {
        type: Sequelize.DataTypes.DATE,
      },
    });
    // Create Book table
    await queryInterface.createTable('books', {
      id: {
        type: Sequelize.DataTypes.BIGINT,
        primaryKey: true,
        autoIncrement: true,
      },
      isbn: {
        type: Sequelize.DataTypes.STRING,
        allowNull: false,
        unique: true,
      },
      title: {
        type: Sequelize.DataTypes.STRING,
        allowNull: false,
      },
      language: {
        type: Sequelize.DataTypes.STRING,
        allowNull: false,
      },
      total_copy: {
        type: Sequelize.DataTypes.INTEGER,
        allowNull: false,
      },
      description: {
        type: Sequelize.DataTypes.STRING,
      },
      created_at: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      updated_at: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      deleted_at: {
        type: Sequelize.DataTypes.DATE,
      },
    });

    // Define associations
    await queryInterface.addColumn('books', 'shelf_id', {
      type: Sequelize.DataTypes.BIGINT,
      references: {
        model: 'shelves',
        key: 'id',
      },
    });

    await queryInterface.addColumn('books', 'category_id', {
      type: Sequelize.DataTypes.BIGINT,
      references: {
        model: 'categories',
        key: 'id',
      },
    });

    // Add foreign key constraints
    await queryInterface.addConstraint('books', {
      fields: ['shelf_id'],
      type: 'foreign key',
      name: 'fk_books_shelf_id',
      references: {
        table: 'shelves',
        field: 'id',
      },
      onDelete: 'cascade',
      onUpdate: 'cascade',
    });

    await queryInterface.addConstraint('books', {
      fields: ['category_id'],
      type: 'foreign key',
      name: 'fk_books_category_id',
      references: {
        table: 'categories',
        field: 'id',
      },
      onDelete: 'cascade',
      onUpdate: 'cascade',
    });
  },

  down: async (queryInterface, Sequelize) => {
    // Drop tables in reverse order
    await queryInterface.dropTable('books');
    await queryInterface.dropTable('shelves');
    await queryInterface.dropTable('categories');
    await queryInterface.dropTable('authors');
  },
};
