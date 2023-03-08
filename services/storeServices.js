const Store = require("../models/Store");

class StoreService{
    getOne(id){
        return Store.findById(id);
    }
    create(data){
        return Store.create(data);
    }
    update(id, data, options = {}){
        if(Object.keys(options).length === 0){
            options = {
                new: true,
                runValidators: true
            };
        }
        return Store.findByIdAndUpdate(id, data, {
            new: true,
            runValidators: true
        })
    }
}
module.exports = new StoreService();