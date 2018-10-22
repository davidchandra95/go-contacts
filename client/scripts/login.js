
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
         getUser(email, password);
      }
   })

   $('#tes').click(function (e) {
      e.preventDefault()
      axios({
         method: 'GET',
         url: 'https://go-contacts-david.herokuapp.com/api/me/contacts/2',
         headers: {
            'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIsIlVzZXJuYW1lIjoiIn0.2n6aTppHGI4YWIBv2jCam-0OXj5_NX-dwGb7Z425Wiw',
            'Content-Type': 'application/x-www-form-urlencoded'
         }
      })
         .then(function (data) {
            console.log(data);
         })
         .catch(function (error) {
            console.log(error);
         });
   })
});

function getUser(email, password) {
   axios({
      method: 'POST',
      url: 'go-contacts-david.herokuapp.com/api/user/login',
      data: {
         email: email,
         password: password
      },
      headers: {
         'Content-Type': 'application/json',
      }
   }).then(function (response) {
      return response.JSON()
   })
      .then(function (data) {
         console.log(data);
      })
      .catch(function (error) {
         console.log(error);
      });
}
