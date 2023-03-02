const Store = require("../models/Store");

class StoreServices{
    get(id){
        return Store.findById(id);
    }
}
module.exports = new StoreServices();