const asyncHandler = require('../middleware/async');
const errorResponse = require("../utils/errorResponse");
const CourseService = require("../services/courseServices");
const StoreService = require("../services/storeServices");

exports.getCourses = asyncHandler( async(req, res, next) => {
    const storeId = req.params.storeId;
    const store = await StoreService.getOne(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    const courses = await CourseService.getBy({
        store: storeId
    });
    return res.status(200).json({
        success: true,
        message: "success getCourses",
        count: courses.length,
        data: courses,
    });
})

exports.getCourse = asyncHandler( async(req, res, next) => {
    const { storeId, id } = req.params;
    const store = await StoreService.getOne(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    const course = await CourseService.getOne(id);
    return res.status(200).json({
        success: true,
        message: "success getCourse",
        data: course,
    });
})

exports.createCourse = asyncHandler( async(req, res, next) => {
    const storeId = req.params.storeId;
    const store = await StoreService.getOne(storeId);

    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    req.body.store = storeId;
    const course = await CourseService.create(req.body);
    return res.status(200).json({
        success: true,
        message: "success createCourse",
        data: course
    })
})

exports.updateCourse = asyncHandler( async(req, res, next) => {
    const { storeId, id } = req.params;
    const store = await StoreService.getOne(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    let course = await CourseService.getOne(id);
    if(!course) return next(new errorResponse(`Course not found id: ${id}`, 404));

    course = await CourseService.update(id, req.body);
    return res.status(200).json({
        success: true,
        message: "success updateCourse",
        data: course
    })
})

exports.deleteCourse = asyncHandler( async(req, res, next) => {
    const { storeId, id } = req.params;
    const store = await StoreService.getOne(storeId);
    if(!store) return next(new errorResponse(`Store not found id: ${storeId}`, 404));

    const course = await CourseService.delete(id);
    return res.status(200).json({
        success: true,
        message: "success deleteCourse",
        data: course
    })
})