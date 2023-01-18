const express = require('express');
const { getUsers, getUser, createUser, updateUser, deleteUser } = require('../controllers/userController');

const userRouter = express.Router();
const courseRouter = require('./stores');
const { protect } = require('../middleware/auth');

userRouter.use('/:userId/stores', courseRouter);

userRouter
    .route('/')
    .get(getUsers)
    .post(protect, createUser);
    
userRouter
    .route('/:id')
    .get(getUser)
    .put(protect, updateUser)
    .delete(protect, deleteUser);

module.exports = userRouter;