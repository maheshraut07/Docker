const express = require('express');
const cors = require('cors');
const app = express();

app.use(cors());

app.get('/', (req, res) => {
    res.send({
        name: "mahesh",
        email: "rautmahesh213@gmail.com"
    });
});

app.listen(5200, () => {
    console.log('Example app listening on port 5200');
});
