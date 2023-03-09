const express = require('express');
const settingRouter = express.Router({ mergeParams: true });
const { protect, authorize } = require('../middleware/auth');
const {settingSlot} = require("../controllers/settingController");

settingRouter
    .route('/settingSlot')
    .post(protect, authorize('store'), settingSlot);

module.exports = settingRouter;
