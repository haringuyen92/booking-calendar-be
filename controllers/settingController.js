const asyncHandler = require("../middleware/async");

module.exports.settingSlot = asyncHandler(async (req, res, err) => {
    return res.status(200).json({
        success: true,
        message: "success createStaff",
    })
})