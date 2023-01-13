const express = require('express');

const router = express.Router();

router.get('/users', (req, res) => {
    res.send("get all User");
});
router.get('/users/:id', (req, res) => {
    res.send("get detail User");
});
router.post('/users', (req, res) => {
    res.send("create User");
});
router.patch('/users/:id', (req, res) => {
    res.send("update User");
});
router.delete('/users/:id', (req, res) => {
    res.send("delete User");
});

module.exports = router;