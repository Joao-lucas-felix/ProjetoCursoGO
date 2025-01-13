$('#form-create-user').on('submit', createUser);
function createUser(event){
    event.preventDefault();

    if ($('#password').val() != $('#confirm-password').val()){
        alert('The passwords are wrong!')
        return;
    }

    $.ajax(
        {
            url: '/users',
            method: 'POST',
            data:{
                name: $('#name').val(),
                email: $('#email').val(),
                nick:  $('#nick').val(),
                password: $('#password').val(),
            }
        }
    ).done(() => {
        alert('User created')
    }).fail((error) => {
        console.log(error);
        
        alert('Operation fail')
    })
}