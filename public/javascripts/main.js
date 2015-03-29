$(document).on('submit', '#new_item', function(e) {
  e.preventDefault();

  var $form = $(this);

  $.post($form.attr('action'), $form.serialize(), function(json) {
    console.log(json);
    // Populate an item template and prepend it in the #items section
  }, 'json')
})

$(document).on('click', '#items a', function(e) {
  e.preventDefault();

  var $link = $(this);

  $.getJSON('/items/', { id: $link.data('id') }, function(json) {
    console.log(json);
    // Replace the item's thumbnail with its iframe (json.html)
  })
})
