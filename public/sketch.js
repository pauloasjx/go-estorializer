
var socket;

function setup() {
	createCanvas(500, 500)
	background(0);

	socket = io.connect();

	socket.on('words', function(words){
		console.log(words)
	});
}

function draw() {

}