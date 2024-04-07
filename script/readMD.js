const fs = require('fs');

let file = '4.debruijn.md'
fs.readFile(file
, 'utf8', (err, file) => {
    if (err) throw err;
    console.log(file);
});