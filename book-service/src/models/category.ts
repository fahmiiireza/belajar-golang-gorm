import { DataTypes, Model } from 'sequelize';
import sequelize from '../../sequelize';

class Category extends Model {
  public id!: number;
  public name!: string;
  public description?: string;
  public created_at!: Date;
  public updated_at!: Date;
  public deleted_at?: Date;
}

Category.init(
  {
    id: {
      type: DataTypes.BIGINT,
      primaryKey: true,
      autoIncrement: true,
    },
    name: {
      type: DataTypes.STRING,
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
  },
  {
    sequelize,
    modelName: 'Category',
    timestamps: true,
    paranoid: true,
    tableName: 'categories',
    underscored: true,
  }
);

export default Category;
