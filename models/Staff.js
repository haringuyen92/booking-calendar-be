const mongoose = require('mongoose');

const Staff = mongoose.Schema({
    store: {
        type: mongoose.Types.ObjectId,
        ref: 'Store',
        required: [true, 'Store invalid!']
    },
    name: {
        type: String,
        trim: true,
        required: [true, 'Name is required'],
        maxlength: [30, 'Name is more than 30 characters!']
    },
    image: {
        type: String,
        default: 'no-photo.jpg'
    },
    description: {
        type: String,
        required: [true, 'Description is required'],
        maxlength: [1500, 'Description is more than 500 characters!']
    },
    cost: {
        type: Number,
        default: 0
    },
    maxBookingSlot: {
        type: Number,
        default: 0
    },
    active: {
        type: Boolean,
        default: false
    },
    color: {
        type: String,
        default: 'black'
    },
    position: {
        type: Number,
        default: 1
    },
    createdAt: {
        type: Date,
        default: Date.now
    },
})

module.exports = mongoose.model('Staff', Staff);