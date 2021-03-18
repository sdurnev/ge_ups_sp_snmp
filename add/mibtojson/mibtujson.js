let asn2json = require('asn2json');
let fs = require('fs');

// Create a new instance of asn2json
let asn = new asn2json();

let data = fs.readFileSync(`GEPARALLELUPSv024.mib`).toString();

let json = asn.parse(data);

//fs.writeFileSync('mib.json', json);
console.log(json);
