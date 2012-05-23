const VALID_CHARS = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
const X2_LETTER_CLASS = 'x2Letter'
const X2_WORD_CLASS = 'x2Word'
const X3_LETTER_CLASS = 'x3Letter'
const X3_WORD_CLASS = 'x3Word'

function handleKey(input, event) {
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
		focusNextInput(input)
	} else if (keyChar === '2') {
		var cls = input.className
		if (cls === X2_LETTER_CLASS) {
			input.className = X2_WORD_CLASS
		} else if (cls === X2_WORD_CLASS) {
			input.className = ''
		} else {
			input.className = X2_LETTER_CLASS
		}
	} else if (keyChar === '3') {
		var cls = input.className
		if (cls === X3_LETTER_CLASS) {
			input.className = X3_WORD_CLASS
		} else if (cls === X3_WORD_CLASS) {
			input.className = ''
		} else {
			input.className = X3_LETTER_CLASS
		}
	}
}

function focusNextInput(currentInput) {
	var form = currentInput.form
	for (var i = 0; i < form.elements.length; i++) {
		var el = form.elements[i]
		if (el === currentInput && i < form.elements.length - 1) {
			form.elements[i + 1].focus()
		}
	}
}

function handleSubmit() {
	var form = document.getElementById('board')
	var boardString = ''
	for (var i = 0; i < 16; i++) {  // Only the first 16 elements
		var el = form.elements[i]

		// Get the modifier, if any
		var m = ''
		var cls = el.className
		if (cls === X2_LETTER_CLASS) {
			m = '2'
		} else if (cls === X2_WORD_CLASS) {
			m = '22'
		} else if (cls === X3_LETTER_CLASS) {
			m = '3'
		} else if (cls === X3_WORD_CLASS) {
			m = '33'
		}
		boardString += m

		// Get the character
		c = el.value.toLowerCase()
		// Handle a 'Qu' cell. This is represented simply as 'q' in a board
		// string.
		if (c === 'qu') {
			c = 'q'
		}
		boardString += c
	}
	window.location.href = '/' + boardString
}
