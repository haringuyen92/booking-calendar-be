const Course = require("../models/Course");
class StaffServices{
    get(id){
        return Course.findById(id);
    }
    async create(){

    }
    async delete(id){
        return Course.findByIdAndDelete(id);
    }
}
module.exports = new StaffServices();