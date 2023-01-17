const mongoose = require('mongoose');
const slugify = require('slugify');

const UserSchema = new mongoose.Schema({
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

UserSchema.pre('save', async function(next){
    this.slug = slugify(this.name, {
        lower: true
    });
    next();
});

UserSchema.pre('remove', async function(next){
    console.log(`Remove Store from User: ${this._id}`);
    await this.model('Store').deleteMany({ user: this._id });
    next();
})

UserSchema.virtual('stores', {
    ref: 'Store',
    localField: '_id',
    foreignField: 'user',
    justOne: false
})
module.exports = mongoose.model('User', UserSchema);