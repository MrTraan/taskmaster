const process = require("process");
const readline = require('readline');
const tmparser = require("./parsing");
const tmproc = require("./tmproc");
const rl = readline.createInterface(process.stdin, process.stdout, completer);

const regIsKeyWord = new RegExp("^(status|start|exit|stop|restart|reload)$");
const compKeyWords = ["status", "start", "exit", "stop", "restart", "reload"];

function completer(line) {
    var hits = [];
    
    var words = line.trim().replace(/\s+/g, " ").split(" ");
    if (words.length == 0) return ([], line);
    if (words.length == 1 && regIsKeyWord.test(words[0])) {
        var m = "*";
        getJobByName(m, (match) => {
            match.forEach((j) => {
                hits.push(j.name);  
            });
        });
    } else if (words.length == 1){
        compKeyWords.forEach((w) => {
            if (w.indexOf(words[0]) == 0) hits.push(w);
        });
    } else {
        var m = words[words.length - 1] + "*";
        getJobByName(m, (match) => {
            match.forEach((j) => {
                var matchArr = words;
                matchArr.splice(-1, 1, j.name);
                hits.push(matchArr.join(" "));  
            });
        });
    }
    return [hits , line];
}

var holder = [];

tmparser.readConfigFile("conf.json", (config) => {
    config.forEach((entry) => {
        holder.push(new tmproc.Process(entry));
    });
    rl.prompt();
});

//rl.setPrompt('tm> ');

rl.on("line", (line) => {
    var av = line.trim().replace(/\s+/g, " ").split(" ");
    switch(av[0]) {
        case "exit":
            console.log("Have a great day!");
            process.exit(0);
            break;
        case "status":
            if (av.length == 1) {
                av[1] = "*";
            }
            for (var i = 1; i < av.length; i++) {
                getJobByName(av[i], (jobs) => {
                    if (jobs.length == 0) {
                        console.log("there is no task called " + av[1]);
                    } else {
                        jobs.forEach((job) => {
                            console.log(job.status());
                        });
                    }
                });
            }
            break;
        case "stop":
            if (av.length == 1) {
                av[1] = "*";
            }
            for (var i = 1; i < av.length; i++) {
                getJobByName(av[i], (jobs) => {
                    if (jobs.length == 0) {
                        console.log("there is no task called " + av[i]);
                    } else {
                        jobs.forEach((job) => {
                            job.stop();
                        });
                    }
                });
            }
            break;
        case "start":
            if (av.length == 1) {
                av[1] = "*";
            }
            for (var i = 1; i < av.length; i++) {
                getJobByName(av[i], (jobs) => {
                    if (jobs.length == 0) {
                        console.log("there is no task called " + av[i]);
                    } else {
                        jobs.forEach((job) => {
                            job.start();
                        });
                    }
                });
            }
            break;
        case "restart":
            if (av.length == 1) {
                av[1] = "*";
            }
            for (var i = 1; i < av.length; i++) {
                getJobByName(av[i], (jobs) => {
                    if (jobs.length == 0) {
                        console.log("there is no task called " + av[i]);
                    } else {
                        jobs.forEach((job) => {
                            job.restart();
                        });
                    }
                });
            }
            break;
        default:
            console.log('Say what? I might have heard `' + line.trim() + '`');
            break;
    }
    rl.prompt();
}).on('close', () => {
    console.log('Have a great day!');
    process.exit(0);
});

function getJobByName(str, callback){
    var match = [];
    
    if (str.charAt(str.length - 1) == '*') {
        try {
            var reg = new RegExp("^" + str.substr(0, str.length - 1));
        } catch(err) {
            callback([]);
            return ;
        }
        holder.forEach((job) => {
            if (reg.test(job.name)) match.push(job);
        });
    } else {
        holder.forEach((job) => {
           if (job.name == str) match.push(job);
        });
    }
    callback(match);
}