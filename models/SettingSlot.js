const mongoose = require("mongoose");

const SettingSlot = mongoose.Schema({
    store: {
        type: mongoose.Types.ObjectId,
        ref: 'Store',
        required: [true, 'Store invalid!']
    },
    isRequiredStaff: {
        type: String,
        default: 'optional' //not_used, used, optional
    },
    isUseCostStaff: {
        type: String,
        default: 'used' //used, not_used
    },
    isRequiredCourse: {
        type: String,
        default: 'optional' //not_used, used, optional
    },
    isUseCostCourse: {
        type: String,
        default: 'used' //used, not_used
    },
    defaultCourseEstimationTime: {
        type: Number,
        default: 15
    },
    createdAt: {
        type: Date,
        default: Date.now
    }
})

module.exports = mongoose.model('SettingSlot', SettingSlot);