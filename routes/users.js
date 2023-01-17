const express = require('express');
const { getUsers, getUser, createUser, updateUser, deleteUser } = require('../controllers/userController');

const userRouter = express.Router();
const courseRouter = require('./stores');

userRouter.use('/:userId/stores', courseRouter);

userRouter
    .route('/')
    .get(getUsers)
    .post(createUser);
    
userRouter
    .route('/:id')
    .get(getUser)
    .put(updateUser)
    .delete(deleteUser);

module.exports = userRouter;