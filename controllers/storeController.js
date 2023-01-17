const errorResponse = require('../utils/errorResponse');
const asyncHandler = require('../middleware/async');
const Users = require('../models/Users');
const Stores = require('../models/Stores');


// @desc   getAll
// @route  GET /api/stores
// @route  GET /api/users/:userId/stores
// @access Public

exports.getStores = asyncHandler( async(req, res, next) => {
    const reqQuery = {...req.query};
    
    const removeFields = ['select', 'sort'];
    removeFields.forEach(q => delete reqQuery[q]);

    let queryStr = JSON.stringify(reqQuery);


    //create operators ($gt,...)
    // queryStr = queryStr.replace(/\b(gt)\b/, m => `$${m}`);
    if(req.params.userId){
        reqQuery.user = req.params.userId;
    }
    let query = Stores.find(reqQuery).populate('user');


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
    const store = await Stores.findById(storeId);

    res.status(200).json({ 
        success: true, 
        message: `success getStore`,
        store: store
    }); 
})

// @desc   create
// @route  POST /api/stores
// @access Public

exports.createStore = asyncHandler( async(req, res, next) => {
    const store = await Stores.create(req.body);
    
    res.status(200).json({ 
        success: true, 
        message: `success createStore`,
        store: store
    }); 
})

// @desc   update 
// @route  PUT /api/stores/:id
// @access Public

exports.updateStore = asyncHandler( async(req, res, next) => {
    const storeId = req.params.id;

    const store = await Stores.findByIdAndUpdate(storeId, req.body, {
        new: true,
        runValidators: true
    });
    if(!store) return next(new errorResponse(`cannot find Store with ID: ${storeId}`, 404));

    res.status(200).json({ 
        success: true, 
        message: `success updateStore ${storeId}`,
        store: store
    }); 
})

// @desc   delete 
// @route  DELETE /api/stores/:id
// @access Public

exports.deleteStore = asyncHandler( async(req, res, next) => {
    const storeId = req.params.id;
    const store = await Stores.findByIdAndDelete(storeId);
    if(!store) return next(new errorResponse(`cannot find Store with ID: ${storeId}`, 404));

    res.status(200).json({ 
        success: true, 
        message: `success deleteStore ${storeId}`,
        store: store
    }); 
})