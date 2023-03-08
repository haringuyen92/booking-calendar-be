const Staff = require("../models/Staff");
class StaffServices{
    getBy(condition = {}){
        return Staff.find(condition).populate('store');
    }
    getOne(id){
        return Staff.findById(id).populate([
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
        return Staff.create(data);
    }
    update(id, data, options = {}){
        if(Object.keys(options).length === 0){
            options = {
                new: true,
                runValidators: true
            }
        }
        return Staff.findByIdAndUpdate(id, data,  options);
    }
    async delete(id){
        return Staff.findByIdAndDelete(id);
    }
}
module.exports = new StaffServices();