function setup() {
  createCanvas(windowWidth, windowHeight);
  background(220);
}

function draw() {
  fill(mouseX, mouseY, mouseX + mouseY, 20);
  rect(mouseX, mouseY, abs(pmouseX - mouseX) * 2, abs(pmouseX - mouseX) * 2);
}
