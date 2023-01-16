const errorResponse = require("../utils/errorResponse");

const errorHandler = (err, req, res, next) => {
    let error = {...err};
    let message = 'Server Error';
    error.message = err.message;

    console.log(err.stack.red);

    if(err.name === 'CastError'){
        message = `Resource not found with id of ${err.value}`;
        error = new errorResponse(message, 404);
    }
    if(err.code === 11000){
        message = `Duplicate pattern`;
        error = new errorResponse(message, 400);
    }
    if(err.name === 'ValidationError'){
        message = Object.values(err.errors).map(val => val.message);
        error = new errorResponse(message, 400);
    }

    res.status(error.statusCode || 500).json({
        success: false,
        error: error.message || message
    });
}

module.exports = errorHandler;