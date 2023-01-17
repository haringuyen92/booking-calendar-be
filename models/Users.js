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
},{
    toJSON: { virtuals: true },
    toObject: { virtuals: true }
});

Users.pre('save', function(next){
    this.slug = slugify(this.name, {
        lower: true
    })
    next();
});

Users.virtual('stores', {
    ref: 'Store',
    localField: '_id',
    foreignField: 'user',
    justOne: false
})
module.exports = mongoose.model('Users', Users);