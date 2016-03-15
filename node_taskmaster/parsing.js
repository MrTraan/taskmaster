const fs = require("fs");
const process = require("process");
const path = require("path");
const tmlogs = require("./tmlogs");

function readConfigFile(filename, callback){
    fs.readFile("conf.json", (err, data) => {
        if (err) throw err;
        var json = JSON.parse(data);
        var conf = [];
        json.forEach((entry) => { 
            if (typeof(entry.numprocs) == "number" && entry.numprocs > 1){
                for (var i = 0; i < entry.numprocs; i++) {
                    var copy = new Conf(entry);
                    copy.name = copy.name + ":" + parseInt(i);
                    conf.push(copy);
                }
            }
            else {
                conf.push(new Conf(entry));   
            }
        });
        callback(conf);
    });
}
module.exports.readConfigFile = readConfigFile;

function parseCmd(cmd){
        av = cmd.split(" ");
        if (av[0].charAt(0) != "/" && av[0].charAt(0) != ".") {
            av[0] = findPath(av[0]);
        }
        return (av);
}

function findPath(cmd) {
    var pathF = process.env.PATH.split(":");
    for (var i = 0; i < pathF.length; i++) {
        if (fs.readdirSync(pathF[i]).indexOf(cmd) > -1)
            return (path.join(pathF[i], cmd));
    }
    return (cmd);
}

function Conf(json) {
    this.name = (json.name != undefined ? json.name : "");
    this.cmd = (json.cmd != undefined ? parseCmd(json.cmd) : ["/bin/echo", json.name, "need a valid cmd"]);
    this.numprocs = (json.numprocs != undefined ? json.numprocs : 1);
    this.umask = (json.umask != undefined ? parseInt(json.umask, 8) : 0o22);
    this.autostart = (json.autostart != undefined ? json.autostart : false);
    this.autorestart = (json.autorestart != undefined ? json.autorestart : "never");
    this.exitcodes = (json.exitcodes != undefined ? json.exitcodes : [0]);
    this.startretries = (json.startretries != undefined ? json.startretries : 0);
    this.starttime = (json.starttime != undefined ? json.starttime : 0);
    this.stopsignal = (json.stopsignal != undefined ? json.stopsignal : "SIGHUP");
    this.stoptime = (json.stoptime != undefined ? json.stoptime : 0);
    this.stdout = (json.stdout != undefined ? json.stdout : "/dev/null");
    this.stderr = (json.stderr != undefined ? json.stderr : "/dev/null");
    this.env = (json.env != undefined ? json.env : process.env);
}