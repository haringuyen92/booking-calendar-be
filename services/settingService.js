const SettingSlot = require("../models/SettingSlot")

class SettingService {
    getSettingSlot(storeId) {
        return SettingSlot.findOne({
            store: storeId
        });
    }

    postSettingSlot(storeId, data, options = {}) {
        if (Object.keys(options).length === 0) {
            options = {
                upsert: true,
                new: true,
                setDefaultsOnInsert: true
            }
        }
        return SettingSlot.findOneAndUpdate({
            store: storeId
        }, data, options);
    }
}

module.exports = new SettingService();