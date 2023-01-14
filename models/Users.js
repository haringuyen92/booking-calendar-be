const mongoose = require('mongoose');

const Users = new mongoose.Schema({
    username: {
        type: String,
        unique: true,
        maxlength: 50
    },
    name: String,
    email: {
        type: String,
    },
    createdAt: {
        type: Date,
        default: Date.now
    }
});

module.exports = mongoose.model('User', Users);