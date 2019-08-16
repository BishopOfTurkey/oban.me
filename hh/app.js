const Y_AXIS = 1;
const X_AXIS = 2;
let b1, b2, c1, c2;
let rainbow, direction;

let RAINBOW_SPEED = 15;

function setup() {
    createCanvas(windowWidth, windowHeight);
    // Define colors
    b1 = color(255);
    b2 = color(0);
    c1 = color(237, 102, 90);
    c2 = color(0, 102, 153);
    
    rainbow = 0;
    direction = 1;

    // noLoop();
    frameRate(25);
}

function draw() {
    c1 = color(rainbow, rainbow + 123 & 256, rainbow - 167 % 256);
    setGradient(0, 0, width, height, c1, c2, Y_AXIS);
    if (rainbow >= 256) {
        direction = -1;
    } else if (rainbow <= 0) {
        direction = 1;
    }
    rainbow += RAINBOW_SPEED * direction;

    mouse(mouseX, mouseY);
    itsHummusTime();
    menu();
}


function mouse(x, y) {
    fill(100);
    textSize(width / 20);
    text('Hummus', x + Math.random()*20 - 10, y + Math.random()*20 - 10);
}

function itsHummusTime() {
    let percent = (second()/60) * width;
    text("It's Hummus TIME", percent, 90);
}

function menu() {
    textSize(20);
    text("Although multiple different theories and claims of origins exist in various parts of the Middle East, evidence is\n insufficient to determine the precise location or time of the invention of hummus. Its basic ingredients—chickpeas, sesame,\n lemon, and garlic—have been combined and eaten in the Levant over centuries. Though regional populations widely ate\n chickpeas, and often cooked them in stews and other hot dishes, puréed chickpeas eaten cold with tahini do not appear before\n the Abbasid period in Egypt and the Levant.", width/5, height/4);
}



function setGradient(x, y, w, h, c1, c2, axis) {
    noFill();

    if (axis === Y_AXIS) {
        // Top to bottom gradient
        for (let i = y; i <= y + h; i++) {
            let inter = map(i, y, y + h, 0, 1);
            let c = lerpColor(c1, c2, inter);
            stroke(c);
            line(x, i, x + w, i);
        }
    } else if (axis === X_AXIS) {
        // Left to right gradient
        for (let i = x; i <= x + w; i++) {
            let inter = map(i, x, x + w, 0, 1);
            let c = lerpColor(c1, c2, inter);
            stroke(c);
            line(i, y, i, y + h);
        }
    }
}