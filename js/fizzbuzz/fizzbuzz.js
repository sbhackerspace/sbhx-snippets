// @ajvbahnken + @elimisteve

for (var i = 1; i < 101; i++) {
    var line = "";
    if (i % 3 === 0) line += "Fizz";
    if (i % 5 === 0) line += "Buzz";
    console.log(line || i);
}
