function setStorage(token) {
    localStorage.setItem("token", token)
}

function queryParams() {
    var token = localStorage.getItem("token")
    return "?token=" + token
}

// load page
$(document).ready(function () {
    csrf = $("meta[name=_csrf]").attr("content")

    // set cursor to the end of autofocus input string
    $("input[autofocus]").each(function () {
        $(this).val($(this).val())
    })


})