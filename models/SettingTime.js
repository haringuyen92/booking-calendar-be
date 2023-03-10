const mongoose = require("mongoose");

const SettingTime = mongoose.Schema({
    store: {
        type: mongoose.Types.ObjectId,
        ref: 'Store',
        required: true
    },
    isOpen: {
        type: Boolean,
        default: false
    },
    isApplyDailySetting: {
        type: Boolean,
        default: false
    },
    dailySetting: {
        type: Object,
        required: true
    },
    mondaySetting: {
        type: Object,
        required: true
    },
    tuesdaySetting: {
        type: Object,
        required: true
    },
    wednesdaySetting: {
        type: Object,
        required: true
    },
    thursdaySetting: {
        type: Object,
        required: true
    },
    fridaySetting: {
        type: Object,
        required: true
    },
    sundaySetting: {
        type: Object,
        required: true
    },
    holidaySetting: {
        type: Object,
        required: true
    }
}, {
    timestamps: true
})

module.exports = mongoose.model('SettingTime', SettingTime)