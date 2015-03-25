$(document).on('submit', '#new_video', function(e) {
  e.preventDefault();

  var $form = $(this);

  $.post($form.attr('action'), $form.serialize(), function(json) {
    console.log(json);
    // Populate a video template and prepend it in the #videos section
  }, 'json')
})

$(document).on('click', '#videos a', function(e) {
  e.preventDefault();

  var $link = $(this);

  $.getJSON('/items/', { url: $link.data('url') }, function(json) {
    console.log(json);
    // Replace the item's thumbnail with its iframe (json.html)
  })
})
