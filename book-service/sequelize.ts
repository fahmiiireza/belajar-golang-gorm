// sequelize.ts

import { Sequelize } from 'sequelize';

const sequelize = new Sequelize('book-db', 'postgres', 'postgres', {
    port: 5434,
    host: 'host.docker.internal',
    password: 'postgres',
    database: 'book-db',
    dialect: 'postgres',
    username: 'postgres'
});

export default sequelize;
