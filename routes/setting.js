const express = require('express');
const settingRouter = express.Router({ mergeParams: true });
const { protect, authorize } = require('../middleware/auth');
const {getSettingSlot, postSettingSlot, getSettingTime, postSettingTime} = require("../controllers/settingController");

settingRouter
    .route('/settingSlot')
    .get(protect, authorize('store'), getSettingSlot)
    .post(protect, authorize('store'), postSettingSlot);
settingRouter
    .route('/settingTime')
    .get(protect, authorize('store', 'admin'), getSettingTime)
    .post(protect, authorize('store', 'admin'), postSettingTime)
module.exports = settingRouter;
