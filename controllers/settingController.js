const asyncHandler = require("../middleware/async");
const SettingService = require("../services/settingService");

module.exports.getSettingSlot = asyncHandler(async (req, res, next) => {
    const {storeId} = req.params;

    const settingSlot = await SettingService.getSettingSlot(storeId);
    return res.status(200).json({
        success: true,
        message: "success getSettingSlot",
        data: settingSlot
    })
})
module.exports.postSettingSlot = asyncHandler(async (req, res, err) => {
    const {storeId} = req.params;
    req.body.store = storeId;

    const settingSlot = await SettingService.postSettingSlot(storeId, req.body);
    return res.status(200).json({
        success: true,
        message: "success postSettingSlot",
        data: settingSlot
    })
})
module.exports.getSettingTime = asyncHandler(async (req, res, err) => {
    const {storeId} = req.params;

    const settingTime = SettingService.getSettingTime(storeId);
    return res.status(200).json({
        success: true,
        message: "success getSettingTime",
        data: settingTime
    })
})
module.exports.postSettingTime = asyncHandler(async (req, res, err) => {
    const {storeId} = req.params;
    req.body.store = storeId;

    const settingTime = await SettingService.postSettingTime(storeId, req.body);
    return res.status(200).json({
        success: true,
        message: "success getSettingTime",
        data: settingTime
    })
})