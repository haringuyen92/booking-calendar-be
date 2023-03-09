const express = require('express');
const { getStores, getStore, createStore, updateStore, deleteStore } = require('../controllers/storeController');
const storeRouter = express.Router({ mergeParams: true });
const staffRouter = require('./staff');
const courseRouter = require('./course');
const { protect, authorize } = require('../middleware/auth');
const settingRouter = require("./setting");

storeRouter.use('/:storeId/staffs', staffRouter);
storeRouter.use('/:storeId/courses', courseRouter);
storeRouter.use('/:id', settingRouter);
storeRouter
    .route('/')
    .get(protect, getStores)
    .post(protect, authorize('store', 'admin'), createStore);
    
storeRouter
    .route('/:id')
    .get(getStore)
    .put(protect, authorize('store', 'admin'), updateStore)
    .delete(protect, authorize('store', 'admin'), deleteStore);

module.exports = storeRouter;
