const mongoose = require('mongoose');
const bcryptjs = require('bcryptjs');
const jwt = require('jsonwebtoken');

const UserSchema = new mongoose.Schema({
    name: {
        type: String,
        required: [true, 'name is require']
    },
    email: {
        type: String,
        unique: true,
        required: [true, 'email is require'],
        match: [
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
            'email invalid'
        ]
    },
    role: {
        type: String,
        enum: ['admin', 'user', 'store'],
        default: 'user'
    },
    password: {
        type: String,
        required: [true, 'password is require'],
        minlength: 6,
        select: false
    },
    resetPasswordToken: String,
    resetPasswordTExpire: Date,
    createdAt: {
        type: Date,
        default: Date.now
    }
},{
    toJSON: { virtuals: true },
    toObject: { virtuals: true }
});

UserSchema.pre('save', async function(next){
    const salt = await bcryptjs.genSalt(10);
    this.password = await bcryptjs.hash(this.password, salt);
    next();
});

UserSchema.pre('remove', async function(next){
    console.log(`Remove Store from User: ${this._id}`);
    await this.model('Store').deleteMany({ user: this._id });
    next();
})

UserSchema.methods.getSignedJwtToken = function() {
    return jwt.sign({ id: this._id }, process.env.JWT_SECRET, {
        expiresIn: process.env.JWT_EXPIRE
    });
}

UserSchema.methods.matchPassword = async function(password) {
    return await bcryptjs.compare(password, this.password);
}

UserSchema.virtual('stores', {
    ref: 'Store',
    localField: '_id',
    foreignField: 'user',
    justOne: false
})
module.exports = mongoose.model('User', UserSchema);