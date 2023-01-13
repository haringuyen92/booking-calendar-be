const logger = (req, res, next) => {
    const fullUrl = req.protocol + '://' + req.get('host') + req.originalUrl;
    console.log(`${req.method} ${fullUrl}`);
    res.setHeader('X-Powered-By', 'Hainv');
    next();
}

module.exports = logger;