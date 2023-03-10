const express = require('express');
const settingRouter = express.Router({ mergeParams: true });
const { protect, authorize } = require('../middleware/auth');
const {getSettingSlot, postSettingSlot} = require("../controllers/settingController");

settingRouter
    .route('/settingSlot')
    .get(protect, authorize('store'), getSettingSlot)
    .post(protect, authorize('store'), postSettingSlot);
settingRouter
    .route('/settingTime')
    .get(protect, authorize('store', 'admin'), )
module.exports = settingRouter;
