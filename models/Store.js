const mongoose = require('mongoose');
const geocoder = require('../utils/geocoder');

const StoreSchema = mongoose.Schema({
    name: {
        type: String,
        trim: true,
        require: [true, 'Vua long nhap mo ta'],
        maxlength: [30, 'Chi chap nhan do dai toi da 30']
    },
    image: {
        type: String,
        default: 'no-photo.jpg'
    },
    description: {
        type: String,
        required: [true, 'Vua long nhap mo ta'],
        maxlength: [500, 'Chi chap nhan do dai toi da 500']
    },
    website: {
        type: String,
        match: [
            /https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)/,
            'Dia chi website khong dung dinh dang'
        ]
    },
    phone: {
        type: String,
        maxlength: [20, 'Chi chap nhan do dai toi da 20']
    },
    email: {
        type: String,
        match: [
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
            'Dia chi email khong dung dinh dang'
        ]
    },
    address: {
        type: String,
        require: [true, 'Vui long nhap dia chi']
    },
    location: {
        type: {
            type: String,
            enum: ['Point'], 
        },
        coordinates: {
            type: [Number],
            index: '2dsphere'
        },
        formattedAddress: String,
        street: String,
        city: String,
        state: String,
        zipcode: String,
        country: String
    },
    user: {
        type: mongoose.Types.ObjectId,
        ref: 'Users',
        require: [true, 'User invalid']
    },
    createdAt: {
        type: Date,
        default: Date.now
    },
})

// StoreSchema.pre('save', async function(next) {
//     const location = await geocoder.geocode('Ha Noi')
//         .then((res) => { 
//             console.log(res.data.results[0]);
//         });
//     next();
// })

module.exports = mongoose.model('Store', StoreSchema);