const mongoose = require('mongoose');
mongoose.set('strictQuery', true);

const connectDB = async () => {
    const conn = await mongoose.connect(process.env.MONGO_URI, {
        useNewUrlParser: true,
        // useFindAndModify: false,
        useUnifiedTopology: true,
        // useCreateIndex: true
    });

    console.log(`connected mongoose  ${conn.connection.host}`);
}

module.exports = connectDB;