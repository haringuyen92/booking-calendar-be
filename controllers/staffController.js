const asyncHandler = require('../middleware/async');
const Staff = require('../models/Staff');
const Store = require('../models/Store');
const errorResponse = require("../utils/errorResponse");
const StaffService = require("../services/StaffServices");
const StoreService = require("../services/StoreServices");

exports.getStaffs = asyncHandler( async(req, res, next) => {
    const storeId = req.params.storeId;
    const store = await StoreService.get(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    const staffs = await Staff.find({
        store: storeId
    }).populate('store');
    return res.status(200).json({
        success: true, 
        message: "success getStaffs",
        count: staffs.length,
        data: staffs,
    }); 
})

exports.getStaff = asyncHandler( async(req, res, next) => {
    const { storeId, id } = req.params;
    const store = await StoreService.get(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    const staff = await Staff.findById(id).populate('store');
    return res.status(200).json({
        success: true, 
        message: "success getStaff",
        data: staff,
    }); 
})

exports.createStaff = asyncHandler( async(req, res, next) => {
    const storeId = req.params.storeId;
    const store = await StoreService.get(storeId);
    
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    req.body.store = storeId;
    const staff = await Staff.create(req.body);
    return res.status(200).json({
        success: true,
        message: "success createStaff",
        data: staff
    })
})

exports.updateStaff = asyncHandler( async(req, res, next) => {
    const { storeId, id } = req.params;
    const store = await StoreService.get(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));
    
    let staff = await StaffService.get(id);
    if(!staff) return next(new errorResponse(`Staff not found id: ${id}`, 404));

    staff = await Store.findByIdAndUpdate(storeId, req.body, {
        new: true,
        runValidators: true
    });
    return res.status(200).json({
        success: true,
        message: "success updateStaff",
        data: staff
    })
})

exports.deleteStaff = asyncHandler( async(req, res, next) => {
    const { storeId, id } = req.params;
    const store = await StoreService.get(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    const staff = await StaffService.delete(id);
    return res.status(200).json({
        success: true,
        message: "success deleteStaff",
        data: staff
    })
})