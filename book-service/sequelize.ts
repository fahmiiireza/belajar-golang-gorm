// sequelize.ts

import { Sequelize } from 'sequelize';

const sequelize = new Sequelize('nodejs-db', 'postgres', 'postgres', {
    port: 5435,
    host: 'host.docker.internal',
    password: 'postgres',
    database: 'nodejs-db',
    dialect: 'postgres',
    username: 'postgres'
});

export default sequelize;
