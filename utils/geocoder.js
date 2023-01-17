const axios = require('axios');
exports.geocode = function(address){
    let mapUrl = 'https://maps.googleapis.com/maps/api/geocode/json';
    return axios.get(`${mapUrl}?address=${address}&key=${process.env.GEOCODER_API_KEY}`)
}
