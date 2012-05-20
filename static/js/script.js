const VALID_CHARS = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'

function handleKey(input, event, nextInputName) {
    event.preventDefault()  // We'll handle the input ourselves
    var key = event.keyCode || event.which;
    var keyChar = event.key || String.fromCharCode(key);
    if (VALID_CHARS.lastIndexOf(keyChar.toUpperCase()) >= 0) {
        input.value = keyChar.toUpperCase()
        if (nextInputName != '') {
            document.getElementsByName(nextInputName)[0].focus()
        }
    }
}

function handleSubmit() {
    var form = document.getElementById('board')
    var boardString = ''
    for (var i = 0; i < 16; i++) {  // Only the first 16 elements
        c = form.elements[i].value.toLowerCase()
        boardString += c
    }
    window.location.href = '/' + boardString
}
