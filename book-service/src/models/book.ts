import { DataTypes, Model } from 'sequelize';
import sequelize from '../../sequelize';
import Shelf from './shelf';
import Category from './category';
import Author from './author'; 

class Book extends Model {
  public id!: number;
  public isbn!: string;
  public title!: string;
  public language!: string;
  public totalCopy!: number;
  public shelfId?: number;
  public categoryId?: number;
  public description?: string;
  public created_at!: Date;
  public updated_at!: Date;
  public deleted_at?: Date;
}

Book.init(
  {
    id: {
      type: DataTypes.BIGINT,
      autoIncrement: true,
      primaryKey: true,
    },
    isbn: {
      type: DataTypes.STRING,
      allowNull: false,
      unique: true,
    },
    title: {
      type: DataTypes.STRING,
      allowNull: false,
      validate: {
        notEmpty: true,
      },
    },
    language: {
      type: DataTypes.STRING,
      allowNull: false,
      validate: {
        notEmpty: true,
      },
    },
    totalCopy: {
      type: DataTypes.INTEGER,
      allowNull: false,
      field: 'total_copy',
      validate: {
        notZero(value: number) {
          if (value <= 0) {
            throw new Error('totalCopy cannot be 0');
          }
        },
      },
    },
    shelfId: {
      type: DataTypes.INTEGER,
      field: 'shelf_id',
      validate: {
        isNumeric: true,
      }
    },
    categoryId: {
      type: DataTypes.INTEGER,
      field: 'category_id',
      validate: {
        isNumeric: true,
      }
    },
    description: {
      type: DataTypes.STRING,
    },
    created_at: {
      type: DataTypes.DATE,
      allowNull: false,
      defaultValue: DataTypes.NOW,
      field: 'created_at',
    },
    updated_at: {
      type: DataTypes.DATE,
      allowNull: false,
      defaultValue: DataTypes.NOW,
      field: 'updated_at',
    },
    deleted_at: {
      type: DataTypes.DATE,
      field: 'deleted_at',
    },
  },
  {
    sequelize,
    modelName: 'Book',
    timestamps: true,
    paranoid: true, 
    tableName: 'books',
    underscored: true, 
  }
);

// Define associations
Book.belongsTo(Shelf, { foreignKey: 'shelfId' });
Book.belongsTo(Category, { foreignKey: 'categoryId' }); 
Book.belongsToMany(Author, { through: 'author_books', foreignKey: 'book_id' });

export default Book;
