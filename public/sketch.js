
var words = [];
var btn, url;
var mult;
const limit = 100;

function Word(word, count) {
	this.word = word;
	this.count = count*mult;
	this.pos = createVector(random(windowWidth), random(windowHeight));
	this.vel = createVector(random(-0.5, 0.5), random(-0.5, 0.5));
	this.col = color(random(255), random(255), random(255));

  	this.intersect = function(other){
    	if(dist(this.pos.x, this.pos.y, other.pos.x, other.pos.y) < this.count/2+other.count/2) {
    		this.vel.mult(-1);
    		other.vel.mult(-1);
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
	url = createInput('http://google.com.br');
	zoom = createInput('5');
	btn = createButton('Enviar');
	
	url.position(0, windowHeight - url.height);
	zoom.position(url.x + url.width, windowHeight - zoom.height);
	btn.position(zoom.x + zoom.width, windowHeight - btn.height);
	
	btn.mousePressed(request);

	textAlign(CENTER, CENTER);
	background(0);
	noStroke();

}

function request() {
	httpPost('/estorializer', url.value(), 'text', setWords);
	mult = zoom.value();
}

function setWords(data) {
	words = [];
	data = JSON.parse(data);
	
	for(var i=0; i < 100; i++) {
		words.push(new Word(data[i].Word, data[i].Count));
	}
}

function draw() {
	background(0);

	for(var i=words.length-1; i>=0; i--) {
		words[i].update();
		words[i].draw();
		/*
		for(var j=0; j<words.length; j++) {
			if(words[i] != words[j]) {
				words[i].intersect(words[j]);
			}
		}
		*/
	}
}