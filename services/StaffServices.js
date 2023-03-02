const Staff = require("../models/Staff");
class StaffServices{
    get(id){
        return Staff.findById(id);
    }
    async create(){

    }
    async delete(id){
        return Staff.findByIdAndDelete(id);
    }
}
module.exports = new StaffServices();