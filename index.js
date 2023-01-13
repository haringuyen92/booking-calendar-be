const express = require('express');
const userRoutes = require('./routes/usersRoutes');
const dotenv = require('dotenv');
// const logger = require('./middleware/logger');
const morgan = require('morgan');
const colors = require('colors');
const connectDB = require('./config/db');

dotenv.config({
    path: './config/config.env'
});

const app = express();
const PORT = process.env.PORT || 3000;
connectDB();

app.use(express.json());

if(process.env.NODE_ENV === 'dev'){
    app.use(morgan('dev'));
}


app.use('/api/users', userRoutes);

const server = app.listen(PORT, () => {
    console.log(`Server Started at ${PORT}`.blue.bold);
})

process.on('unhandledRejection', (err, promise) => {
    console.log(`Error: ${err.message}`.red.bold);
    server.close(() => process.exit(1));
})