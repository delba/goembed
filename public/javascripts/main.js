$(document).on('submit', '#new_video', function(e) {
  e.preventDefault();

  var $form = $(this);

  $.post($form.attr('action'), $form.serialize(), function(json) {
    console.log(json);
    // Populate a video template and prepend it in the #videos section
  }, 'json')
})
