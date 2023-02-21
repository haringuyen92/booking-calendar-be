const express = require('express');
const authRouter = require('./routes/auth');
const userRouter = require('./routes/users');
const storeRouter = require('./routes/stores');
const dotenv = require('dotenv');
// const logger = require('./middleware/logger');
const morgan = require('morgan');
const colors = require('colors');
const cors = require('cors');
const cookieParser = require('cookie-parser');
const errorHandler = require('./middleware/error');
const connectDB = require('./config/db');

dotenv.config({
    path: './config/config.env'
});
const corsOptions = {
    origin: [
        "http://localhost:8080"
    ],
    credentials: true,
}

const app = express();
const PORT = process.env.PORT || 3000;
connectDB();

app.use(express.json());
app.use(cookieParser());
app.use(cors(corsOptions));


if(process.env.NODE_ENV === 'dev'){
    app.use(morgan('dev'));
}

app.use('/api/auth', authRouter);
app.use('/api/users', userRouter);
app.use('/api/stores', storeRouter);


app.use(errorHandler);

const server = app.listen(PORT, () => {
    console.log(`Server Started at ${PORT}`.blue.bold);
})

process.on('unhandledRejection', (err, promise) => {
    console.log(`Error: ${err.message}`.red.bold);
    server.close(() => process.exit(1));
})