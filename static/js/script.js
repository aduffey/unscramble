const VALID_CHARS = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'

function handleKey(input, event, nextInputName) {
    event.preventDefault()  // We'll handle the input ourselves
    var key = event.keyCode || event.which;
    var keyChar = event.key || String.fromCharCode(key);
    keyChar = keyChar.toUpperCase()
    if (VALID_CHARS.lastIndexOf(keyChar) >= 0) {
        // Special case: 'Q' really means 'Qu'. There is no 'Q' cell.
        if (keyChar === 'Q') {
            input.value = 'Qu'
        } else {
            input.value = keyChar
        }
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
        // Handle a 'Qu' cell. This is represented simply as 'q' in a board
        // string.
        if (c === 'qu') {
            c = 'q'
        }
        boardString += c
    }
    window.location.href = '/' + boardString
}
