$('#new-user-register').on('submit',createUser);


function createUser(event){
    event.preventDefault();
    console.log("inside of the function")

    if ($('#password').val() != $('#confirmPassword').val()){
        alert("Password does not match");
        return;
    }

    $.ajax({
        url:"/create-user",
        method: "POST",
        data:{
            
        }
    }


    )
}