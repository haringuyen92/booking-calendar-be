const errorResponse = require('../utils/errorResponse');
const asyncHandler = require('../middleware/async');
const User = require('../models/Users');

exports.getUsers = asyncHandler(async (req, res, next) => {
    const users = await User.find();
    res.status(200).json({ 
        success: true, 
        message: "success getUsers",
        users: users
    }); 
})

exports.getUser = asyncHandler(async (req, res, next) => {
    const userId = req.params.id;
    const user = await User.findById(userId);
    if(!user) return next(new errorResponse(`user not found with ID: ${userId}`, 404))
    res.status(200).json({
        success: true, 
        message: `success getUser ${userId}`,
        user: user
    });
})

exports.createUser = asyncHandler(async (req, res, next) => {
    const user = await User.create(req.body);
    res.status(200).json({ 
        success: true, 
        message: "success createUser createUser",
        user: user
    });
})

exports.updateUser = asyncHandler(async (req, res, next) => {
    const userId = req.params.id;
    const user = await User.findByIdAndUpdate(userId, req.body, {
        new: true,
        runValidators: true
    });

    if(!user) return next(new errorResponse(`user not found with ID: ${userId}`, 404))

    res.status(200).json({ 
        success: true, 
        message: `success updateUser ${userId}`,
        user: user
    });    
})

exports.deleteUser = asyncHandler(async (req, res, next) => {
    const userId = req.params.id;
    const user = await User.findByIdAndDelete(userId);

    if(!user) return next(new errorResponse(`user not found with ID: ${userId}`, 404));
    
    res.status(200).json({ 
        success: true, 
        message: `success deleteUser ${userId}`,
        user: user
    });    
})