const errorResponse = require('../utils/errorResponse');
const asyncHandler = require('../middleware/async');
const Store = require('../models/Store');

exports.getStores = asyncHandler( async(req, res, next) => {
    const reqQuery = {...req.query};
    const removeFields = ['select', 'sort'];
    removeFields.forEach(q => delete reqQuery[q]);

    let queryStr = JSON.stringify(reqQuery);


    //create operators ($gt,...)
    // queryStr = queryStr.replace(/\b(gt)\b/, m => `$${m}`);
    let query = Store.find(JSON.parse(queryStr));

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
    const total = await Store.countDocuments();
    //pagination
    const page = parseInt(req.query.page, 10) || 1;
    const limit = parseInt(req.query.limit, 10) || 10;
    const skip = (page - 1)*limit;

    query.skip(skip).limit(limit);

    const stores = await query;
    res.status(200).json({ 
        success: true, 
        message: "success getStores",
        total: total,
        count: stores.length,
        users: stores
    }); 
})

exports.getStore = asyncHandler( async(req, res, next) => {
    const storeId = req.params.id;
    const store = await Store.findById(storeId);

    res.status(200).json({ 
        success: true, 
        message: `success getStore`,
        store: store
    }); 
})

exports.createStore = asyncHandler( async(req, res, next) => {
    const store = await Store.create(req.body);
    
    res.status(200).json({ 
        success: true, 
        message: `success createStore`,
        store: store
    }); 
})

exports.updateStore = asyncHandler( async(req, res, next) => {
    const storeId = req.params.id;

    const store = await Store.findByIdAndUpdate(storeId, req.body, {
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

exports.deleteStore = asyncHandler( async(req, res, next) => {
    const storeId = req.params.id;
    const store = await Store.findByIdAndDelete(storeId);
    if(!store) return next(new errorResponse(`cannot find Store with ID: ${storeId}`, 404));

    res.status(200).json({ 
        success: true, 
        message: `success deleteStore ${storeId}`,
        store: store
    }); 
})