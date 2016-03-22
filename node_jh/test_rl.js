const readline = require('readline')
const colors = require('colors')

const rl = readline.createInterface({
	input:		process.stdin,
	output:		process.stdout,
	completer:	completer
})

var completionsCmd = ['status', 'load', 'reload', 'stop', 'shutdown', 'exit']
var completionsArg = ['lel', 'lal']


function completer(line) {
	var tmp = line.split(' ')

	if (tmp.length == 1) {
		var hits = completionsCmd.filter((c) => { return c.indexOf(line) == 0 })
		return [hits.length ? hits : completionsCmd, line]
	} else {
		var hits = completionsArg.filter((c) => { return c.indexOf(tmp[tmp.length - 1]) == 0 })
		return [hits.length ? hits : completionsArg, line]
	}
}

function addCompletionKey(tab, key) {
	tab.push(key)
}

function displayCompletionTab(tab) {
	tab.forEach((key) => {
		console.log('key is: ' + key)
	})
}

addCompletionKey(completionsArg, 'lol')
addCompletionKey(completionsArg, 'test')
displayCompletionTab(completionsArg)

rl.setPrompt('lol> ')
rl.prompt()

rl.on('line', (line) => {
	switch (line.trim()) {
		case 'exit':
			rl.close()
		default:
			console.log('Say what ? I might have heard ' + line.rainbow)
	}
	rl.prompt()
}).on('close', () => {
	console.log('this is over'.rainbow)
	process.exit(1)
})
