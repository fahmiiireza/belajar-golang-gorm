import { DataTypes, Model } from 'sequelize';
import sequelize from '../../sequelize';
import Shelf from './shelf';
import Category from './category';
import Author from './author'; // Import Author model

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
    },
    language: {
      type: DataTypes.STRING,
      allowNull: false,
    },
    totalCopy: {
      type: DataTypes.INTEGER,
      allowNull: false,
      field: 'total_copy',
      validate: {
        notZero(value: number) {
          if (value === 0) {
            throw new Error('totalCopy cannot be 0');
          }
        },
      },
    },
    shelfId: {
      type: DataTypes.INTEGER,
      field: 'shelf_id',
    },
    categoryId: {
      type: DataTypes.INTEGER,
      field: 'category_id',
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
    paranoid: true, // Enable soft deletes
    tableName: 'books', // Specify the table name if different from the model name
    underscored: true, // Enable snake_case column names
  }
);

// Define associations
Book.belongsTo(Shelf, { foreignKey: 'shelfId' }); // Many-to-one with Shelf
Book.belongsTo(Category, { foreignKey: 'categoryId' }); // Many-to-one with Category
Book.belongsToMany(Author, { through: 'author_books', foreignKey: 'book_id' }); // Many-to-many with Author

export default Book;
