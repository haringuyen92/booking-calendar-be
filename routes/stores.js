const express = require('express');
const { getStores, getStore, createStore, updateStore, deleteStore } = require('../controllers/storeController');
const storeRouter = express.Router({ mergeParams: true });

storeRouter
    .route('/')
    .get(getStores)
    .post(createStore);
    
storeRouter
    .route('/:id')
    .get(getStore)
    .put(updateStore)
    .delete(deleteStore);

module.exports = storeRouter;
