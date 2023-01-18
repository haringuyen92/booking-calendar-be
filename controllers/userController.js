const errorResponse = require('../utils/errorResponse');
const asyncHandler = require('../middleware/async');
const Users = require('../models/User');

// @desc   getAll 
// @route  GET /api/users
// @access Private

exports.getUsers = asyncHandler(async (req, res, next) => {
    const reqQuery = {...req.query};
    const removeFields = ['select', 'sort'];
    removeFields.forEach(q => delete reqQuery[q]);

    //create operators ($gt,...)
    // queryStr = queryStr.replace(/\b(lt|lte|gt|gte|ne)\b/, m => `$${m}`);
    let query = Users.find(reqQuery);

    //select fields ex: select=s1,s2,...
    if(req.query.select){
        const selectFields = req.query.select.split(',').join(' ');
        query.select(selectFields);
    }
    //sort ex: ASC sort=s1,s2,s3,...  - DESC sort=-s1,-s2,-s3,...
    if(req.query.sort){
        const sortFields = req.query.sort.split(',').join(' ');
        query.sort(sortFields);
    }else{
        query.sort('-createdAt');
    }
    const total = await Users.countDocuments();
    //pagination
    const page = parseInt(req.query.page, 10) || 1;
    const limit = parseInt(req.query.limit, 10) || 10;
    const skip = (page - 1)*limit;

    query.skip(skip).limit(limit);

    const users = await query.populate('stores');
    res.status(200).json({ 
        success: true, 
        message: "success getUsers",
        total: total,
        count: users.length,
        users: users
    }); 
})

// @desc   getOne 
// @route  GET /api/users/:id
// @access Private


exports.getUser = asyncHandler(async (req, res, next) => {
    const userId = req.params.id;
    const user = await Users.findById(userId).populate('stores');
    if(!user) return next(new errorResponse(`user not found with ID: ${userId}`, 404))
    res.status(200).json({
        success: true, 
        message: `success getUser ${userId}`,
        user: user
    });
})

// @desc   create 
// @route  POST /api/users
// @access private


exports.createUser = asyncHandler(async (req, res, next) => {
    const user = await Users.create(req.body);
    res.status(200).json({ 
        success: true, 
        message: "success createUser createUser",
        user: user
    });
})

// @desc   update 
// @route  POST /api/users/:id
// @access private


exports.updateUser = asyncHandler(async (req, res, next) => {
    const userId = req.params.id;
    const user = await Users.findByIdAndUpdate(userId, req.body, {
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

// @desc   delete 
// @route  DELETE /api/users/:id
// @access private


exports.deleteUser = asyncHandler(async (req, res, next) => {
    const userId = req.params.id;
    const user = await Users.findById(userId);

    if(!user) return next(new errorResponse(`user not found with ID: ${userId}`, 404));
    user.remove();
    
    res.status(200).json({ 
        success: true, 
        message: `success deleteUser ${userId}`,
        user: user
    });    
})