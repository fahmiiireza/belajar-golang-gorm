import { DataTypes, Model } from 'sequelize';
import sequelize from '../../sequelize';

class Book extends Model {
  public id!: number;
  public isbn!: string;
  public title!: string;
  public language!: string;
  public totalCopy!: number;
  public shelfId?: number;
  public categoryId?: number;
  public description?: string;
  public createdAt!: Date;
  public updatedAt!: Date;
  public deletedAt?: Date;

  // public readonly authors?: Author[]; // Define the association with Author
  // public readonly shelf?: Shelf; // Define the association with Shelf
  // public readonly category?: Category; // Define the association with Category
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
        notZero(value : number) {
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
    createdAt: {
      type: DataTypes.DATE,
      allowNull: false,
      defaultValue: DataTypes.NOW,
      field: 'created_at',
    },
    updatedAt: {
      type: DataTypes.DATE,
      allowNull: false,
      defaultValue: DataTypes.NOW,
      field: 'updated_at',
    },
    deletedAt: {
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
  }
);

// Define associations
// Book.belongsToMany(Author, { through: 'author_books', foreignKey: 'book_id' }); // Many-to-many with Author
// Book.belongsTo(Shelf, { foreignKey: 'shelf_id' }); // One-to-one with Shelf
// Book.belongsTo(Category, { foreignKey: 'category_id' }); // One-to-one with Category

export default Book;