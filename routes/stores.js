const express = require('express');
const { getStores, getStore, createStore, updateStore, deleteStore } = require('../controllers/storeController');
const storeRouter = express.Router({ mergeParams: true });
const { protect, authorize } = require('../middleware/auth');

storeRouter
    .route('/')
    .get(getStores)
    .post(protect, authorize('store', 'admin'), createStore);
    
storeRouter
    .route('/:id')
    .get(getStore)
    .put(protect, authorize('store', 'admin'), updateStore)
    .delete(protect, authorize('store', 'admin'), deleteStore);

module.exports = storeRouter;
