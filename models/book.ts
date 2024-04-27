// File: models/book.js
import { DataTypes, Model, Optional } from 'sequelize';
import sequelize from '../sequelize';
import Shelf from './shelf';
import Category from './category';  

const Book = sequelize.define('Book', {
  id: {
    type: DataTypes.BIGINT,
    primaryKey: true,
    autoIncrement: true,
  },
  isbn: {
    type: DataTypes.STRING,
    allowNull: false,
    unique: true,
  },
  title: {
    type: DataTypes.STRING,
    allowNull: false,
  },
  language: {
    type: DataTypes.STRING,
    allowNull: false,
  },
  total_copy: {
    type: DataTypes.INTEGER,
    allowNull: false,
  },
  description: {
    type: DataTypes.STRING,
  },
  created_at: {
    type: DataTypes.DATE,
    defaultValue: DataTypes.NOW,
    allowNull: false,
  },
  updated_at: {
    type: DataTypes.DATE,
    defaultValue: DataTypes.NOW,
    allowNull: false,
  },
  deleted_at: {
    type: DataTypes.DATE,
  },
});



// Define associations
Book.belongsTo(Shelf, { foreignKey: 'shelf_id' });
Shelf.hasMany(Book, { foreignKey: 'shelf_id' });

Book.belongsTo(Category, { foreignKey: 'category_id' });
Category.hasMany(Book, { foreignKey: 'category_id' });

module.exports = { Book };
