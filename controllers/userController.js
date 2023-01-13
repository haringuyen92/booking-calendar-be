exports.getUsers = (req, res, next) => {
    res.status(200).json({ success: true, message: "success getUsers"});
}

exports.getUser = (req, res, next) => {
    const userId = req.params.id;
    res.status(200).json({ success: true, message: `success getUser ${userId}`});
}

exports.createUser = (req, res, next) => {
    res.status(200).json({ success: true, message: "success createUser createUser"});
}

exports.updateUser = (req, res, next) => {
    const userId = req.params.id;
    res.status(200).json({ success: true, message: `success updateUser ${userId}`});
}

exports.deleteUser = (req, res, next) => {
    const userId = req.params.id;
    res.status(200).json({ success: true, message: `success deleteUser ${userId}`});
}