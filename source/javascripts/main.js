//= require jquery
//= require bootstrap


$(window).scroll(function () {
  if ($(window).scrollTop() <= 50) {
    $('#devopsdays-menu').removeClass('scrolled')
  } else {
    $('#devopsdays-menu').addClass('scrolled')
  }
});
$(window).load(function () {
  if ($(window).scrollTop() <= 30) {
    $('#devopsdays-menu').removeClass('scrolled')
  } else {
    $('#devopsdays-menu').addClass('scrolled')
  }
});
