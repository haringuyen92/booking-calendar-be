const errorResponse = require('../utils/errorResponse');
const asyncHandler = require('../middleware/async');
const User = require('../models/User');
const bcryptjs = require('bcryptjs');


// @desc Register
// @route POST /api/auth/register
// @access Public

exports.register = asyncHandler( async(req, res, next) => {
    const user = await User.create(req.body);
    sendTokenResponse(user, res);
})

// @desc Login
// @route POST /api/auth/login
// @access Public

exports.login = asyncHandler( async(req, res, next) => {
    const { email, password } = { ...req.body };

    if(!email || !password) return next(new errorResponse(`missing params email or password`, 404));

    const user = await User.findOne({ email:email }).select('+password');

    if(!user) return next(new errorResponse(`Invalid credentials`, 401));

    const isMatch = await user.matchPassword(password);

    if(!isMatch) return next(new errorResponse(`Invalid password`, 401));

    sendTokenResponse(user, res);
})

const sendTokenResponse = (user, res) => {
    const token = user.getSignedJwtToken();

    const options = {
        expires: new Date(Date.now() + process.env.COOKIE_TOKEN_EXPIRE*24*3600*1000),
        httpOnly: true,
    }
    if(process.env.NODE_ENV === 'production'){
        options.secure = true;
    }

    res.cookie('token', token, options)
        .status(200)
        .json({
            success: true,
            user: {
                id: user.id,
                name: user.name,
                email: user.email,
                role: user.role,
                token: token,
            }
        })
}