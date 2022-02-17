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


