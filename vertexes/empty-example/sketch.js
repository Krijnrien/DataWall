var circles = [];
let url = 'http://145.93.118.9:3000/devices';

let time;
let wait = 20000;
let tick = false;

function preload() {
    fetchJson(fetchToCircles);
}
function setup() {
    createCanvas(window.innerWidth, window.innerHeight);
    time = millis();

    // loadJSON('dummy-data.json', fetchToCircles);

    setBackground();
    /*
    for (var i = 0; i < 1050; i++) {
        var circle = {
            x: random(width),
            y: random(height),
            r: 4,
        };
        circles.push(circle);
    }
   */
    for (var j = 0; j < circles.length; j++) {
        for (var k = 0; k < circles.length; k++) {
            if (
                circles[j].x - circles[k].x < 40 &&
        circles[j].x - circles[k].x > -40 &&
        circles[j].y - circles[k].y < 40 &&
        circles[j].y - circles[k].y > -40
            ) {
                r = random(255);
                g = random(255);
                b = random(255);
                stroke(r, g, b);
                line(circles[k].x, circles[k].y, circles[j].x, circles[j].y);

                //break;
            }
        }
    }
}
function draw() {
    if (millis() - time >= wait) {
        tick = !tick;
        time = millis();
        preload();
        clear();
        setBackground();
        for (var j = 0; j < circles.length; j++) {
            for (var k = 0; k < circles.length; k++) {
                if (
                    circles[j].x - circles[k].x < 70 &&
          circles[j].x - circles[k].x > -70 &&
          circles[j].y - circles[k].y < 70 &&
          circles[j].y - circles[k].y > -70
                ) {
                    r = random(255);
                    g = random(255);
                    b = random(255);
                    stroke(r, g, b);
                    line(circles[k].x, circles[k].y, circles[j].x, circles[j].y);

                    //break;
                }
            }
        }
        for (var v = 0; v < circles.length; v++) {
            r = random(255);
            g = random(255);
            b = random(255);

            fill(r, g, b);
            ellipse(circles[v].x, circles[v].y, circles[v].r * 2, circles[v].r * 2);

            //line(circles[i].x,circles[i].y,circles[i].x,circles[i].y);
        }

        circles = [];
    }

    /*
    for (var i = 0; i < circles.length; i++) {
      circles[i].move();
      circles[i].display();      
    } */
    /*
    for (var j = 0; j < circles.length; j++) {
        for (var k = 0; k < circles.length; k++) {
            if (
                circles[j].x - circles[k].x < 40 &&
        circles[j].x - circles[k].x > -40 &&
        circles[j].y - circles[k].y < 40 &&
        circles[j].y - circles[k].y > -40
            ) {
                r = random(255);
                g = random(255);
                b = random(255);
                stroke(r, g, b);
                line(circles[k].x, circles[k].y, circles[j].x, circles[j].y);

                //break;
            }
        }
    }
    */
}
function fetchJson(func) {
    fetch(url)
        .then(res => {
            return res.json();
        })
        .then(json => {
            console.log(json);
            func(json);
            return json;
        })
        .catch(err => {
            console.log(err);
        });
}
function fetchToCircles(json) {
    for (let i = Object.keys(json).length - 1; i >= 0; i--) {
        if (json[i].z == 0) {
            let currentCircle = json[i];
            let temp = new circle(currentCircle.x * 2.3, currentCircle.y * 2.3, 4);
            circles.push(temp);
        }
    }
}

function setBackground() {
    background(30, 30, 30);
}
class circle {
    constructor(x, y, r) {
        this.x = x;
        this.y = y;
        this.r = r;
    }
}
