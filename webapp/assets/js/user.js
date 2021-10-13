$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);
$('#edit-user').on('submit', edit);
$('#update-password').on('submit', updatePassword);
$('#delete-user').on('click', deleteUser);

function unfollow(){
    const userId = $(this).data('user-id');
    $('#unfollow').prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/unfollow`,
        method: "POST"
    }).done(function(){
        window.location = `/users/${userId}`;
    }).fail(function(){
        Swal.fire("Ops...", "ERROR while unfollowing the user", "error");
        $('#unfollow').prop('disabled', false);
    });
}

function follow(){
    const userId = $(this).data('user-id');
    $('#follow').prop('disabled', true);

    $.ajax({
        url: `/users/${userId}/follow`,
        method: "POST"
    }).done(function(){
        window.location = `/users/${userId}`;
    }).fail(function(){
        Swal.fire("Ops...", "ERROR while following the user", "error");
        $('#follow').prop('disabled', false);
    });
}

function edit(evt){
    evt.preventDefault();

    $.ajax({
        url: "/edit-user",
        method: "PUT",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
        }
    }).done(function() {
        Swal.fire("Success!", "The user was updated!", "success")
            .then(function() {
                window.location = "/profile";
            });
    }).fail(function() {
        Swal.fire("Ops...", "ERROR while updating the user!", "error");
    });
}

function updatePassword(evt){
    evt.preventDefault();

    if( $('#new-password').val() != $('#confirm-password').val()){
        Swal.fire("Ops...", "The passwords are different!", "warning");
        return;
    }

    $.ajax({
        url: "/update-password",
        method: "POST",
        data: {
            password: $('#password').val(),
            newPassword: $('#new-password').val()
        }
    }).done(function(){
        Swal.fire("Success!", "The password was updated!", "success")
        .then(function() {
            window.location = "/profile";
        });
    }).fail(function(){
        Swal.fire("Ops...", "ERROR while updating the password!", "error");
    })
}

function deleteUser(evt){
    Swal.fire({
        title: "Atention!",
        text: "Are you sure you want to delete your account? You won't be able to recover your account!",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Yes, delete it!'
    }).then((result) => {
        if (result.isConfirmed) {
            $.ajax({
                url: "/delete-user",
                method: "DELETE"
            }).done(function(){
                Swal.fire("Success!", "Your user was successfully deleted", "success").then(function(){
                    window.location = "/logout";
                })
            }).fail(function(){
                Swal.fire("Ops...", "ERROR while deleting your user!", "error");
            });
        }
    })
}