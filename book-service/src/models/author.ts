import { DataTypes, Model } from 'sequelize';
import sequelize from '../../sequelize';

class Author extends Model {
  public id!: number;
  public biography?: string;
  public nationality!: string;
}

Author.init(
  {
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
  },
  {
    sequelize,
    modelName: 'Author',
    timestamps: true,
    paranoid: true,
    tableName: 'authors',
    underscored: true,
  }
);

export default Author;
