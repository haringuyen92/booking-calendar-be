const asyncHandler = require('../middleware/async');
const Course = require('../models/Course');
const errorResponse = require("../utils/errorResponse");
const CourseService = require("../services/CourseServices");
const StoreService = require("../services/StoreServices");

exports.getCourses = asyncHandler( async(req, res, next) => {
    const storeId = req.params.storeId;
    const store = await StoreService.get(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    const courses = await Course.find({
        store: storeId
    }).populate('store');
    return res.status(200).json({
        success: true,
        message: "success getCourses",
        count: courses.length,
        data: courses,
    });
})

exports.getCourse = asyncHandler( async(req, res, next) => {
    const { storeId, id } = req.params;
    const store = await StoreService.get(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    const course = await Course.findById(id).populate('store');
    return res.status(200).json({
        success: true,
        message: "success getCourse",
        data: course,
    });
})

exports.createCourse = asyncHandler( async(req, res, next) => {
    const storeId = req.params.storeId;
    const store = await StoreService.get(storeId);

    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    req.body.store = storeId;
    const course = await Course.create(req.body);
    return res.status(200).json({
        success: true,
        message: "success createCourse",
        data: course
    })
})

exports.updateCourse = asyncHandler( async(req, res, next) => {
    const { storeId, id } = req.params;
    const store = await StoreService.get(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    let course = await CourseService.get(id);
    if(!Course) return next(new errorResponse(`Course not found id: ${id}`, 404));

    course = await Course.findByIdAndUpdate(id, req.body, {
        new: true,
        runValidators: true
    });
    return res.status(200).json({
        success: true,
        message: "success updateCourse",
        data: course
    })
})

exports.deleteCourse = asyncHandler( async(req, res, next) => {
    const { storeId, id } = req.params;
    const store = await StoreService.get(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    const course = await CourseService.delete(id);
    return res.status(200).json({
        success: true,
        message: "success deleteCourse",
        data: course
    })
})