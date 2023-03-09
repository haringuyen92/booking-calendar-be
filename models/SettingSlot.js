const mongoose = require("mongoose");

const SettingSlot = mongoose.Schema({
    store: {
        type: mongoose.Types.ObjectId,
        ref: 'Store',
        required: [true, 'Store invalid!']
    },
    isRequiredStaff: {
        type: String,
        default: 'optional' //not_used, required, optional
    },
    isUseCostStaff: {
        type: Boolean,
        default: false
    },
    isRequiredCourse: {
        type: String,
        default: 'optional' //not_used, required, optional
    },
    isUseCostCourse: {
        type: Boolean,
        default: false
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