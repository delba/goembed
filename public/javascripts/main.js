var itemTemplate = _.template($('#item-template').html())

$(document).on('submit', '#new_item', function(e) {
  e.preventDefault();

  var $form = $(this);

  $.post($form.attr('action'), $form.serialize(), function(json) {
    var html = itemTemplate({item: json});
    $('#items').prepend(html)
    $form.get(0).reset()
  }, 'json')
})

$(document).on('click', '#items a', function(e) {
  e.preventDefault();

  var $link = $(this), id = $link.data('id');

  $.getJSON('/items/' + id, function(json) {
    console.log(json);
    // Replace the item's thumbnail with its iframe (json.html)
  })
})

$(document).on({
  mouseenter: function(e) {
    $(this).addClass('overlay');
  },
  mouseleave: function(e) {
    $(this).removeClass('overlay');
  }
}, '#items a')
