'use strict';

module.exports = {
  up: async (queryInterface, Sequelize) => {
    // Create Author table
    await queryInterface.createTable('Authors', {
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
      createdAt: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      updatedAt: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      deletedAt: {
        type: Sequelize.DataTypes.DATE,
      },
    });

    // Create Shelf table
    await queryInterface.createTable('Shelves', {
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
      createdAt: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      updatedAt: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      deletedAt: {
        type: Sequelize.DataTypes.DATE,
      },
    });

    // Create Category table
    await queryInterface.createTable('Categories', {
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
      createdAt: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      updatedAt: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      deletedAt: {
        type: Sequelize.DataTypes.DATE,
      },
    });
    // Create Book table
    await queryInterface.createTable('Books', {
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
      createdAt: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      updatedAt: {
        type: Sequelize.DataTypes.DATE,
        allowNull: false,
        defaultValue: Sequelize.DataTypes.NOW,
      },
      deletedAt: {
        type: Sequelize.DataTypes.DATE,
      },
    });

    // Define associations
    await queryInterface.addColumn('Books', 'shelf_id', {
      type: Sequelize.DataTypes.BIGINT,
      references: {
        model: 'Shelves',
        key: 'id',
      },
    });

    await queryInterface.addColumn('Books', 'category_id', {
      type: Sequelize.DataTypes.BIGINT,
      references: {
        model: 'Categories',
        key: 'id',
      },
    });

    // Add foreign key constraints
    await queryInterface.addConstraint('Books', {
      fields: ['shelf_id'],
      type: 'foreign key',
      name: 'fk_books_shelf_id',
      references: {
        table: 'Shelves',
        field: 'id',
      },
      onDelete: 'cascade',
      onUpdate: 'cascade',
    });

    await queryInterface.addConstraint('Books', {
      fields: ['category_id'],
      type: 'foreign key',
      name: 'fk_books_category_id',
      references: {
        table: 'Categories',
        field: 'id',
      },
      onDelete: 'cascade',
      onUpdate: 'cascade',
    });
  },

  down: async (queryInterface, Sequelize) => {
    // Drop tables in reverse order
    await queryInterface.dropTable('Books');
    await queryInterface.dropTable('Shelves');
    await queryInterface.dropTable('Categories');
    await queryInterface.dropTable('Authors');
  },
};
