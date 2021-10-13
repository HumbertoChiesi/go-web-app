$('#register-form').on('submit', createUser);

function createUser(event){
    event.preventDefault();

    if ($('#password').val() != $('#confirm-password').val()){
        Swal.fire("Ops...", "The Passwords are different!", "error")
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
        Swal.fire("Success!", "user successfully signed up", "success").then(function(){
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: $('#email').val(),
                    password: $('#password').val(),
                }
            }).done(function(){
                window.location = "/home";
            }).fail(function(){
                Swal.fire("Ops...", "ERROR, could not authenticate the user!", "error");
            })
        })

    }).fail(function(err){
        console.log(err)
        Swal.fire("Ops...", "ERROR while signing up the user!", "error")
    })
}