// File: models/author.js
import { DataTypes, Model, Optional } from 'sequelize';
import sequelize from '../sequelize';

const Author = sequelize.define('Author', {
  id: {
    type: DataTypes.BIGINT,
    primaryKey: true,
    autoIncrement: true,
  },
  biography: {
    type: DataTypes.STRING,
  },
  nationality: {
    type: DataTypes.STRING,
    allowNull: false,
  },
});

export default Author;