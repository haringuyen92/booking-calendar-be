const express = require('express');
const { getStaffs, getStaff, createStaff, updateStaff, deleteStaff } = require('../controllers/staffController');
const staffRouter = express.Router({ mergeParams: true });
const { protect, authorize } = require('../middleware/auth');

staffRouter
    .route('/')
    .get(protect, getStaffs)
    .post(protect, authorize('store', 'admin'), createStaff);
    
    staffRouter
    .route('/:id')
    .get(getStaff)
    .put(protect, authorize('store', 'admin'), updateStaff)
    .delete(protect, authorize('store', 'admin'), deleteStaff);

module.exports = staffRouter;
