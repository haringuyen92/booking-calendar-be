const errorResponse = require('../utils/errorResponse');
const asyncHandler = require('../middleware/async');
const User = require('../models/User');
const Store = require('../models/Store');


// @desc   getAll
// @route  GET /api/stores
// @route  GET /api/users/:userId/stores
// @access Public

exports.getStores = asyncHandler( async(req, res, next) => {
    const reqQuery = {...req.query};
    
    const removeFields = ['select', 'sort'];
    removeFields.forEach(q => delete reqQuery[q]);

    //create operators ($gt,...)
    // queryStr = queryStr.replace(/\b(gt)\b/, m => `$${m}`);
    if(req.params.userId && req.user.role === 'store'){
        reqQuery.user = req.params.userId;
    }
    let query = Store.find(reqQuery).populate('user');


    const stores = await query;
    res.status(200).json({ 
        success: true, 
        message: "success getStores",
        count: stores.length,
        data: stores
    }); 
})

// @desc   getOne
// @route  GET /api/stores/:id
// @access Public

exports.getStore = asyncHandler( async(req, res, next) => {
    const storeId = req.params.id;
    const store = await Store.findById(storeId).populate({
        path: 'user'
    });

    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    res.status(200).json({ 
        success: true, 
        message: `success getStore`,
        data: store
    }); 
})

// @desc   create
// @route  POST /api/users/:userId/store
// @access Private

exports.createStore = asyncHandler( async(req, res, next) => {
    const userId = req.params.userId;
    const user = User.findById(userId);

    if(!user) return next(new errorResponse(`User not found id: ${userId}`, 404));

    req.body.user = userId;
    const store = await Store.create(req.body);
    
    res.status(200).json({ 
        success: true, 
        message: `success createStore`,
        data: store
    }); 
})

// @desc   update 
// @route  PUT /api/stores/:id
// @access Private

exports.updateStore = asyncHandler( async(req, res, next) => {
    const storeId = req.params.id;
    let store = await Store.findById(storeId);

    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    store = await Store.findByIdAndUpdate(storeId, req.body, {
        new: true,
        runValidators: true
    });

    res.status(200).json({ 
        success: true, 
        message: `success updateStore ${storeId}`,
        data: store
    }); 
})

// @desc   delete 
// @route  DELETE /api/stores/:id
// @access Private

exports.deleteStore = asyncHandler( async(req, res, next) => {
    const storeId = req.params.id;
    let store = await Store.findById(storeId);
    
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    await store.remove();

    res.status(200).json({ 
        success: true, 
        message: `success deleteStore ${storeId}`,
        data: store
    }); 
})