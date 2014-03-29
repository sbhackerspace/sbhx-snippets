// @ajvbahnken + @elimisteve

// Printing to the console
for (var i = 1; i =< 100; i++) {
    var line = "";
    if (i % 3 === 0) line += "Fizz";
    if (i % 5 === 0) line += "Buzz";
    console.log(line || i);
}

// Printing to the browser
var d = document.getElementsByTagName('div')[0];
for (var i = 1; i =< 100; i++) {
    var line = "";
    if (i % 3 === 0) line += "Fizz";
    if (i % 5 === 0) line += "Buzz";
    d.innerHTML += (line || i) + '<br>';
}
