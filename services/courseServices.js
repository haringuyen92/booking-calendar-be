const Course = require("../models/Course");
class CourseServices{
    getBy(condition = {}){
        return Course.find(condition).populate('store');
    }
    getOne(id){
        return Course.findById(id).populate([
            {
                path: 'store'
            },
            {
                path: 'courses',
                select: 'name'
            }
        ]);
    }
    create(data){
        return Course.create(data);
    }
    update(id, data, options = {}){
        if(Object.keys(options).length === 0){
            options = {
                new: true,
                runValidators: true
            }
        }
        return Course.findByIdAndUpdate(id, data,  options);
    }
    async delete(id){
        return Course.findByIdAndDelete(id);
    }
}
module.exports = new CourseServices();