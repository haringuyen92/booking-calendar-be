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
        res.status(400).json({ 
            success: false,
            message: error.message
         });
    }
}

exports.getUser = async (req, res, next) => {
    const userId = req.params.id;
    try {
        const user = await User.findById(userId);
        res.status(200).json({
            success: true, 
            message: `success getUser ${userId}`,
            user: user
        });
    } catch(error){
        res.status(400).json({ 
            success: false, 
            message: error.message,
            user_id: userId
        });
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
        res.status(400).json({ 
            success: false, 
            message: error.message
        });
    }
}

exports.updateUser = async (req, res, next) => {
    const userId = req.params.id;
    try {
        const user = await User.findByIdAndUpdate(userId, req.body, {
            new: true,
            runValidators: true
        });
        res.status(200).json({ 
            success: true, 
            message: `success updateUser ${userId}`,
            user: user
        });    
    } catch (error) {
        res.status(400).json({
            success: false,
            message: error.message
        })
    }
    
}

exports.deleteUser = async (req, res, next) => {
    const userId = req.params.id;
    try {
        const user = await User.findByIdAndDelete(userId);
        res.status(200).json({ 
            success: true, 
            message: `success deleteUser ${userId}`,
            user: user
        });        
    } catch (error) {
        res.status(200).json({ 
            success: false, 
            message: error.message
        });        
    }

}