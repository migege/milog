$('a.external').click(function(){
  window.open($(this).attr('href'),'_blank')
  return false
})
$('.post-tags').select2({tags:true})
$('table').addClass("table table-striped")
$('#commentform').submit(function(){
  $.ajax({
    url:$(this).attr('action'),
    type:$(this).attr('method'),
    data:$(this).serialize(),
    dataType:'json',
    success:function(res){
      window.location.reload()
    },
    error:function(){
    },
  })
  return false
})
function reply(comment_id){
  $('#comment_parent_id').val(comment_id)
}
$('#loginform').submit(function(){
  var msg=CryptoJS.MD5($('#user_login').val()+CryptoJS.MD5($('#user_pass').val()))
  var sig=CryptoJS.HmacSHA256($('#logints').val()+msg,$('#user_login').val())
  $('#user_pass').val(sig)
  $.ajax({
    url:$(this).attr('action'),
    type:$(this).attr('method'),
    data:$(this).serialize(),
    dataType:'json',
    success:function(res){
      if(res.Code==0){
        //window.location.reload()
        window.location.href=$('#loginredirect').val()
      }else{
        /*
        $('#user_pass').val("")
        $('#loginerror').text(res.Message)
        */
        window.location.reload()
      }
    },
    error:function(){
    },
  })
  return false
})
function logout(){
  $.ajax({
    url:'/logout',
    type:'get',
    success:function(res){
      if(res.Code==0){
        window.location.href='/login'
      }
    },
  })
}
