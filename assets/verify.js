function ValidateEmail()
{  
    var inputEmail = $("#EmailInput").val();

    if(!validateEmailRegex(inputEmail))
    {
        alert("Please input valid email!");
        return
    }
    // var data, xhr;

    // data = new FormData();
    // data.append('email', email);

    // xhr = new XMLHttpRequest();

    // xhr.open( 'POST', '/verifyEmail', true );
    // xhr.onreadystatechange = function ( response ) {
    //     alert(JSON.parse(response));
    // };
    // xhr.send( data );


    $.ajax({
                url: "/verifyEmail",
                type: "POST",
                data: { email: inputEmail},
                success: function (result) {
                    var response = JSON.parse(result);
                    if(response.result)
                        $('#ResultLabel').text("Your email is compromised!");
                    else
                        $('#ResultLabel').text("Your email is safe!");
                    
                },
                error: function (result) {
                    alert(JSON.parse(result));
                }
            });
}

function validateEmailRegex(email) {
    const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}
