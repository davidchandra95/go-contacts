
$(document).ready(function () {
   $('.ui.form').form({
      fields: {
         email: {
            identifier: 'email',
            rules: [{
               type: 'empty',
               prompt: 'Please enter your e-mail'
            },
            {
               type: 'email',
               prompt: 'Please enter a valid e-mail'
            }
            ]
         },
         password: {
            identifier: 'password',
            rules: [{
               type: 'empty',
               prompt: 'Please enter your password'
            },
            {
               type: 'length[6]',
               prompt: 'Your password must be at least 6 characters'
            }
            ]
         }
      }
   });

   $('#btnLogin').click(function (event) {
      event.preventDefault();
      if ($('.ui.form').form('is valid')) {
         let email = $('#email').val();
         let password = $('#password').val();
         Login(email, password);
      }
   })

   $('#tes').click(function (e) {
      e.preventDefault()
      axios({
         method: 'GET',
         url: 'https://go-contacts-david.herokuapp.com/api/me/contacts',
         headers: {
            'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIsIlVzZXJuYW1lIjoiIn0.2n6aTppHGI4YWIBv2jCam-0OXj5_NX-dwGb7Z425Wiw',
         },
      })
         .then(function (data) {
            console.log(data);
         })
         .catch(function (error) {
            console.log(error);
         });
   })
});

function Login(email, password) {
   axios({
      method: 'POST',
      url: 'https://go-contacts-david.herokuapp.com/api/user/login',
      data: {
         email: email,
         password: password
      }
   })
      .then(function (data) {
         console.log(data);
      })
      .catch(function (error) {
         console.log(error);
      });
}
