$('#new-post').on('submit', createPost);

$(document).on('click', '.like-post', likePost);
$(document).on('click', '.dislike-post', dislikePost);

$('#update-post').on('click', updatePost);
$('.delete-post').on('click', deletePost);



function createPost(event){
    event.preventDefault();

    $.ajax({
        url: "/posts",
        method : "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        }
    }).done(function(){
        window.location="/home"
    }).fail(function(){
        Swal.fire("Ops...", "ERROR while creating the post", "error")
    })
}

function likePost(event){
    event.preventDefault();

    const clickedElement = $(event.target);
    const postId = clickedElement.closest('div').data('post-id');

    clickedElement.prop('disabled', true);
    $.ajax({
        url: `/posts/${postId}/like`,
        method: "POST"
    }).done(function(){
        const likeCounter = clickedElement.next('span');
        const likes = parseInt(likeCounter.text());

        likeCounter.text(likes + 1);

        clickedElement.addClass('dislike-post')
        clickedElement.addClass('text-primary')
        clickedElement.removeClass('like-post')
    }).fail(function(){
        Swal.fire("Ops...", "ERROR", "error")
    }).always(function (){
        clickedElement.prop('disabled', false);
    });
}

function dislikePost(event){
    event.preventDefault();

    const clickedElement = $(event.target);
    const postId = clickedElement.closest('div').data('post-id');

    clickedElement.prop('disabled', true);
    $.ajax({
        url: `/posts/${postId}/dislike`,
        method: "POST"
    }).done(function(){
        const likeCounter = clickedElement.next('span');
        const likes = parseInt(likeCounter.text());

        likeCounter.text(likes - 1);

        clickedElement.addClass('like-post')
        clickedElement.removeClass('text-primary')
        clickedElement.removeClass('dislike-post')
    }).fail(function(){
        Swal.fire("Ops...", "ERROR", "error")
    }).always(function (){
        clickedElement.prop('disabled', false);
    });
}

function updatePost(event){
    $(this).prop('disabled', true);
    event.preventDefault();

    const postId = $(this).data('post-id');

    $.ajax({
        url: `/posts/${postId}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val(),
        }
    }).done(function (){
        Swal.fire('Success!', 'Post updated with success!', 'success').then(function() {
            window.location = "/home";
        })
    }).fail(function(){
        Swal.fire("Ops...", "Error while editing the post!", "error");
    }).always(function(){
        $('#update-post').prop('disabled', false);
    })
}

function deletePost(event){
    event.preventDefault();

    Swal.fire({
        title: 'Are you sure?',
        text: "You won't be able to revert this!",
        icon: 'warning',
        showCancelButton: true,
        confirmButtonColor: '#3085d6',
        cancelButtonColor: '#d33',
        confirmButtonText: 'Yes, delete it!'
    }).then((result) => {
        if (result.isConfirmed) {
            const clickedElement = $(event.target);
            const post = clickedElement.closest('div');
            const postId = post.data('post-id');
        
            clickedElement.prop('disabled', true);
            $.ajax({
                url: `/posts/${postId}`,
                method: "DELETE",
            }).done(function(){
                post.fadeOut("slow", function(){
                    $(this).remove();
                });
            }).fail(function(){
                Swal.fire("Ops...", "ERROR while deleting the post!", "error")
            })
        }
    })
}