import { DataTypes, Model } from 'sequelize';
import sequelize from '../../sequelize';

class Shelf extends Model {
  public id!: number;
  public floor!: number;
  public section!: number;
  public bookCapacity!: number;
  public created_at!: Date;
  public updated_at!: Date;
  public deleted_at?: Date;
}

Shelf.init(
  {
    id: {
      type: DataTypes.BIGINT,
      primaryKey: true,
      autoIncrement: true,
    },
    floor: {
      type: DataTypes.INTEGER,
      allowNull: false,
    },
    section: {
      type: DataTypes.INTEGER,
      allowNull: false,
    },
    book_capacity: {
      type: DataTypes.INTEGER,
      allowNull: false,
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
  },
  {
    sequelize,
    modelName: 'Shelf',
    timestamps: true,
    paranoid: true,
    tableName: 'shelves',
    underscored: true,
  }
);

export default Shelf;
