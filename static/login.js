
    $(document).ready(function(){    
            
            $('#username-exists').hide();
            $('#account-created').hide();
            $('#shortInput').hide();
            $('.white-panel').css('transition','.3s ease-in-out');

            $('#wrong-password').attr('data')=='0'?$('#wrong-password').show():$('#wrong-password').hide();
            $('#new-account').attr('data')=='-1'?$('#new-account').show():$('#new-account').hide();

            $('.login-info-box').fadeOut();
            $('.login-show').addClass('show-log-panel');

            if($('.login-reg-panel').attr('data')=='true')
            {
                $('.white-panel').css('transition','0s');
                $('#label-login').click();
                $('#regSubmit').attr('data-userExists')=='true'? $('#username-exists').show(): $('#account-created').show();   
            }


     


    });
    
     
    $('input[type="radio"][name=active-log-panel]').on('change', function() {
        $('#new-account').hide();
        $('#wrong-password').hide();
        $('#username-exists').hide();
        $('#account-created').hide();   
        $('#shortInput').hide();
        $('#usernameId').val("");
        $('#passwordId').val("");
        $('#RegUsernameId').val("");
        $('#RegPasswordId').val("");

        if($('#log-login-show').is(':checked')) {
            $('.register-info-box').fadeOut(); 
            $('.login-info-box').fadeIn();
            
            $('.white-panel').addClass('right-log');
            $('.register-show').addClass('show-log-panel');
            $('.login-show').removeClass('show-log-panel');
            $('#log-login-show').prop('checked',false);
            
        }
        else if($('#log-reg-show').is(':checked')) {
            $('.register-info-box').fadeIn();
            $('.login-info-box').fadeOut();
            
            $('.white-panel').removeClass('right-log');
            
            $('.login-show').addClass('show-log-panel');
            $('.register-show').removeClass('show-log-panel');
            $('#log-reg-show').prop('checked',false);
        }
    });

    window.addEventListener( "pageshow", function ( event ) {
        var historyTraversal = event.persisted || 
                               ( typeof window.performance != "undefined" && 
                                    window.performance.navigation.type === 2 );
        if ( historyTraversal ) {
          // Handle page restore.
          window.location.reload();
        }
      });

      function logInF(){
        if($('#usernameId').val().length<3 || $('#passwordId').val().length<3 || $('#usernameId').val().length>30 || $('#passwordId').val().length>30)
        {
            $('#shortInput').show();
        }

        else{
        $('#loginForm').submit();
        }
      }

      function regF(){

        if(($('#RegUsernameId').val()).length<3 || ($('#RegPasswordId').val()).length<3 || ($('#RegUsernameId').val()).length>30 || ($('#RegPasswordId').val()).length>30)
        {
            $('#shortInput').show();
        }

        else{
            $('#regForm').submit();
        }
        
      }

      var input = document.getElementById("loginForm");
      input.addEventListener("keypress", function(event) {
        if (event.key === "Enter") {
          event.preventDefault();
        document.getElementById("regSubmit").click();
        }
      });

      var input = document.getElementById("regForm");
      input.addEventListener("keypress", function(event) {
        if (event.key === "Enter") {
          event.preventDefault();
        document.getElementById("registerSubmit").click();
        }
      });

      document.addEventListener('click', function(event) {
        if ((!document.getElementById("regForm").contains(event.target)||!document.getElementById("loginForm").contains(event.target))&& event.target.type!="button")
       { $('#new-account').hide();
        $('#wrong-password').hide();
        $('#username-exists').hide();
        $('#account-created').hide();   
        $('#shortInput').hide();
      }

      });
   
      



    