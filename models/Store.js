const mongoose = require('mongoose');
const geocoder = require('../utils/geocoder');

const StoreSchema = mongoose.Schema({
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
        maxlength: [500, 'Description is more than 500 characters!']
    },
    website: {
        type: String,
        match: [
            /https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)/,
            'Website invalid!'
        ]
    },
    phone: {
        type: String,
        maxlength: [20, 'Phone is more than 20 characters!']
    },
    email: {
        type: String,
        match: [
            /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/,
            'Email invalid!'
        ]
    },
    address: {
        type: String,
        required: [true, 'Address is required!']
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
        ref: 'User',
        required: [true, 'User invalid!']
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
StoreSchema.set('toJSON', {
    transform: function (doc, ret) {
      ret.id = ret._id;
    }
  });

module.exports = mongoose.model('Store', StoreSchema);