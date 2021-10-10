$('#new-post').on('submit', createPost)

$(document).on('click', '.like-post', likePost)
$(document).on('click', '.dislike-post', dislikePost)

$('#update-post').on('click', updatePost)



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
        alert("ERROR while creating the post")
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
        alert('ERROR' + postId)
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
        alert('ERROR' + postId)
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
        alert("post edited")
    }).fail(function(){
        alert("Error while editing the post")
    }).always(function(){
        $('#update-post').prop('disabled', false);
    })
}