const mongoose = require('mongoose');

const Users = new mongoose.Schema({
    username: {
        type: String,
        unique: true,
        maxlength: 50,
        require: true
    },
    name: String,
    email: {
        type: String,
        unique: true,
        require: true
    },
    createdAt: {
        type: Date,
        default: Date.now
    }
});

module.exports = mongoose.model('User', Users);