const User = require("../models/User");

class UserServices {
    getOne(id) {
        return User.findById(id);
    }
}

module.exports = new UserServices();