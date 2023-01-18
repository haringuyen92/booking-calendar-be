const express = require('express');
const { getUsers, getUser, createUser, updateUser, deleteUser } = require('../controllers/userController');

const userRouter = express.Router();
const courseRouter = require('./stores');
const { protect, authorize } = require('../middleware/auth');

userRouter.use('/:userId/stores', courseRouter);

userRouter
    .route('/')
    .get(protect, getUsers)
    .post(protect, authorize('admin', 'store'), createUser);
    
userRouter
    .route('/:id')
    .get(protect, getUser)
    .put(protect, authorize('admin', 'store'), updateUser)
    .delete(protect, authorize('admin', 'store'), deleteUser);

module.exports = userRouter;