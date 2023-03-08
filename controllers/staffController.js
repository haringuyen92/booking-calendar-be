const asyncHandler = require('../middleware/async');
const errorResponse = require("../utils/errorResponse");
const StaffService = require("../services/staffServices");
const StoreService = require("../services/storeServices");

exports.getStaffs = asyncHandler( async(req, res, next) => {
    const storeId = req.params.storeId;
    const store = await StoreService.getOne(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    const staffs = await StaffService.getBy({
        store: storeId
    });
    return res.status(200).json({
        success: true, 
        message: "success getStaffs",
        count: staffs.length,
        data: staffs,
    }); 
})

exports.getStaff = asyncHandler( async(req, res, next) => {
    const { storeId, id } = req.params;
    const store = await StoreService.getOne(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    const staff = await StaffService.getOne(id);
    return res.status(200).json({
        success: true, 
        message: "success getStaff",
        data: staff,
    }); 
})

exports.createStaff = asyncHandler( async(req, res, next) => {
    const storeId = req.params.storeId;
    const store = await StoreService.getOne(storeId);
    
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    req.body.store = storeId;
    const staff = await StaffService.create(req.body);
    return res.status(200).json({
        success: true,
        message: "success createStaff",
        data: staff
    })
})

exports.updateStaff = asyncHandler( async(req, res, next) => {
    const { storeId, id } = req.params;
    const store = await StoreService.getOne(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));
    
    const currentStaff = await StaffService.getOne(id);
    if(!currentStaff) return next(new errorResponse(`Staff not found id: ${id}`, 404));
    const { isAllCourse } = req.body;
    if(isAllCourse) req.body.courses = [];

    const staff = await StaffService.update(id, req.body);
    return res.status(200).json({
        success: true,
        message: "success updateStaff",
        data: staff
    })
})

exports.deleteStaff = asyncHandler( async(req, res, next) => {
    const { storeId, id } = req.params;
    const store = await StoreService.getOne(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    const staff = await StaffService.delete(id);
    return res.status(200).json({
        success: true,
        message: "success deleteStaff",
        data: staff
    })
})