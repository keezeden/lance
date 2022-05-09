const message = "hello world"

out(message)

const age = 21

const greet = (name) => {
    log("Hello " + name)
}

const name = in("Name: ")

if (name == "lance") {
    out("Thats me, not you")
} else {
    greet(name)
}

const data = [1,2,3,4,5]
const iterations = 5
let count = 0

while (count < iterations) {
    out("We are at item: " + data[count])
}





