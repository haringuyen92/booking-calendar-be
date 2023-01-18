const jwt = require('jsonwebtoken');
const errorResponse = require('../utils/errorResponse');
const User = require('../models/User');
const asyncHandler = require('./async');

exports.protect = asyncHandler(async (req, res, next) => {
    let token;
    if(req.headers.authorization && req.headers.authorization.startsWith('Bearer')){
        token = req.headers.authorization.split(' ')[1];
    }

    if(!token) return next (new errorResponse(`Not authorize to access this route`, 401));

    try {
        const decoded = jwt.verify(token, process.env.JWT_SECRET);        
        const user = User.findById(decoded.id);
        if(!user) return next(new errorResponse(`verify token is failed!`, 401));
        req.user = user;
    } catch (error) {
        return next(new errorResponse(`Token invalid`, 404));
    }
    next();
})