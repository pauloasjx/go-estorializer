
var words = [];

function Word(word, count) {
	this.word = word;
	this.count = count;
	this.pos = createVector(random(windowWidth), random(windowHeight));
	this.vel = createVector(random(-1, 1), random(-1, 1));
	this.col = color(random(255), random(255), random(255));

  	this.intersect = function(other){
  		var aux = this.pos.copy();
    	if(aux.sub(other.pos) >= this.count*5+other.count*5) {
    		this.vel.mult(-1.1);
    		other.vel.mult(-1.1);
    	}
  	}

  	this.check = function() {
  		if(this.pos.x < 0 || this.pos.y < 0 || this.pos.x > windowWidth || this.pos.y > windowHeight) {
  			return true;
  		}
  		return false;
  	}

	this.draw = function() {
		fill(this.col);
		ellipse(this.pos.x, this.pos.y, this.count, this.count);
		textSize(this.count/3-this.word.length);
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
/*
	socket = io.connect();

	socket.on('new-word', function(word){
		words.push(new Word(word.Word, word.Count))
	});

	socket.on('count-word', function(word){
		for(var i=0; i<words.length; i++) {
			if(words[i].word == word.Word) {
				words[i].count++;
			}
		}
	});
*/
}

function draw() {
	background(0);

	for(var i=words.length-1; i>=0; i--) {
		words[i].update();
		words[i].draw();

		if(words[i].check()) {
			words.splice(i, 1);
		}

		for(var j=0; j<words.length; j++) {
			if(words[i] != words[j]) {
				words[i].intersect(words[j]);
			}
		}
	}
}