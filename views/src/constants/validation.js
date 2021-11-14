export function checkValidInput(inputTxt) {
    const letters = /^[0-9a-zA-Z_-]+$/;
    if (letters.test(inputTxt)) {
        return true
    } else {
        alert("invalid character(s)");
        return false
    }
}