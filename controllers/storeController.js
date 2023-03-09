const errorResponse = require('../utils/errorResponse');
const asyncHandler = require('../middleware/async');
const Store = require('../models/Store');
const UserService = require("../services/userServices");
const StoreService = require("../services/storeServices");


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
    console.log(req.user);
    if(req.user && req.user.role === 'store'){
        reqQuery.user = req.user.id;
    }
    const stores = await Store.find(reqQuery).populate('user');
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
    req.body.user = req.user.id;
    const store = await StoreService.create(req.body);
    
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

    store = await StoreService.update(storeId, req.body);

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