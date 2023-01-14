const mongoose = require('mongoose');
const slugify = require('slugify');

const Users = new mongoose.Schema({
    username: {
        type: String,
        unique: true,
        maxlength: 50,
        require: true
    },
    name: String,
    slug: String,
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
Users.pre('save', function(next){
    this.slug = slugify(this.name, {
        lower: true
    })
    next();
});
module.exports = mongoose.model('User', Users);