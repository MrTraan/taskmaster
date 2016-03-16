var fs = require("fs");

const logFile = fs.createWriteStream("taskmaster_logs.txt", {flags: "a+"});

function logProcStart(name) {
    logFile.write(Date().toString() + " " + name + " has started\n");
}

module.exports.logProcStart = logProcStart;

function logProcStop(name, code) {
    logFile.write(`${Date().toString()} ${name} exited with code ${code}\n`);
}
module.exports.logProcStop = logProcStop;

function log(str) {
    logFile.write(`${Date().toString()} ${str}`);
}
module.exports.log = log;