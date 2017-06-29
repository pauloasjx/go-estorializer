
var socket;
var words = [];

function Word(word, count) {
	this.word = word;
	this.count = count*20;
	this.pos = createVector(random(windowWidth), random(windowHeight));
	this.vel = createVector(0);
	this.col = color(random(255), random(255), random(255));

  	this.intersect = function(other){
    	if(this.pos.sub(other.pos) <= this.count) {
    		this.vel.mult(-1);
    	}
  	}

	this.draw = function() {
		fill(this.col);
		ellipse(this.pos.x, this.pos.y, this.count, this.count);
		textSize(this.count/2);
		fill(255);
		text(this.word, this.pos.x, this.pos.y);
	}

	this.update = function() {
		this.pos.add(this.vel);
	}

}

function setup() {
	createCanvas(windowWidth, windowHeight)
	background(0);
	noStroke();

	socket = io.connect();

	socket.on('new-word', function(word){
		words.push(new Word(word.Word, word.Count))
	});

	socket.on('count-word', function(word){
		for(var i=0; i<words.length; i++) {
			if(words[i].Word == word) {
				word[i].count++;
			}
		}
	});

}

function draw() {
	background(0);

	for(var i=0; i<words.length; i++) {
		words[i].update();
		words[i].draw();
		for(var j=0; j<words.length; j++) {
			if(words[i] != words[j]) {
				words[i].intersect(words[j]);
			}
		}
	}
}