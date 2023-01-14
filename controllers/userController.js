const errorResponse = require('../utils/errorResponse');
const User = require('../models/Users');

exports.getUsers = async (req, res, next) => {
    try {
        const users = await User.find();
        res.status(200).json({ 
            success: true, 
            message: "success getUsers",
            users: users
        });        
    } catch (error) {
        next(error);
    }
}

exports.getUser = async (req, res, next) => {
    const userId = req.params.id;
    try {
        const user = await User.findById(userId);
        if(!user) return next(new errorResponse(`user not found with ID: ${userId}`, 404))
        res.status(200).json({
            success: true, 
            message: `success getUser ${userId}`,
            user: user
        });
    } catch(error){
        next(error);
    }
}

exports.createUser = async (req, res, next) => {
    try {
        const user = await User.create(req.body);
        res.status(200).json({ 
            success: true, 
            message: "success createUser createUser",
            user: user
        });
    } catch(error) {
        next(error);
    }
}

exports.updateUser = async (req, res, next) => {
    const userId = req.params.id;
    try {
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
    } catch (error) {
        next(error);
    }
    
}

exports.deleteUser = async (req, res, next) => {
    const userId = req.params.id;
    try {
        const user = await User.findByIdAndDelete(userId);

        if(!user) return next(new errorResponse(`user not found with ID: ${userId}`, 404));
        
        res.status(200).json({ 
            success: true, 
            message: `success deleteUser ${userId}`,
            user: user
        });        
    } catch (error) {
        next(error);
    }

}