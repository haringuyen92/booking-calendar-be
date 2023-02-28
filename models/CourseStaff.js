const mongoose = require('mongoose');

const Staff = mongoose.Schema({
    staff: {
        type: mongoose.Types.ObjectId,
        ref: 'Staff',
        required: [true, 'Staff invalid!']
    },
    course: {
        type: mongoose.Types.ObjectId,
        ref: 'Course',
        required: [true, 'Course invalid!']
    },
    createdAt: {
        type: Date,
        default: Date.now
    },
})

module.exports = mongoose.model('Staff', Staff);