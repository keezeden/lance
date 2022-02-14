const message = "hello world"

out(message)

const age = 21

error(age)

// This is a comment
const greet = (name) => {
    log("Hello " + name)
}

const name = in("Name: ")

if (name == "lance") {
    error("Thats me, not you")
} else {
    greet(name)
}


