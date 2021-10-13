$('#login').on('submit', signIn);

function signIn(event) {
    event.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            password: $('#password').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function(err) {
        console.log(err)
        Swal.fire("Ops...", "The email or password are wrong", "error")
    });
}