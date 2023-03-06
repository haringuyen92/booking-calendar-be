const express = require('express');
const courseRouter = express.Router({ mergeParams: true });

const { protect, authorize } = require('../middleware/auth');
const {getCourses, createCourse, getCourse, updateCourse, deleteCourse} = require("../controllers/courseController");

courseRouter
    .route('/')
    .get(protect, getCourses)
    .post(protect, authorize('store', 'admin'), createCourse);

courseRouter
    .route('/:id')
    .get(getCourse)
    .put(protect, authorize('store', 'admin'), updateCourse)
    .delete(protect, authorize('store', 'admin'), deleteCourse);

module.exports = courseRouter;