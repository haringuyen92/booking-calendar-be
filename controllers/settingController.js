const asyncHandler = require("../middleware/async");
const ErrorResponse = require("../utils/errorResponse");
const StoreService = require("../services/storeServices");
const SettingService = require("../services/settingService");

module.exports.getSettingSlot = asyncHandler(async (req, res, next) => {
    const {id} = req.params;

    const settingSlot = await SettingService.getSettingSlot(id);

    return res.status(200).json({
        success: true,
        message: "success getSettingSlot",
        data: settingSlot
    })
})
module.exports.postSettingSlot = asyncHandler(async (req, res, err) => {
    const {id} = req.params;
    req.body.store = id;

    const settingSlot = await SettingService.postSettingSlot(id, req.body);

    return res.status(200).json({
        success: true,
        message: "success postSettingSlot",
        settingSlot
    })
})