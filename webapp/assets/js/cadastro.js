$('#register-form').on('submit', createUser);

function createUser(event){
    event.preventDefault();

    if ($('#password').val() != $('#confirm-password').val()){
        alert("Diferent Passwords!");
        return;
    }
    
    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            password: $('#password').val(),
        }
    }).done(function (){
        alert("user successfully signed up")
    }).fail(function(err){
        console.log(err)
        alert("error while signing up the user")
    })
}