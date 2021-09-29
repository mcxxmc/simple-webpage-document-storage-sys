const urlDefaultView = "http://localhost:8080/default-view"

$(function(){

    $.ajax({
        type: "GET",
        url: urlDefaultView,
        success: function(data) {
            alert(data);
        },
        error: function(data) {
            alert("error");
        }
    });


})