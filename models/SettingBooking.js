const mongoose = require('mongoose');

const Staff = mongoose.Schema({
    store: {
        type: mongoose.Types.ObjectId,
        ref: 'Store',
        required: [true, 'Store invalid!']
    },
    durations: {
        type: Number,
        default: 10
    },
    maxBookingSlot: {
        type: Number,
        default: 0
    },
    isRequiredApprove: {
        type: Number,
        default: 2
    },
    isRequiredChange: {
        type: Number,
        default: 2
    },
    changeWithBeforeTime: {
        type: Number,
        default: 0
    },
    isRequiredCancel: {
        type: Number,
        default: 2
    },
    cancelWithBeforeTime: {
        type: Number,
        default: 0
    },
    createdAt: {
        type: Date,
        default: Date.now
    },
})

module.exports = mongoose.model('Staff', Staff);