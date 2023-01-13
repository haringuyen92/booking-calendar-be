const express = require('express');
const mongoose = require('mongoose');
const routes = require('./routes/routes');
const dotenv = require('dotenv');

dotenv.config({
    path: './config/config.env'
});

const app = express();
const PORT = process.env.PORT || 3000;

app.use(express.json());

app.use(function(req, res, next) {
    res.setHeader('X-Powered-By', 'Hainv');
    next();
  });

app.use('/', (req,res) => res.send('<h1>hello world!!</h1>'));
app.use('/api', routes);

app.listen(PORT, () => {
    console.log(`Server Started at ${PORT}`)
})